package database

import (
	"time"

	"github.com/labstack/echo"
	"github.com/quocbang/api-server-basic/impl/requests"
)

type Services interface {
	Users() Users
	Tasks() Tasks
}

type Users interface {
	CreateAccount(echo.Context, requests.CreateAccountRequest) (requests.CreateAccountReply, error)
	DeleteAccounts(echo.Context, requests.DeleteAccountRequest) (requests.DeleteAccountReply, error)
	Login(echo.Context, requests.LoginRequest) (requests.LoginReply, error)
	SignOut(echo.Context, string, time.Duration) error
}

type Tasks interface {
	GetTasks(echo.Context, requests.GetTaskRequest) (requests.GetTaskReply, error)
	CreateTasks(echo.Context, requests.CreateTaskRequest) (requests.CreateTaskReply, error)
	UpdateTask(echo.Context, requests.UpdateTaskRequest) (requests.UpdateTaskReply, error)
	DeleteTasks(echo.Context, requests.DeleteTaskRequest) (requests.DeleteTaskReply, error)
}
