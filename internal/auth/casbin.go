package auth

import (
	"github.com/casbin/casbin/v2"
	"log"
)

var enforcer *casbin.Enforcer

func InitCasbin() {
	var err error
	enforcer, err = casbin.NewEnforcer("configs/auth_model.conf", "configs/auth_policy.csv")
	if err != nil {
		log.Fatalf("[error] init casbin enforcer failed: %v", err)
	}
}

func CheckPermission(userID, action, resource string) bool {
	ok, err := enforcer.Enforce(userID, action, resource)
	if err != nil {
		log.Fatalf("[error] check permission failed: %v", err)
		return false
	}

	return ok
}
