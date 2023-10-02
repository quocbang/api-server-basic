package impl

import (
	"github.com/quocbang/api-server-basic/impl/service"
	"github.com/quocbang/api-server-basic/impl/service/tasks"
	"github.com/quocbang/api-server-basic/impl/service/users"
)

// Task is task service that config data for method of this service.
func (s *ServiceConfig) Tasks() service.TaskService {
	return tasks.NewService(s.DM)
}

// Users is users service that reparing data for all user services.
func (s *ServiceConfig) Users() service.UsersService {
	return users.NewService(s.DM)
}
