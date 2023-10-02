package users

import (
	"github.com/jackc/fake"
	"github.com/labstack/echo"

	localErr "github.com/quocbang/api-server-basic/errors"
	"github.com/quocbang/api-server-basic/impl/requests"
	"github.com/quocbang/api-server-basic/utils/roles"
)

func (s *Suite) TestCreateAccount() {
	assertion := s.Assertions
	ctx := echo.New().AcquireContext()

	// good case
	{
		// arange
		s.ClearData()
		req := requests.CreateAccountRequest{
			Email:    "test_user@gmail.com",
			Password: "password",
			Roles:    []roles.Roles{roles.Roles_USER},
		}

		// act
		reply, err := s.DM.Users().CreateAccount(ctx, req)
		assertion.NoError(err)

		// assertion
		expected := requests.CreateAccountReply{
			RowsAffected: int64(1),
		}
		assertion.Equal(expected, reply)
	}
}

func (s *Suite) TestLogin() {
	assertion := s.Assertions
	ctx := echo.New().AcquireContext()

	// user not found.
	{
		// arrange
		s.ClearData()
		req := requests.LoginRequest{
			Email:    "user_not_found@gmail.com",
			Password: "not found",
		}

		// act
		_, err := s.DM.Users().Login(ctx, req)

		// assertion
		expected := localErr.Error{Code: localErr.Code_ERR_DATA_NOT_FOUND, Detail: "user not found"}
		assertion.Equal(expected, err)
	}

	// wrong password.
	{
		// arrange
		s.ClearData()
		wrongPwdUser := "wrong_password@gmail.com"
		req := requests.LoginRequest{
			Email:    wrongPwdUser,
			Password: "wrong_password",
		}
		_, err := s.DM.Users().CreateAccount(ctx, requests.CreateAccountRequest{
			Email:    wrongPwdUser,
			Password: "password",
		})
		assertion.NoError(err)

		// act
		_, err = s.DM.Users().Login(ctx, req)

		// assertion
		expected := localErr.Error{Code: localErr.Code_WRONG_PASSWORD, Detail: "wrong password"}
		assertion.Equal(expected, err)
	}

	// login successfully.
	{
		// arrange
		s.ClearData()
		goodUser := "good_user@gmail.com"
		goodPwd := "good_password"
		req := requests.LoginRequest{
			Email:    goodUser,
			Password: goodPwd,
		}
		_, err := s.DM.Users().CreateAccount(ctx, requests.CreateAccountRequest{
			Email:    goodUser,
			Password: goodPwd,
		})
		assertion.NoError(err)

		// act
		reply, err := s.DM.Users().Login(ctx, req)

		// assertion
		assertion.NotNil(reply)
		assertion.NoError(err)
	}
}

func (s *Suite) TestDeleteAccounts() {
	assertion := s.Assertions
	ctx := echo.New().AcquireContext()

	// delete accounts successfully.
	{
		// arrange
		s.ClearData()
		emails := []string{"user_1@gmail.com", "user_2@gmail.com"}
		for _, email := range emails {
			createAccount := requests.CreateAccountRequest{
				Email:    email,
				Password: fake.FullName(),
				Roles:    []roles.Roles{roles.Roles_USER},
			}
			_, err := s.DM.Users().CreateAccount(ctx, createAccount)
			assertion.NoError(err)
		}
		deleteRequest := requests.DeleteAccountRequest{
			Emails: emails,
		}

		// act
		reply, err := s.DM.Users().DeleteAccounts(ctx, deleteRequest)

		//assert
		expected := requests.DeleteAccountReply{
			RowsAffected: 2,
		}
		assertion.NoError(err)
		assertion.Equal(expected, reply)
	}
}
