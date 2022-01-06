package ldap

import (
	"github.com/go-ldap/ldap/v3"
	"github.com/sirupsen/logrus"
	"strings"
)

type Client struct {
	l *ldap.Conn
	// where to search and add for users
	peopleDN string
}

func NewClient(url, peopleDN, adminUser, adminPassword string) (*Client, error) {
	l, err := ldap.DialURL(url)
	if err != nil {
		return nil, err
	}
	err = l.Bind(adminUser, adminPassword)
	if err != nil {
		return nil, err
	}
	return &Client{l: l, peopleDN: peopleDN}, nil
}

func (c Client) ResetPasswordByEmail(email string) (string, error) {
	cn, err := EmailToCN(email)
	if err != nil {
		return "", err
	}
	userDN := cnToDN(cn, c.peopleDN)
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

	dn := strings.Join([]string{"cn=" + attributes["cn"][0], c.peopleDN}, ",")
	return c.l.Add(&ldap.AddRequest{
		DN:         dn,
		Attributes: toLDAPAttributes(attributes),
		Controls:   nil,
	})
}
