package usermodule

import (
	"context"
	"errors"
	"net/mail"
	"strings"
)

const Entity = "User"

type User struct {
	Id             int    `json:"id" gorm:"column:id;"`
	Name           string `json:"name" gorm:"column:name;"`
	Email          string `json:"email" gorm:"column:email;unique"`
	HashedPassword string `json:"-" gorm:"column:hashed_password;"`
	Status         bool   `json:"-" gorm:"column:status"`
	Role           string `json:"-" gorm:"column:role"`
	RefreshToken   string `json:"-" gorm:"column:refresh_token"`
}

func (User) TableName() string {
	return "users"
}

func (u *User) GetUserId() int {
	return u.Id
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) GetRole() string {
	return u.Role
}

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

type UserUpdate struct {
	Name         string `json:"name" gorm:"column:name;"`
	RefreshToken string `json:"-" gorm:"column:refresh_token"`
	Status       bool   `json:"-" gorm:"column:status"`
}

func (UserUpdate) TableName() string {
	return "users"
}

type UserRepo interface {
	FindByEmail(ctx context.Context, email string) (*User, error)
	Create(ctx context.Context, data *UserCreate) error
	FindByCondition(ctx context.Context, condition map[string]interface{}, moreKeys ...string) (*User, error)
	UpdateById(ctx context.Context, id int, params *UserUpdate) error
}
