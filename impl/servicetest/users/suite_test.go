package users

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
			function.FunctionOperationID_CREATE_ACCOUNT.String(): {
				roles.Roles_ADMINISTRATOR.String(),
			},
			function.FunctionOperationID_DELETE_ACCOUNTS.String(): {
				roles.Roles_ADMINISTRATOR.String(),
			},
		},
	})

	return &Suite{BasicSuite: *basicSuite}
}

func TestSuite(t *testing.T) {
	s.Run(t, NewSuite())
}
