package ldap

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func init() {
	logrus.SetLevel(logrus.DebugLevel)
}

func newClient() *Client {
	c, err := NewClient("ldap://localhost:389", "cn=admin,dc=example,dc=com", "123456", "ou=people,dc=example,dc=com", "ou=group,dc=example,dc=com")
	if err != nil {
		panic(err)
	}
	return c
}

func TestClient_ResetPassword(t *testing.T) {
	c := newClient()
	newPassword, err := c.ResetPassword("cn=john.xu,ou=people,dc=example,dc=com")
	assert.NoError(t, err)
	println(newPassword)
}

func TestClient_CreateAccount(t *testing.T) {
	c := newClient()
	err := c.CreateAccount("tom.xu@example.com")
	assert.NoError(t, err)
	newPassword, err := c.ResetPasswordByEmail("tom.xu@example.com")
	assert.NoError(t, err)
	println(newPassword)
}

func TestClient_GetGroup(t *testing.T) {
	c := newClient()
	group, err := c.GetGroup("ou=admin,ou=group,dc=example,dc=com")
	assert.NoError(t, err)
	spew.Dump(group)
}

func TestClient_GetGroups(t *testing.T) {
	c := newClient()
	groups, err := c.GetGroups()
	assert.NoError(t, err)
	spew.Dump(groups)
}

func TestClient_AddUserToGroup(t *testing.T) {
	c := newClient()
	err := c.RemoveUserFromGroup("cn=tom.xu,ou=people,dc=example,dc=com", "ou=admin,ou=group,dc=example,dc=com")
	assert.NoError(t, err)
	TestClient_GetGroups(t)
	err = c.AddUserToGroup("cn=tom.xu,ou=people,dc=example,dc=com", "ou=admin,ou=group,dc=example,dc=com")
	assert.NoError(t, err)
	TestClient_GetGroups(t)
}
