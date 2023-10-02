package tasks

import (
	"net/http"
	"time"

	"github.com/quocbang/api-server-basic/database"
	"github.com/quocbang/api-server-basic/database/mocks"
	localErr "github.com/quocbang/api-server-basic/errors"
	"github.com/quocbang/api-server-basic/impl/requests"
	"github.com/quocbang/api-server-basic/impl/servicetest/internal/httptest/context"
	"github.com/quocbang/api-server-basic/impl/servicetest/internal/httptest/testexpect"
	"github.com/quocbang/api-server-basic/impl/servicetest/testutil"
	"github.com/quocbang/api-server-basic/middleware/authorization"
	"github.com/quocbang/api-server-basic/utils/roles"
	"github.com/stretchr/testify/mock"
)

func (s *Suite) TestCreateTasks() {
	assertion := s.Assertions
	testEmail := ""
	testRoles := []roles.Roles{roles.Roles_ADMINISTRATOR, roles.Roles_LEADER, roles.Roles_USER}
	testExpiryTime := time.Now().Add(time.Microsecond * 2)
	mockServer := testutil.NewMock[mock.Mock, database.Tasks]([]testutil.Scripts[mock.Mock, database.Tasks]{})

	goodRequest := func() requests.CreateTaskRequest {
		return requests.CreateTaskRequest{
			Tasks: []requests.CreateTaskDetail{
				{
					ID:             1,
					Status:         "Doing",
					ProcessPercent: 50,
					Content:        "test create tasks 01",
					EndTime:        8,
				},
			},
		}
	}

	// case permission denine.
	{
		// Arrange.
		ctx, rec, err := context.NewTestContext(http.MethodPost, "/api/task", requests.CreateTaskRequest{})
		assertion.NoError(err)
		ctx.Set(authorization.UserPrincipalKey, authorization.Principal{
			Email:      testEmail,
			Roles:      testRoles,
			ExpiryTime: testExpiryTime,
		})

		// Act.
		err = mockServer.Tasks().CreateTasks(ctx)

		// Assert.
		assertion.NoError(err)
		expected, err := testexpect.ToString(localErr.Error{
			Code:   localErr.Code_FORBIDDEN,
			Detail: "permission denied",
		})
		assertion.NoError(err)
		assertion.Equal(http.StatusForbidden, rec.Result().StatusCode)
		assertion.Equal(expected, rec.Body.String())
	}

	// missing tasks.
	{
		// Arrange.
		ctx, rec, err := context.NewTestContext(http.MethodPost, "/api/tasks", requests.CreateTaskRequest{})
		assertion.NoError(err)
		ctx.Set(authorization.UserPrincipalKey, &authorization.Principal{
			Email:      testEmail,
			Roles:      testRoles,
			ExpiryTime: testExpiryTime,
		})

		// Act.
		err = mockServer.Tasks().CreateTasks(ctx)

		// Assert.
		assertion.NoError(err)
		expected, err := testexpect.ToString(localErr.Error{
			Code:   localErr.Code_REQUEST_INSUFFICIENT,
			Detail: "missing field [Tasks]",
		})
		assertion.NoError(err)
		assertion.Equal(http.StatusBadRequest, ctx.Response().Status)
		assertion.Equal(expected, rec.Body.String())
		assertion.Equal(expected, rec.Body.String())
	}

	// missing ID.
	{
		// Arrange.
		request := goodRequest()
		request.Tasks[0].ID = 0
		ctx, rec, err := context.NewTestContext(http.MethodPost, "/api/tasks", request)
		assertion.NoError(err)
		ctx.Set(authorization.UserPrincipalKey, &authorization.Principal{
			Email:      testEmail,
			Roles:      testRoles,
			ExpiryTime: testExpiryTime,
		})

		// Act.
		err = mockServer.Tasks().CreateTasks(ctx)

		// Assert.
		assertion.NoError(err)
		expected, err := testexpect.ToString(localErr.Error{
			Code:   localErr.Code_REQUEST_INSUFFICIENT,
			Detail: "missing field [ID]",
		})
		assertion.NoError(err)
		assertion.Equal(http.StatusBadRequest, ctx.Response().Status)
		assertion.Equal(expected, rec.Body.String())
	}

	// missing content.
	{
		// Arrange.
		request := goodRequest()
		request.Tasks[0].Content = ""
		ctx, rec, err := context.NewTestContext(http.MethodPost, "/api/tasks", request)
		assertion.NoError(err)
		ctx.Set(authorization.UserPrincipalKey, &authorization.Principal{
			Email:      testEmail,
			Roles:      testRoles,
			ExpiryTime: testExpiryTime,
		})

		// Act.
		err = mockServer.Tasks().CreateTasks(ctx)

		// Assert.
		assertion.NoError(err)
		expected, err := testexpect.ToString(localErr.Error{
			Code:   localErr.Code_REQUEST_INSUFFICIENT,
			Detail: "missing field [Content]",
		})
		assertion.NoError(err)
		assertion.Equal(http.StatusBadRequest, ctx.Response().Status)
		assertion.Equal(expected, rec.Body.String())
	}

	// missing end time.
	{
		// Arrange.
		request := goodRequest()
		request.Tasks[0].EndTime = 0
		ctx, rec, err := context.NewTestContext(http.MethodPost, "/api/tasks", request)
		assertion.NoError(err)
		ctx.Set(authorization.UserPrincipalKey, &authorization.Principal{
			Email:      testEmail,
			Roles:      testRoles,
			ExpiryTime: testExpiryTime,
		})

		// Act.
		err = mockServer.Tasks().CreateTasks(ctx)

		// Assert.
		assertion.NoError(err)
		expected, err := testexpect.ToString(localErr.Error{
			Code:   localErr.Code_REQUEST_INSUFFICIENT,
			Detail: "missing field [EndTime]",
		})
		assertion.NoError(err)
		assertion.Equal(http.StatusBadRequest, ctx.Response().Status)
		assertion.Equal(expected, rec.Body.String())
	}

	// create tasks successfully.
	{
		// Arrange
		ctx, rec, err := context.NewTestContext(http.MethodPost, "/api/task", goodRequest())
		assertion.NoError(err)
		ctx.Set(authorization.UserPrincipalKey, &authorization.Principal{
			Email:      "test_user@gmail.com",
			Roles:      []roles.Roles{roles.Roles_ADMINISTRATOR, roles.Roles_LEADER, roles.Roles_USER},
			ExpiryTime: time.Now().Add(time.Minute * 2),
		})

		s := []testutil.Scripts[*mock.Mock, database.Tasks]{
			{
				MethodName: "Tasks",
				ParentReturn: func(m *mock.Mock) database.Tasks {
					return &mocks.Tasks{Mock: m}
				},
				IsParent: true,
			},
			{
				MethodName: "CreateTasks",
				Input:      append(make([]interface{}, 0), ctx, goodRequest()),
				Output:     append(make([]interface{}, 0), requests.CreateTaskReply{RowsAffected: 1}, nil),
			},
		}

		mockServer := testutil.NewMock[*mock.Mock, database.Tasks](s)

		// Act
		err = mockServer.Tasks().CreateTasks(ctx)

		// Assert
		assertion.NoError(err)
		expected, err := testexpect.ToString(requests.CreateTaskReply{
			RowsAffected: 1,
		})
		assertion.NoError(err)
		assertion.Equal(http.StatusOK, ctx.Response().Status)
		assertion.Equal(expected, rec.Body.String())
	}
}

func (s *Suite) TestGetTasks() {
	assertion := s.Assertions
	req := func() requests.GetTaskRequest {
		return requests.GetTaskRequest{
			ID: 1,
		}
	}

	// permission denined
	{
		// Arrange
		ctx, rec, err := context.NewTestContext(http.MethodGet, "/api/task", req())
		assertion.NoError(err)
		ctx.Set(authorization.UserPrincipalKey, authorization.Principal{
			Email:      "no_body@gmail.com",
			Roles:      []roles.Roles{},
			ExpiryTime: time.Now(),
		})
		mockServer := testutil.NewMock[mock.Mock, database.Tasks]([]testutil.Scripts[mock.Mock, database.Tasks]{})

		// Act
		err = mockServer.Tasks().ListTasks(ctx)

		// Assert
		assertion.NoError(err)
		expected, err := testexpect.ToString(localErr.Error{
			Code:   localErr.Code_FORBIDDEN,
			Detail: "permission denied",
		})
		assertion.NoError(err)
		assertion.Equal(http.StatusForbidden, rec.Code)
		assertion.Equal(expected, rec.Body.String())
	}

	// get task with id
	{
		// Arrange
		ctx, rec, err := context.NewTestContext(http.MethodGet, "/api/task", req())
		assertion.NoError(err)
		ctx.Set(authorization.UserPrincipalKey, &authorization.Principal{
			Email:      "admin_test@gmail.com",
			Roles:      []roles.Roles{roles.Roles_ADMINISTRATOR},
			ExpiryTime: time.Now().Add(time.Microsecond * 2),
		})
		reply := requests.GetTaskReply{
			Tasks: []requests.TaskDetail{
				{
					ID:             1,
					Status:         "Doing",
					ProcessPercent: 20,
					Content:        "test get task",
					EndTime:        time.Now().Add(time.Duration(time.Hour * 12)),
					Updated:        time.Hour.Nanoseconds(),
					Created:        time.Hour.Nanoseconds(),
				},
			},
		}
		var (
			mockInput, mockOutput []any
		)
		mockInput = append(mockInput, ctx, req())
		mockOutput = append(mockOutput, reply, nil)
		scripts := []testutil.Scripts[*mock.Mock, database.Tasks]{
			{
				MethodName: "Tasks",
				IsParent:   true,
				ParentReturn: func(m *mock.Mock) database.Tasks {
					return &mocks.Tasks{Mock: m}
				},
			},
			{
				MethodName: "GetTasks",
				Input:      mockInput,
				Output:     mockOutput,
			},
		}
		mockServer := testutil.NewMock[*mock.Mock, database.Tasks](scripts)

		// Act
		err = mockServer.Tasks().ListTasks(ctx)

		// Assert
		assertion.NoError(err)
		expected, err := testexpect.ToString(reply)
		assertion.NoError(err)
		assertion.Equal(http.StatusOK, rec.Code)
		assertion.Equal(expected, rec.Body.String())
	}

	// get task without id
	{
		// Arrange
		ctx, rec, err := context.NewTestContext(http.MethodGet, "/api/task", requests.GetTaskRequest{})
		assertion.NoError(err)
		ctx.Set(authorization.UserPrincipalKey, &authorization.Principal{
			Email:      "admin_test@gmail.com",
			Roles:      []roles.Roles{roles.Roles_ADMINISTRATOR},
			ExpiryTime: time.Now().Add(time.Microsecond * 2),
		})
		reply := requests.GetTaskReply{
			Tasks: []requests.TaskDetail{
				{
					ID:             1,
					Status:         "Doing",
					ProcessPercent: 20,
					Content:        "test get task",
					EndTime:        time.Now().Add(time.Duration(time.Hour * 12)),
					Updated:        time.Hour.Nanoseconds(),
					Created:        time.Hour.Nanoseconds(),
				},
				{
					ID:             2,
					Status:         "Finish",
					ProcessPercent: 100,
					Content:        "test get task 2",
					EndTime:        time.Now().Add(time.Duration(time.Hour * 12)),
					Updated:        time.Hour.Nanoseconds(),
					Created:        time.Hour.Nanoseconds(),
				},
			},
		}
		var mockInput []any
		var mockOutput []any
		mockInput = append(mockInput, ctx, requests.GetTaskRequest{})
		mockOutput = append(mockOutput, reply, nil)
		scripts := []testutil.Scripts[*mock.Mock, database.Tasks]{
			{
				MethodName: "Tasks",
				IsParent:   true,
				ParentReturn: func(m *mock.Mock) database.Tasks {
					return &mocks.Tasks{Mock: m}
				},
			},
			{
				MethodName: "GetTasks",
				Input:      mockInput,
				Output:     mockOutput,
			},
		}
		mockServer := testutil.NewMock[*mock.Mock, database.Tasks](scripts)

		// Act
		err = mockServer.Tasks().ListTasks(ctx)

		// Assert
		assertion.NoError(err)
		expected, err := testexpect.ToString(reply)
		assertion.NoError(err)
		assertion.Equal(http.StatusOK, rec.Code)
		assertion.Equal(expected, rec.Body.String())
	}
}

func (s *Suite) TestUpdateTask() {
	assertion := s.Assertions
	req := func() requests.UpdateTaskRequest {
		return requests.UpdateTaskRequest{
			ID:             1,
			Status:         "finish",
			ProcessPercent: 100,
		}
	}
	var (
		mockRequest, mockReply []any
	)

	// permission denied
	{
		// Arrange
		ctx, rec, err := context.NewTestContext(http.MethodPut, "/api/task", req())
		assertion.NoError(err)
		ctx.Set(authorization.UserPrincipalKey, &authorization.Principal{
			Email:      "user_role@gmail.com",
			Roles:      []roles.Roles{roles.Roles_USER},
			ExpiryTime: time.Now().Add(time.Duration(time.Microsecond * 2)),
		})
		mockServer := testutil.NewMock[*mock.Mock, database.Tasks]([]testutil.Scripts[*mock.Mock, database.Tasks]{})

		// Act
		err = mockServer.Tasks().UpdateTask(ctx)

		// Assert
		assertion.NoError(err)
		expected, err := testexpect.ToString(localErr.Error{
			Code:   localErr.Code_FORBIDDEN,
			Detail: "permission denied",
		})
		assertion.NoError(err)
		assertion.Equal(http.StatusForbidden, rec.Code)
		assertion.Equal(expected, rec.Body.String())
	}

	// missing id
	{
		// Arrange
		request := req()
		request.ID = 0
		ctx, rec, err := context.NewTestContext(http.MethodPut, "/api/task", request)
		assertion.NoError(err)
		ctx.Set(authorization.UserPrincipalKey, &authorization.Principal{
			Email:      "leader_role@gmail.com",
			Roles:      []roles.Roles{roles.Roles_LEADER},
			ExpiryTime: time.Now().Add(time.Duration(time.Microsecond * 2)),
		})
		mockServer := testutil.NewMock[*mock.Mock, database.Tasks]([]testutil.Scripts[*mock.Mock, database.Tasks]{})

		// Act
		err = mockServer.Tasks().UpdateTask(ctx)

		// Assert
		assertion.NoError(err)
		expected, err := testexpect.ToString(localErr.Error{
			Code:   localErr.Code_REQUEST_INSUFFICIENT,
			Detail: "missing field [ID]",
		})
		assertion.NoError(err)
		assertion.Equal(http.StatusBadRequest, rec.Code)
		assertion.Equal(expected, rec.Body.String())
	}

	// update task successfully
	{
		// Arrange
		ctx, rec, err := context.NewTestContext(http.MethodPut, "/api/task", req())
		assertion.NoError(err)
		ctx.Set(authorization.UserPrincipalKey, &authorization.Principal{
			Email:      "leader_role@gmail.com",
			Roles:      []roles.Roles{roles.Roles_LEADER},
			ExpiryTime: time.Now().Add(time.Duration(time.Microsecond * 2)),
		})
		mockRequest = append(mockRequest, ctx, req())
		mockReply = append(mockReply, requests.UpdateTaskReply{
			RowsAffected: 1,
		}, nil)
		mockServer := testutil.NewMock[*mock.Mock, database.Tasks]([]testutil.Scripts[*mock.Mock, database.Tasks]{
			{
				MethodName: "Tasks",
				IsParent:   true,
				ParentReturn: func(m *mock.Mock) database.Tasks {
					return &mocks.Tasks{Mock: m}
				},
			},
			{
				MethodName: "UpdateTask",
				Input:      mockRequest,
				Output:     mockReply,
			},
		})

		// Act
		err = mockServer.Tasks().UpdateTask(ctx)

		// Assert
		assertion.NoError(err)
		expected, err := testexpect.ToString(requests.UpdateTaskReply{
			RowsAffected: 1,
		})
		assertion.NoError(err)
		assertion.Equal(http.StatusOK, rec.Code)
		assertion.Equal(expected, rec.Body.String())
	}
}

func (s *Suite) TestDeleteStacks() {
	assertion := s.Assertions
	req := func() requests.DeleteTaskRequest {
		return requests.DeleteTaskRequest{
			IDs: []int32{1, 2, 3},
		}
	}
	var (
		mockRequest []any
		mockReply   []any
	)

	// permission denied
	{
		// Arrange
		ctx, rec, err := context.NewTestContext(http.MethodDelete, "/api/task", req())
		ctx.Set(authorization.AuthorizationKey, &authorization.Principal{
			Email:      "user_role@gmail.com",
			Roles:      []roles.Roles{roles.Roles_USER},
			ExpiryTime: time.Now().Add(time.Duration(time.Millisecond * 2)),
		})
		assertion.NoError(err)
		mockServer := testutil.NewMock[*mock.Mock, database.Tasks]([]testutil.Scripts[*mock.Mock, database.Tasks]{})

		// Act
		err = mockServer.Tasks().DeleteTasks(ctx)

		// Assert
		assertion.NoError(err)
		expected, err := testexpect.ToString(localErr.Error{
			Code:   localErr.Code_FORBIDDEN,
			Detail: "permission denied",
		})
		assertion.NoError(err)
		assertion.Equal(http.StatusForbidden, rec.Code)
		assertion.Equal(expected, rec.Body.String())
	}
	// missing ids field.
	{
		// Arrange
		ctx, rec, err := context.NewTestContext(http.MethodDelete, "/api/task", requests.DeleteTaskRequest{})
		ctx.Set(authorization.UserPrincipalKey, &authorization.Principal{
			Email:      "leader_role@gmail.com",
			Roles:      []roles.Roles{roles.Roles_LEADER},
			ExpiryTime: time.Now().Add(time.Duration(time.Millisecond * 2)),
		})
		assertion.NoError(err)
		mockServer := testutil.NewMock[*mock.Mock, database.Tasks]([]testutil.Scripts[*mock.Mock, database.Tasks]{})

		// Act
		err = mockServer.Tasks().DeleteTasks(ctx)

		// Assert
		assertion.NoError(err)
		expected, err := testexpect.ToString(localErr.Error{
			Code:   localErr.Code_REQUEST_INSUFFICIENT,
			Detail: "missing field [IDs]",
		})
		assertion.NoError(err)
		assertion.Equal(http.StatusBadRequest, rec.Code)
		assertion.Equal(expected, rec.Body.String())
	}
	// element of field less than expected number.
	{
		// Arrange
		request := req()
		request.IDs = []int32{1, 0}
		ctx, rec, err := context.NewTestContext(http.MethodDelete, "/api/task", request)
		ctx.Set(authorization.UserPrincipalKey, &authorization.Principal{
			Email:      "leader_role@gmail.com",
			Roles:      []roles.Roles{roles.Roles_LEADER},
			ExpiryTime: time.Now().Add(time.Duration(time.Millisecond * 2)),
		})
		assertion.NoError(err)
		mockServer := testutil.NewMock[*mock.Mock, database.Tasks]([]testutil.Scripts[*mock.Mock, database.Tasks]{})

		// Act
		err = mockServer.Tasks().DeleteTasks(ctx)

		// Assert
		assertion.NoError(err)
		expected, err := testexpect.ToString(localErr.Error{
			Code:   localErr.Code_REQUEST_INSUFFICIENT,
			Detail: "field [IDs[1]] should be greater than or equal to 1 but actual got 0",
		})
		assertion.NoError(err)
		assertion.Equal(http.StatusBadRequest, rec.Code)
		assertion.Equal(expected, rec.Body.String())
	}

	// delete stasks successfully.
	{
		// Arrange
		ctx, rec, err := context.NewTestContext(http.MethodDelete, "/api/task", req())
		ctx.Set(authorization.UserPrincipalKey, &authorization.Principal{
			Email:      "leader_role@gmail.com",
			Roles:      []roles.Roles{roles.Roles_LEADER},
			ExpiryTime: time.Now().Add(time.Duration(time.Millisecond * 2)),
		})
		assertion.NoError(err)
		mockRequest = append(mockRequest, ctx, req())
		mockReply = append(mockReply, requests.DeleteTaskReply{
			RowsAffected: 2,
		}, nil)
		mockServer := testutil.NewMock[*mock.Mock, database.Tasks]([]testutil.Scripts[*mock.Mock, database.Tasks]{
			{
				MethodName: "Tasks",
				IsParent:   true,
				ParentReturn: func(m *mock.Mock) database.Tasks {
					return &mocks.Tasks{Mock: m}
				},
			},
			{
				MethodName: "DeleteTasks",
				Input:      mockRequest,
				Output:     mockReply,
			},
		})

		// Act
		err = mockServer.Tasks().DeleteTasks(ctx)

		// Assert
		assertion.NoError(err)
		expected, err := testexpect.ToString(requests.DeleteTaskReply{
			RowsAffected: 2,
		})
		assertion.NoError(err)
		assertion.Equal(http.StatusOK, rec.Code)
		assertion.Equal(expected, rec.Body.String())
	}
}
