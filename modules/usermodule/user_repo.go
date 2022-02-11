package usermodule

import (
	"context"

	"gorm.io/gorm"
)

type UserRepoMysql struct {
	db *gorm.DB
}

func NewUserRepoMysql(db *gorm.DB) *UserRepoMysql {
	return &UserRepoMysql{db: db}
}

func (repo *UserRepoMysql) FindByEmail(ctx context.Context, email string) (*User, error) {
	var user User
	if err := repo.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *UserRepoMysql) Create(ctx context.Context, data *UserCreate) error {
	return repo.db.Create(data).Error
}

func (repo *UserRepoMysql) FindByCondition(ctx context.Context, condition map[string]interface{}, moreKeys ...string) (*User, error) {
	db := repo.db
	var user User

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if err := db.Where(condition).First(&user).Error; err != nil {
		// if err == gorm.ErrRecordNotFound {
		// 	return nil, common.RecordNotFound
		// }

		// return nil, common.RecordNotFound

		return nil, err
	}

	return &user, nil
}

func (repo *UserRepoMysql) UpdateById(ctx context.Context, id int, params *UserUpdate) error {
	return repo.db.Where("id = ?", id).Updates(params).Error
}
