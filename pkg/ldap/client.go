package ldap

import (
	"errors"
	"strings"

	"github.com/go-ldap/ldap/v3"
	"github.com/sirupsen/logrus"
)

type Client struct {
	l *ldap.Conn
	// where to search and add for users
	peopleBaseDN string
	groupBaseDN  string
}

func NewClient(url, adminUser, adminPassword, peopleBaseDN, groupBaseDN string) (*Client, error) {
	l, err := ldap.DialURL(url)
	if err != nil {
		return nil, err
	}
	err = l.Bind(adminUser, adminPassword)
	if err != nil {
		return nil, err
	}
	return &Client{l: l, peopleBaseDN: peopleBaseDN, groupBaseDN: groupBaseDN}, nil
}

func (c Client) ResetPasswordByEmail(email string) (string, error) {
	cn, err := EmailToCN(email)
	if err != nil {
		return "", err
	}
	userDN := cnToDN(cn, c.peopleBaseDN)
	return c.ResetPassword(userDN)
}

func (c Client) ResetPassword(userDN string) (string, error) {
	// Prepare the password change request
	// newPassword = "", means LDAP will generate a random password for us
	changeReq := ldap.NewPasswordModifyRequest(userDN, "", "")

	// Send the password change request
	result, err := c.l.PasswordModify(changeReq)
	if err != nil {
		return "", err
	}

	logrus.Debug(result)

	return result.GeneratedPassword, nil
}

// CreateAccount creates a new account entry in the LDAP directory without specify password, and you need reset the password
func (c Client) CreateAccount(email string) error {
	attributes, err := genAttributes(email)
	if err != nil {
		return err
	}
	logrus.Debug(attributes)

	dn := strings.Join([]string{"cn=" + attributes["cn"][0], c.peopleBaseDN}, ",")
	return c.l.Add(&ldap.AddRequest{
		DN:         dn,
		Attributes: toLDAPAttributes(attributes),
		Controls:   nil,
	})
}

func (c Client) GetGroups() ([]*ldap.Entry, error) {
	logrus.Info("getting groups...")
	searchRequest := ldap.NewSearchRequest(
		c.groupBaseDN,
		2, //ScopeWholeSubtree
		0, //NeverDerefAliases
		0,
		0,
		false,
		"(&(objectClass=groupOfUniqueNames))", // The filter to apply
		[]string{"dn", "uniqueMember"},        // A list attributes to retrieve
		nil,
	)

	sr, err := c.l.Search(searchRequest)
	if err != nil {
		return nil, err
	}

	return sr.Entries, nil
}

// GetGroup get group by group name, group attribute is dn and uniqueMember
func (c Client) GetGroup(groupDN string) (*ldap.Entry, error) {
	logrus.Info("getting group...")
	entries, err := c.GetGroups()
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if entry.DN == groupDN {
			return entry, nil
		}
	}

	return nil, nil
}

func (c Client) AddUserToGroup(userDN, groupDN string) error {
	logrus.Info("adding user to group...")
	group, err := c.GetGroup(groupDN)
	if err != nil {
		return err
	}

	if group == nil {
		return errors.New("group not found")
	}

	if contains(group.GetAttributeValues("uniqueMember"), userDN) {
		logrus.Infof("user %s is already in group %s", userDN, groupDN)
		return nil
	}

	modify := ldap.NewModifyRequest(group.DN, nil)
	newMembers := group.GetAttributeValues("uniqueMember")
	newMembers = append(newMembers, userDN)
	modify.Replace("uniqueMember", newMembers)

	return c.l.Modify(modify)
}

func (c Client) RemoveUserFromGroup(userDN, groupDN string) error {
	logrus.Info("removing user from group...")
	group, err := c.GetGroup(groupDN)
	if err != nil {
		return err
	}

	if group == nil {
		return errors.New("group not found")
	}

	if !contains(group.GetAttributeValues("uniqueMember"), userDN) {
		logrus.Infof("user %s is not in group %s", userDN, groupDN)
		return nil
	}

	modify := ldap.NewModifyRequest(group.DN, nil)
	oldMembers := group.GetAttributeValues("uniqueMember")
	// remove userDN from oldMembers
	var newMembers []string
	for _, v := range oldMembers {
		if v != userDN {
			newMembers = append(newMembers, v)
		}
	}
	modify.Replace("uniqueMember", newMembers)

	return c.l.Modify(modify)
}
func contains(ss []string, s string) bool {
	for _, v := range ss {
		if v == s {
			return true
		}
	}
	return false
}

func (c Client) AddUserToGroupEasy(email, groupName string) error {
	cn, err := EmailToCN(email)
	if err != nil {
		return err
	}
	userDN := cnToDN(cn, c.peopleBaseDN)
	groupDN := ouToDN(groupName, c.groupBaseDN)

	return c.AddUserToGroup(userDN, groupDN)
}

func (c Client) RemoveUserFromGroupEasy(email, groupName string) error {
	cn, err := EmailToCN(email)
	if err != nil {
		return err
	}
	userDN := cnToDN(cn, c.peopleBaseDN)
	groupDN := ouToDN(groupName, c.groupBaseDN)

	return c.RemoveUserFromGroup(userDN, groupDN)
}
