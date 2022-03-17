package ldap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_genAttributes(t *testing.T) {
	attributes, err := genUserAttributes("tom.xu@example.com")
	assert.NoError(t, err)
	assert.EqualValues(t, map[string][]string{
		"objectclass": []string{"inetOrgPerson", "top"},
		"cn":          []string{"tom.xu"},
		"sn":          []string{"xu"},
		"givenname":   []string{"tom"},
		"mail":        []string{"tom.xu@example.com"},
	}, attributes)
}
