package userrepo

import (
	"context"
	"hienviluong125/trello-clone-be/modules/usermodule/usermodel"

	"gorm.io/gorm"
)

type UserRepo interface {
	FindByEmail(ctx context.Context, email string) (*usermodel.User, error)
	Create(ctx context.Context, data *usermodel.UserCreate) error
	FindByCondition(ctx context.Context, condition map[string]interface{}, moreKeys ...string) (*usermodel.User, error)
	UpdateById(ctx context.Context, id int, params *usermodel.UserUpdate) error
}

type UserRepoMysql struct {
	db *gorm.DB
}

func NewUserRepoMysql(db *gorm.DB) *UserRepoMysql {
	return &UserRepoMysql{db: db}
}

func (repo *UserRepoMysql) FindByEmail(ctx context.Context, email string) (*usermodel.User, error) {
	var user usermodel.User
	if err := repo.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *UserRepoMysql) Create(ctx context.Context, data *usermodel.UserCreate) error {
	return repo.db.Create(data).Error
}

func (repo *UserRepoMysql) FindByCondition(ctx context.Context, condition map[string]interface{}, moreKeys ...string) (*usermodel.User, error) {
	db := repo.db
	var user usermodel.User

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if err := db.Where(condition).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *UserRepoMysql) UpdateById(ctx context.Context, id int, params *usermodel.UserUpdate) error {
	return repo.db.Where("id = ?", id).Updates(params).Error
}
