package usermodel

import (
	"errors"
	"strings"
)

type UserCreate struct {
	Email    string `json:"email" gorm:"column:email;"`
	Password string `json:"password" gorm:"column:hashed_password;"`
	Name     string `json:"name" gorm:"column:name;"`
	Status   bool   `json:"-" gorm:"column:status"`
	Role     string `json:"-" gorm:"column:role"`
}

func (UserCreate) TableName() string {
	return "users"
}

func (user *UserCreate) Validate() error {
	user.Email = strings.TrimSpace(user.Email)
	user.Password = strings.TrimSpace(user.Password)

	if len(user.Email) == 0 {
		return errors.New("email can't be blank")
	}

	if len(user.Password) == 0 {
		return errors.New("password can't be blank")
	}

	return nil
}
