package security

import (
	"testing"

	"github.com/Alonso-Arias/test-cleverit/services/model"
	"github.com/stretchr/testify/assert"
)

func Test_IsAuthorized(t *testing.T) {

	au := model.AuthenticatedUser{
		Roles: []model.Role{{Code: "ROL_1"}},
	}

	assert.True(t, IsAuthorized(au, "GET", "/api/v1/user/sksksks"))
}
