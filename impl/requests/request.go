package requests

import "github.com/quocbang/api-server-basic/utils/roles"

type CreateAccountRequest struct {
	Email    string        `validate:"required,email"`
	Password string        `validate:"required"`
	Roles    []roles.Roles `validate:"required"`
}

type CreateAccountReply struct {
	RowsAffected RowsAffected
}

type LoginReply struct {
	Token string
}

type DeleteAccountRequest struct {
	Emails []string `json:"emails" validate:"required,dive,email"`
}

type DeleteAccountReply struct {
	RowsAffected RowsAffected
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
