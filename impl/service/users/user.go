package users

import (
	"net/http"
	"time"

	"github.com/labstack/echo"

	validate "gitlab.com/quocbang/common-util/requests"

	"github.com/quocbang/api-server-basic/database"
	localErr "github.com/quocbang/api-server-basic/errors"
	"github.com/quocbang/api-server-basic/impl/requests"
	"github.com/quocbang/api-server-basic/impl/service"
	"github.com/quocbang/api-server-basic/middleware/authorization"
	"github.com/quocbang/api-server-basic/middleware/context"
	"github.com/quocbang/api-server-basic/utils/function"
	"github.com/quocbang/api-server-basic/utils/roles"
)

type services struct {
	dm database.DataManager
}

func NewService(dm database.DataManager) service.UsersService {
	return &services{dm: dm}
}

// CreateAccount is create new account.
func (s *services) CreateAccount(ctx echo.Context) error {
	if !roles.HasPermission(function.FunctionOperationID_CREATE_ACCOUNT, context.GetPrincipal(ctx).Roles) {
		return ctx.JSON(http.StatusForbidden, localErr.Error{
			Code:   localErr.Code_FORBIDDEN,
			Detail: "permission denied",
		})
	}

	req, err := validate.BindRequest[requests.CreateAccountRequest](ctx)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	if err := validate.ValidateStruct[requests.CreateAccountRequest](req); err != nil {
		return ctx.JSON(http.StatusBadRequest, localErr.Error{
			Code:   localErr.Code_REQUEST_INSUFFICIENT,
			Detail: err.Error(),
		})
	}

	rowsAffected, err := s.dm.Users().CreateAccount(ctx, req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, requests.CreateAccountReply{RowsAffected: rowsAffected.RowsAffected})
}

// Login is sign-in user account and receiving the,
// token for access to other methods.
func (s *services) Login(ctx echo.Context) error {
	req, err := validate.BindRequest[requests.LoginRequest](ctx)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	if err := validate.ValidateStruct[requests.LoginRequest](req); err != nil {
		return ctx.JSON(http.StatusBadRequest, localErr.Error{
			Code:   localErr.Code_REQUEST_INSUFFICIENT,
			Detail: err.Error(),
		})
	}

	token, err := s.dm.Users().Login(ctx, req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, token)
}

// DeleteAccounts is delete existing accounts.
func (s *services) DeleteAccounts(ctx echo.Context) error {
	if !roles.HasPermission(function.FunctionOperationID_DELETE_ACCOUNTS, context.GetPrincipal(ctx).Roles) {
		return ctx.JSON(http.StatusForbidden, localErr.Error{
			Code:   localErr.Code_FORBIDDEN,
			Detail: "permission denied",
		})
	}

	req, err := validate.BindRequest[requests.DeleteAccountRequest](ctx)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, localErr.Error{
			Code:   http.StatusBadRequest,
			Detail: err.Error(),
		})
	}

	if err := validate.ValidateStruct[requests.DeleteAccountRequest](req); err != nil {
		return ctx.JSON(http.StatusBadRequest, localErr.Error{
			Code:   localErr.Code_REQUEST_INSUFFICIENT,
			Detail: err.Error(),
		})
	}

	// check whether delete it self or not.
	if email := context.GetPrincipal(ctx).Email; email != "" {
		for _, v := range req.Emails {
			if v == email {
				return ctx.JSON(http.StatusBadRequest, localErr.Error{
					Code:   localErr.Code_NOT_SUPPORTED_DELETE_YOUR_SELF,
					Detail: "can not delete your self",
				})
			}
		}
	}

	reply, err := s.dm.Users().DeleteAccounts(ctx, req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, reply)
}

// Logout is sign-out
func (s *services) Logout(ctx echo.Context) error {
	token := ctx.Request().Header.Get(authorization.AuthorizationKey)
	if token == "" {
		return ctx.JSON(http.StatusNonAuthoritativeInfo, localErr.Error{
			Code:   localErr.Code_NONE,
			Detail: "token key not found",
		})
	}

	expiryTime := context.GetPrincipal(ctx).ExpiryTime
	if expiryTime.Before(time.Now()) {
		return ctx.JSON(http.StatusOK, "Logout successfully")
	}

	timeRemaining := time.Until(expiryTime)
	err := s.dm.Users().SignOut(ctx, token, timeRemaining)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, "Logout successfully")
}
