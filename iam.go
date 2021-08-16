package main

import (
	"log"

	"github.com/casbin/casbin/v2"
	redisadapter "github.com/casbin/redis-adapter/v2"
)

const (
	ROLE_ADMIN   = "admin"
	ROLE_TEACHER = "teacher"

	PERMISSION_READ   = "read"
	PERMISSION_CREATE = "create"
	PERMISSION_UPDATE = "update"
	PERMISSION_DELETE = "delete"
)

func NewIam() *casbin.Enforcer {
	adapter := redisadapter.NewAdapter("tcp", "localhost:6379")
	enforcer, err := casbin.NewEnforcer("model.conf", adapter)
	if err != nil {
		log.Fatalf("fail to configure casbin: %s", err.Error())
	}

	policiesWillEnforce := [][]interface{}{
		{ROLE_ADMIN, "students", PERMISSION_READ},
		{ROLE_ADMIN, "students", PERMISSION_CREATE},
		{ROLE_ADMIN, "students", PERMISSION_UPDATE},
		{ROLE_ADMIN, "students", PERMISSION_DELETE},
		{ROLE_TEACHER, "students", PERMISSION_READ},
	}
	for _, policy := range policiesWillEnforce {
		_, err = enforcer.AddPolicy(policy...)
		if err != nil {
			log.Fatalf("fail to add policy: %s", err.Error())
		}
	}

	groupsWillEnforce := [][]interface{}{
		{"tes1", ROLE_ADMIN},
		{"tes2", ROLE_TEACHER},
	}
	for _, group := range groupsWillEnforce {
		_, err = enforcer.AddGroupingPolicy(group...)
		if err != nil {
			log.Fatalf("fail to add policy: %s", err.Error())
		}
	}

	return enforcer
}
