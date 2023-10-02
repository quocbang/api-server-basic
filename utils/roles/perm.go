package roles

import (
	"github.com/quocbang/api-server-basic/utils/function"
)

var permission map[function.FunctionOperationID]map[Roles]interface{}

func HasPermission(operationID function.FunctionOperationID, roles []Roles) bool {
	if rls, ok := permission[operationID]; ok {
		for _, role := range roles {
			if _, ok := rls[role]; ok {
				return true
			}
		}
	}
	return false
}

func InitializePermission(roles map[string][]string) {
	permission = make(map[function.FunctionOperationID]map[Roles]interface{}, len(roles))

	for id, roles := range roles {
		permission[function.FunctionOperationID(function.FunctionOperationID_value[id])] = parseRoles(roles)
	}
}

func parseRoles(r []string) map[Roles]interface{} {
	roles := make(map[Roles]interface{}, len(r))
	for _, role := range r {
		roles[Roles(Roles_value[role])] = nil
	}
	return roles
}
