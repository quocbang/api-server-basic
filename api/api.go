package api

import (
	"github.com/labstack/echo"

	"github.com/quocbang/api-server-basic/api/tasks"
	"github.com/quocbang/api-server-basic/api/users"
	"github.com/quocbang/api-server-basic/impl/service"
)

const (
	API             = "/api"
	taskServicePath = "/task"
	userServicePath = "/user"
)

func RegisterAPI(e *echo.Echo, service service.DataManagerService) {
	// register tasks api service.
	tasks.RegisterTasksAPI(e, API+taskServicePath, service.Tasks())
	// register users api service.
	users.RegisterUsersAPI(e, API+userServicePath, service.Users())
}
