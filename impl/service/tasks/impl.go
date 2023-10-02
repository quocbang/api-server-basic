package tasks

import (
	"net/http"

	"github.com/labstack/echo"

	requestCm "gitlab.com/quocbang/common-util/requests"

	"github.com/quocbang/api-server-basic/database"
	localErr "github.com/quocbang/api-server-basic/errors"
	"github.com/quocbang/api-server-basic/impl/requests"
	serviceImpl "github.com/quocbang/api-server-basic/impl/service"
	"github.com/quocbang/api-server-basic/middleware/context"
	"github.com/quocbang/api-server-basic/utils/function"
	"github.com/quocbang/api-server-basic/utils/roles"
)

type services struct {
	dm database.DataManager
}

func NewService(dm database.DataManager) serviceImpl.TaskService {
	return &services{dm: dm}
}

// CreateTasks implementation.
func (s *services) CreateTasks(ctx echo.Context) error {
	// check permission.
	if !roles.HasPermission(function.FunctionOperationID_CREATE_TASKS, context.GetPrincipal(ctx).Roles) {
		return ctx.JSON(http.StatusForbidden, localErr.Error{
			Code:   localErr.Code_FORBIDDEN,
			Detail: "permission denied",
		})
	}

	// bind request.
	req, err := requestCm.BindRequest[requests.CreateTaskRequest](ctx)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	// check request struct.
	if err := requestCm.ValidateStruct[requests.CreateTaskRequest](req); err != nil {
		return ctx.JSON(http.StatusBadRequest, localErr.Error{
			Code:   localErr.Code_REQUEST_INSUFFICIENT,
			Detail: err.Error(),
		})
	}

	// do insert into database.
	reply, err := s.dm.Tasks().CreateTasks(ctx, req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, reply)
}

// ListTasks implementation.
func (s *services) ListTasks(ctx echo.Context) error {
	// check permission.
	if !roles.HasPermission(function.FunctionOperationID_GET_TASKS, context.GetPrincipal(ctx).Roles) {
		return ctx.JSON(http.StatusForbidden, localErr.Error{
			Code:   localErr.Code_FORBIDDEN,
			Detail: "permission denied",
		})
	}

	// bind request.
	req, err := requestCm.BindRequest[requests.GetTaskRequest](ctx)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	// check request struct.
	reply, err := s.dm.Tasks().GetTasks(ctx, req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, reply)
}

// UpdateTask implementation.
func (s *services) UpdateTask(ctx echo.Context) error {
	// check permission.
	if !roles.HasPermission(function.FunctionOperationID_UPDATE_TASK, context.GetPrincipal(ctx).Roles) {
		return ctx.JSON(http.StatusForbidden, localErr.Error{
			Code:   localErr.Code_FORBIDDEN,
			Detail: "permission denied",
		})
	}

	req, err := requestCm.BindRequest[requests.UpdateTaskRequest](ctx)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	if err := requestCm.ValidateStruct[requests.UpdateTaskRequest](req); err != nil {
		return ctx.JSON(http.StatusBadRequest, localErr.Error{
			Code:   localErr.Code_REQUEST_INSUFFICIENT,
			Detail: err.Error(),
		})
	}

	reply, err := s.dm.Tasks().UpdateTask(ctx, req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, reply)
}

// DeleteTasks implementation.
func (s *services) DeleteTasks(ctx echo.Context) error {
	// check permission.
	if !roles.HasPermission(function.FunctionOperationID_DELETE_TASKS, context.GetPrincipal(ctx).Roles) {
		return ctx.JSON(http.StatusForbidden, localErr.Error{
			Code:   localErr.Code_FORBIDDEN,
			Detail: "permission denied",
		})
	}

	req, err := requestCm.BindRequest[requests.DeleteTaskRequest](ctx)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	if err := requestCm.ValidateStruct[requests.DeleteTaskRequest](req); err != nil {
		return ctx.JSON(http.StatusBadRequest, localErr.Error{
			Code:   localErr.Code_REQUEST_INSUFFICIENT,
			Detail: err.Error(),
		})
	}

	reply, err := s.dm.Tasks().DeleteTasks(ctx, req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, reply)
}
