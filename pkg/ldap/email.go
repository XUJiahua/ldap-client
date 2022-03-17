package ldap

import (
	"errors"
	"strings"

	"github.com/go-ldap/ldap/v3"
)

var ErrInvalidEmail = errors.New("invalid email, expect xxx.xxx@xxx")

func EmailToCN(email string) (string, error) {
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return "", ErrInvalidEmail
	}
	cn := parts[0]

	return cn, nil
}

func cnToDN(cn string, baseDN string) string {
	return "cn=" + cn + "," + baseDN
}

func ouToDN(ou string, baseDN string) string {
	return "ou=" + ou + "," + baseDN
}

func genAttributes(email string) (map[string][]string, error) {
	cn, err := EmailToCN(email)
	if err != nil {
		return nil, err
	}
	parts := strings.Split(cn, ".")
	if len(parts) != 2 {
		return nil, ErrInvalidEmail
	}
	givenName := parts[0]
	sn := parts[1]
	attributes := make(map[string][]string)
	attributes["mail"] = []string{email}
	attributes["cn"] = []string{cn}
	attributes["givenname"] = []string{givenName}
	attributes["sn"] = []string{sn}
	attributes["objectclass"] = []string{"inetOrgPerson", "top"}
	// dn is the entry key, not entry attribute
	//attributes["dn"] = []string{strings.Join([]string{"cn=" + cn, baseDN}, ",")}

	return attributes, nil
}

func toLDAPAttributes(attributes map[string][]string) []ldap.Attribute {
	var ldapAttributes []ldap.Attribute
	for k, v := range attributes {
		ldapAttributes = append(ldapAttributes, ldap.Attribute{
			Type: k,
			Vals: v,
		})
	}
	return ldapAttributes
}
