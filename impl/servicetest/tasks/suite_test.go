package tasks

import (
	"testing"

	s "github.com/stretchr/testify/suite"

	"github.com/quocbang/api-server-basic/impl/servicetest/internal/suite"
	"github.com/quocbang/api-server-basic/utils/function"
	"github.com/quocbang/api-server-basic/utils/roles"
)

type Suite struct {
	suite.BasicSuite
}

func NewSuite() *Suite {
	basicSuite := suite.NewSuite(suite.Setup{
		Roles: map[string][]string{
			function.FunctionOperationID_CREATE_TASKS.String(): {
				roles.Roles_ADMINISTRATOR.String(),
				roles.Roles_LEADER.String(),
			},
			function.FunctionOperationID_GET_TASKS.String(): {
				roles.Roles_ADMINISTRATOR.String(),
				roles.Roles_LEADER.String(),
				roles.Roles_USER.String(),
			},
			function.FunctionOperationID_UPDATE_TASK.String(): {
				roles.Roles_ADMINISTRATOR.String(),
				roles.Roles_LEADER.String(),
			},
			function.FunctionOperationID_DELETE_TASKS.String(): {
				roles.Roles_ADMINISTRATOR.String(),
				roles.Roles_LEADER.String(),
			},
		},
	})

	return &Suite{BasicSuite: *basicSuite}
}

func TestSuite(t *testing.T) {
	s.Run(t, NewSuite())
}
