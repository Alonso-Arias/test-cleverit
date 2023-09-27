package security

import (
	"os"

	"github.com/Alonso-Arias/test-cleverit/services/model"
	"github.com/casbin/casbin/v2"
)

var e *casbin.Enforcer

func init() {

	BASE_PATH := os.Getenv("BASE_PATH")

	var err error

	e, err = casbin.NewEnforcer(BASE_PATH+"/security/casbin_model.conf", BASE_PATH+"/security/casbin_policy.csv")

	if err != nil {

		loggerf.WithError(err).Fatal("Failed to load access policy")

	}
}

// IsAuthorized -
func IsAuthorized(au model.AuthenticatedUser, method string, path string) bool {

	//log := loggerf.WithField("func", "IsAuthorized")

	ok := false
	for _, r := range au.Roles {
		ok, _ = e.Enforce(r.Code, path, method)
		if ok {
			break
		}
	}

	return ok

}
