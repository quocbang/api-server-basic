package users

import (
	"github.com/labstack/echo"

	"github.com/quocbang/api-server-basic/impl/service"
	"github.com/quocbang/api-server-basic/middleware/authorization"
)

func RegisterUsersAPI(e *echo.Echo, serivePath string, userService service.UsersService) {
	// CreateUser is create new user.
	//
	// DO:
	// 	- create user
	//
	// Required:
	//  - Email
	//  - Password
	//  - Roles
	e.POST(serivePath, userService.CreateAccount, authorization.Authorization)
	// Login is sign in with email and password was registed.
	//
	// DO:
	// 	- Login
	//
	// Required:
	//  - Email
	//  - Password
	e.POST(serivePath+"/login", userService.Login)
	// DeleteAccounts is delete account stored in database.
	//
	// DO:
	//  - delete
	//
	// Required:
	//  - Emails
	e.DELETE(serivePath, userService.DeleteAccounts, authorization.Authorization)
	// Logout is sign-out
	//
	// DO:
	//  - logout
	//
	// Required:
	//  - verified token key.
	e.POST(serivePath+"/logout", userService.Logout, authorization.Authorization)
}
