package users

import (
	"net/http"
	"time"

	"github.com/stretchr/testify/mock"

	"github.com/quocbang/api-server-basic/database"
	"github.com/quocbang/api-server-basic/database/mocks"
	localErr "github.com/quocbang/api-server-basic/errors"
	"github.com/quocbang/api-server-basic/impl/requests"
	"github.com/quocbang/api-server-basic/impl/servicetest/internal/httptest/context"
	"github.com/quocbang/api-server-basic/impl/servicetest/internal/httptest/testexpect"
	"github.com/quocbang/api-server-basic/impl/servicetest/testutil"
	"github.com/quocbang/api-server-basic/middleware/authorization"
	"github.com/quocbang/api-server-basic/utils/roles"
)

func (s *Suite) TestLogin() {
	assertion := s.Assertions
	goodRequest := func() requests.LoginRequest {
		return requests.LoginRequest{
			Email:    "test@gmail.com",
			Password: "test_password",
		}
	}

	// missing field email
	{
		// Arrange
		misEmailRequest := goodRequest()
		misEmailRequest.Email = ""
		ctx, rec, err := context.NewTestContext(http.MethodPost, "/api/user/login", misEmailRequest)
		assertion.NoError(err)
		mockServer := testutil.NewMock[*mock.Mock, database.Users]([]testutil.Scripts[*mock.Mock, database.Users]{})

		// Act
		err = mockServer.Users().Login(ctx)

		// Assert
		assertion.NoError(err)
		expected, err := testexpect.ToString(localErr.Error{
			Code:   localErr.Code_REQUEST_INSUFFICIENT,
			Detail: "missing field [Email]",
		})
		assertion.NoError(err)
		assertion.Equal(http.StatusBadRequest, rec.Code)
		assertion.Equal(expected, rec.Body.String())
	}
	// missing field password
	{
		// Arrange
		misEmailRequest := goodRequest()
		misEmailRequest.Password = ""
		ctx, rec, err := context.NewTestContext(http.MethodPost, "/api/user/login", misEmailRequest)
		assertion.NoError(err)
		mockServer := testutil.NewMock[*mock.Mock, database.Users]([]testutil.Scripts[*mock.Mock, database.Users]{})

		// Act
		err = mockServer.Users().Login(ctx)

		// Assert
		assertion.NoError(err)
		expected, err := testexpect.ToString(localErr.Error{
			Code:   localErr.Code_REQUEST_INSUFFICIENT,
			Detail: "missing field [Password]",
		})
		assertion.NoError(err)
		assertion.Equal(http.StatusBadRequest, rec.Code)
		assertion.Equal(expected, rec.Body.String())
	}

	// login successfully
	{
		// Arrange
		ctx, rec, err := context.NewTestContext(http.MethodPost, "/api/user/login", goodRequest())
		assertion.NoError(err)
		var mockInput []any
		var mockOutput []any
		mockInput = append(mockInput, ctx, goodRequest())
		mockOutput = append(mockOutput, requests.LoginReply{
			Token: "test_token",
		}, nil)
		mockServer := testutil.NewMock[*mock.Mock, database.Users]([]testutil.Scripts[*mock.Mock, database.Users]{
			{
				MethodName: "Users",
				IsParent:   true,
				ParentReturn: func(m *mock.Mock) database.Users {
					return &mocks.Users{
						Mock: m,
					}
				},
			},
			{
				MethodName: "Login",
				Input:      mockInput,
				Output:     mockOutput,
			},
		})

		// Act
		err = mockServer.Users().Login(ctx)

		// Assert
		assertion.NoError(err)
		expected, err := testexpect.ToString(requests.LoginReply{
			Token: "test_token",
		})
		assertion.NoError(err)
		assertion.Equal(http.StatusOK, rec.Code)
		assertion.Equal(expected, rec.Body.String())
	}
}

func (s *Suite) TestLogout() {
	assertion := s.Assertions

	// token key not found
	{
		// Arrange
		ctx, rec, err := context.NewTestContext(http.MethodPost, "/api/user/logout", nil)
		assertion.NoError(err)
		mockServer := testutil.NewMock[*mock.Mock, database.Users]([]testutil.Scripts[*mock.Mock, database.Users]{})

		// Act
		err = mockServer.Users().Logout(ctx)

		// Assert
		assertion.NoError(err)
		expected, err := testexpect.ToString(localErr.Error{
			Code:   localErr.Code_NONE,
			Detail: "token key not found",
		})
		assertion.NoError(err)
		assertion.Equal(http.StatusNonAuthoritativeInfo, rec.Code)
		assertion.Equal(expected, rec.Body.String())
	}
	// logout successfully
	{
		// Arrange
		expiryTime := time.Now().Add(time.Minute * 2)
		ctx, rec, err := context.NewTestContext(http.MethodPost, "/api/user/logout", nil)
		assertion.NoError(err)
		ctx.Request().Header.Set(authorization.AuthorizationKey, "test_token_key")
		ctx.Set(authorization.UserPrincipalKey, &authorization.Principal{
			Email: "test@gmail.com",
			Roles: []roles.Roles{
				roles.Roles_ADMINISTRATOR,
				roles.Roles_LEADER,
			},
			ExpiryTime: expiryTime,
		})
		mockServer := testutil.NewMock[*mock.Mock, database.Users]([]testutil.Scripts[*mock.Mock, database.Users]{
			{
				MethodName: "Users",
				IsParent:   true,
				ParentReturn: func(m *mock.Mock) database.Users {
					return &mocks.Users{
						Mock: m,
					}
				},
			},
			{
				MethodName: "SignOut",
				Input:      []any{ctx, "test_token_key", time.Until(expiryTime)},
				Output:     []any{nil},
			},
		})

		// Act
		err = mockServer.Users().Logout(ctx)

		// Assert
		assertion.NoError(err)
		expected, err := testexpect.ToString("Logout successfully")
		assertion.NoError(err)
		assertion.Equal(http.StatusOK, rec.Code)
		assertion.Equal(expected, rec.Body.String())
	}
}

func (s *Suite) TestCreateAccount() {
	assertion := s.Assertions
	goodRequest := func() requests.CreateAccountRequest {
		return requests.CreateAccountRequest{
			Email:    "test@gmail.com",
			Password: "test_password",
			Roles:    []roles.Roles{roles.Roles_ADMINISTRATOR, roles.Roles_LEADER, roles.Roles_USER},
		}
	}

	// permission denied
	{
		// Arrange
		ctx, rec, err := context.NewTestContext(http.MethodPost, "/api/user", requests.CreateAccountRequest{})
		assertion.NoError(err)
		ctx.Set(authorization.UserPrincipalKey, &authorization.Principal{
			Email:      "user_role@gmail.com",
			Roles:      []roles.Roles{roles.Roles_USER},
			ExpiryTime: time.Now().Add(time.Minute * 2),
		})
		mockServer := testutil.NewMock[*mock.Mock, database.Users]([]testutil.Scripts[*mock.Mock, database.Users]{})

		// Act
		err = mockServer.Users().CreateAccount(ctx)

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
	// missing email
	{
		// Arrange
		misingEmailRequest := goodRequest()
		misingEmailRequest.Email = ""
		ctx, rec, err := context.NewTestContext(http.MethodPost, "/api/user", misingEmailRequest)
		assertion.NoError(err)
		ctx.Set(authorization.UserPrincipalKey, &authorization.Principal{
			Email:      "admin@gmail.com",
			Roles:      []roles.Roles{roles.Roles_ADMINISTRATOR},
			ExpiryTime: time.Now().Add(time.Minute * 2),
		})
		mockServer := testutil.NewMock[*mock.Mock, database.Users]([]testutil.Scripts[*mock.Mock, database.Users]{})

		// Act
		err = mockServer.Users().CreateAccount(ctx)

		// Assert
		assertion.NoError(err)
		expected, err := testexpect.ToString(localErr.Error{
			Code:   localErr.Code_REQUEST_INSUFFICIENT,
			Detail: "missing field [Email]",
		})
		assertion.NoError(err)
		assertion.Equal(http.StatusBadRequest, rec.Code)
		assertion.Equal(expected, rec.Body.String())
	}
	// missing password
	{
		// Arrange
		misingPasswordRequest := goodRequest()
		misingPasswordRequest.Password = ""
		ctx, rec, err := context.NewTestContext(http.MethodPost, "/api/user", misingPasswordRequest)
		assertion.NoError(err)
		ctx.Set(authorization.UserPrincipalKey, &authorization.Principal{
			Email:      "admin@gmail.com",
			Roles:      []roles.Roles{roles.Roles_ADMINISTRATOR},
			ExpiryTime: time.Now().Add(time.Minute * 2),
		})
		mockServer := testutil.NewMock[*mock.Mock, database.Users]([]testutil.Scripts[*mock.Mock, database.Users]{})

		// Act
		err = mockServer.Users().CreateAccount(ctx)

		// Assert
		assertion.NoError(err)
		expected, err := testexpect.ToString(localErr.Error{
			Code:   localErr.Code_REQUEST_INSUFFICIENT,
			Detail: "missing field [Password]",
		})
		assertion.NoError(err)
		assertion.Equal(http.StatusBadRequest, rec.Code)
		assertion.Equal(expected, rec.Body.String())
	}
	// missing roles
	{
		// Arrange
		misingRoleRequest := goodRequest()
		misingRoleRequest.Roles = nil
		ctx, rec, err := context.NewTestContext(http.MethodPost, "/api/user", misingRoleRequest)
		assertion.NoError(err)
		ctx.Set(authorization.UserPrincipalKey, &authorization.Principal{
			Email:      "admin@gmail.com",
			Roles:      []roles.Roles{roles.Roles_ADMINISTRATOR},
			ExpiryTime: time.Now().Add(time.Minute * 2),
		})
		mockServer := testutil.NewMock[*mock.Mock, database.Users]([]testutil.Scripts[*mock.Mock, database.Users]{})

		// Act
		err = mockServer.Users().CreateAccount(ctx)

		// Assert
		assertion.NoError(err)
		expected, err := testexpect.ToString(localErr.Error{
			Code:   localErr.Code_REQUEST_INSUFFICIENT,
			Detail: "missing field [Roles]",
		})
		assertion.NoError(err)
		assertion.Equal(http.StatusBadRequest, rec.Code)
		assertion.Equal(expected, rec.Body.String())
	}

	// create account successfully
	{
		// Arrange
		ctx, rec, err := context.NewTestContext(http.MethodPost, "/api/user", goodRequest())
		assertion.NoError(err)
		ctx.Set(authorization.UserPrincipalKey, &authorization.Principal{
			Email:      "admin@gmail.com",
			Roles:      []roles.Roles{roles.Roles_ADMINISTRATOR},
			ExpiryTime: time.Now().Add(time.Minute * 2),
		})
		var mockInput []any
		var mockOutput []any
		mockInput = append(mockInput, ctx, goodRequest())
		mockOutput = append(mockOutput, requests.CreateAccountReply{
			RowsAffected: 1,
		}, nil)
		mockServer := testutil.NewMock[*mock.Mock, database.Users]([]testutil.Scripts[*mock.Mock, database.Users]{
			{
				MethodName: "Users",
				IsParent:   true,
				ParentReturn: func(m *mock.Mock) database.Users {
					return &mocks.Users{
						Mock: m,
					}
				},
			},
			{
				MethodName: "CreateAccount",
				Input:      mockInput,
				Output:     mockOutput,
			},
		})

		// Act
		err = mockServer.Users().CreateAccount(ctx)

		// Assert
		assertion.NoError(err)
		expected, err := testexpect.ToString(requests.CreateAccountReply{
			RowsAffected: 1,
		})
		assertion.NoError(err)
		assertion.Equal(http.StatusOK, rec.Code)
		assertion.Equal(expected, rec.Body.String())
	}
}

func (s *Suite) TestDeleteDeleteAcconts() {
	assertion := s.Assertions
	goodRequest := func() requests.DeleteAccountRequest {
		return requests.DeleteAccountRequest{
			Emails: []string{"admin_01@gmail.com", "leader_01@gmail.com", "user_01@gmail.com"},
		}
	}

	// permission denied
	{
		// Arrange
		ctx, rec, err := context.NewTestContext(http.MethodDelete, "/api/user", goodRequest())
		assertion.NoError(err)
		ctx.Set(authorization.UserPrincipalKey, &authorization.Principal{
			Email:      "user_role@gmail.com",
			Roles:      []roles.Roles{roles.Roles_USER},
			ExpiryTime: time.Now().Add(time.Minute * 2),
		})
		mockServer := testutil.NewMock[*mock.Mock, database.Users]([]testutil.Scripts[*mock.Mock, database.Users]{})

		// Act
		err = mockServer.Users().DeleteAccounts(ctx)

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
	// cannot delete yourself
	{
		// Arrange
		request := goodRequest()
		request.Emails = append(request.Emails, "your_self@gmail.com")
		ctx, rec, err := context.NewTestContext(http.MethodDelete, "/api/user", request)
		assertion.NoError(err)
		ctx.Set(authorization.UserPrincipalKey, &authorization.Principal{
			Email:      "your_self@gmail.com",
			Roles:      []roles.Roles{roles.Roles_ADMINISTRATOR},
			ExpiryTime: time.Now().Add(time.Minute * 2),
		})
		mockServer := testutil.NewMock[*mock.Mock, database.Users]([]testutil.Scripts[*mock.Mock, database.Users]{})

		// Act
		err = mockServer.Users().DeleteAccounts(ctx)

		// Assert
		assertion.NoError(err)
		expected, err := testexpect.ToString(localErr.Error{
			Code:   localErr.Code_NOT_SUPPORTED_DELETE_YOUR_SELF,
			Detail: "can not delete your self",
		})
		assertion.NoError(err)
		assertion.Equal(http.StatusBadRequest, rec.Code)
		assertion.Equal(expected, rec.Body.String())
	}
	// missing field emails
	{
		// Arrange
		request := requests.DeleteAccountRequest{}
		ctx, rec, err := context.NewTestContext(http.MethodDelete, "/api/user", request)
		assertion.NoError(err)
		ctx.Set(authorization.UserPrincipalKey, &authorization.Principal{
			Email:      "admin@gmail.com",
			Roles:      []roles.Roles{roles.Roles_ADMINISTRATOR},
			ExpiryTime: time.Now().Add(time.Minute * 2),
		})
		mockServer := testutil.NewMock[*mock.Mock, database.Users]([]testutil.Scripts[*mock.Mock, database.Users]{})

		// Act
		err = mockServer.Users().DeleteAccounts(ctx)

		// Assert
		assertion.NoError(err)
		expected, err := testexpect.ToString(localErr.Error{
			Code:   localErr.Code_REQUEST_INSUFFICIENT,
			Detail: "missing field [Emails]",
		})
		assertion.NoError(err)
		assertion.Equal(http.StatusBadRequest, rec.Code)
		assertion.Equal(expected, rec.Body.String())
	}
	// wrong email format
	{
		// Arrange
		request := requests.DeleteAccountRequest{
			Emails: []string{"wrong_format.gmail.com"},
		}
		ctx, rec, err := context.NewTestContext(http.MethodDelete, "/api/user", request)
		assertion.NoError(err)
		ctx.Set(authorization.UserPrincipalKey, &authorization.Principal{
			Email:      "admin@gmail.com",
			Roles:      []roles.Roles{roles.Roles_ADMINISTRATOR},
			ExpiryTime: time.Now().Add(time.Minute * 2),
		})
		mockServer := testutil.NewMock[*mock.Mock, database.Users]([]testutil.Scripts[*mock.Mock, database.Users]{})

		// Act
		err = mockServer.Users().DeleteAccounts(ctx)

		// Assert
		assertion.NoError(err)
		expected, err := testexpect.ToString(localErr.Error{
			Code:   localErr.Code_REQUEST_INSUFFICIENT,
			Detail: "field [Emails[0]] should be a valid email address example: abc@gmail.com but actual got wrong_format.gmail.com",
		})
		assertion.NoError(err)
		assertion.Equal(http.StatusBadRequest, rec.Code)
		assertion.Equal(expected, rec.Body.String())
	}

	// delete account successfully
	{
		// Arrange
		request := requests.DeleteAccountRequest{
			Emails: []string{"leader@gmail.com"},
		}
		ctx, rec, err := context.NewTestContext(http.MethodDelete, "/api/user", request)
		assertion.NoError(err)
		ctx.Set(authorization.UserPrincipalKey, &authorization.Principal{
			Email:      "admin@gmail.com",
			Roles:      []roles.Roles{roles.Roles_ADMINISTRATOR},
			ExpiryTime: time.Now().Add(time.Minute * 2),
		})
		var mockInput []any
		var mockOutput []any
		mockInput = append(mockInput, ctx, request)
		mockOutput = append(mockOutput, requests.DeleteAccountReply{
			RowsAffected: 1,
		}, nil)
		mockServer := testutil.NewMock[*mock.Mock, database.Users]([]testutil.Scripts[*mock.Mock, database.Users]{
			{
				MethodName: "Users",
				IsParent:   true,
				ParentReturn: func(m *mock.Mock) database.Users {
					return &mocks.Users{
						Mock: m,
					}
				},
			},
			{
				MethodName: "DeleteAccounts",
				Input:      mockInput,
				Output:     mockOutput,
			},
		})

		// Act
		err = mockServer.Users().DeleteAccounts(ctx)

		// Assert
		assertion.NoError(err)
		expected, err := testexpect.ToString(requests.DeleteAccountReply{
			RowsAffected: 1,
		})
		assertion.NoError(err)
		assertion.Equal(http.StatusOK, rec.Code)
		assertion.Equal(expected, rec.Body.String())
	}
}
