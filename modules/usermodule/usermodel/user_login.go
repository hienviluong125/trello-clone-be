package usermodel

import (
	"errors"
	"net/mail"
	"strings"
)

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (UserLogin) TableName() string {
	return "users"
}

func (user *UserLogin) Validate() error {
	user.Email = strings.TrimSpace(user.Email)
	user.Password = strings.TrimSpace(user.Password)

	if len(user.Email) == 0 {
		return errors.New("email can't be blank")
	}

	if _, err := mail.ParseAddress(user.Email); err != nil {
		return errors.New("invalid email format")
	}

	if len(user.Password) == 0 {
		return errors.New("password can't be blank")
	}

	return nil
}
