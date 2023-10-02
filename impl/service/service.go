package service

import (
	"github.com/labstack/echo"
)

// DataManagerService definition.
type DataManagerService interface {
	Tasks() TaskService
	Users() UsersService
}

// TaskService definition.
type TaskService interface {
	CreateTasks(echo.Context) error
	ListTasks(echo.Context) error
	UpdateTask(echo.Context) error
	DeleteTasks(echo.Context) error
}

// UserService definition.
type UsersService interface {
	CreateAccount(echo.Context) error
	DeleteAccounts(echo.Context) error
	Login(echo.Context) error
	Logout(echo.Context) error
}
