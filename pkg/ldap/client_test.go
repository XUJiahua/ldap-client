package ldap

import (
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
)

func init() {
	logrus.SetLevel(logrus.DebugLevel)
}

func TestClient_ResetPassword(t *testing.T) {
	c, err := NewClient("ldap://localhost:389", "ou=people,dc=example,dc=com", "cn=admin,dc=example,dc=com", "123456")
	assert.NoError(t, err)
	newPassword, err := c.ResetPassword("cn=john.xu,ou=people,dc=example,dc=com")
	assert.NoError(t, err)
	println(newPassword)
}

func TestClient_CreateAccount(t *testing.T) {
	c, err := NewClient("ldap://localhost:389", "ou=people,dc=example,dc=com", "cn=admin,dc=example,dc=com", "123456")
	assert.NoError(t, err)
	err = c.CreateAccount("tom.xu@example.com")
	assert.NoError(t, err)
	newPassword, err := c.ResetPasswordByEmail("tom.xu@example.com")
	assert.NoError(t, err)
	println(newPassword)
}
