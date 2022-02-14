package userboardrepo

import (
	"context"
	"hienviluong125/trello-clone-be/modules/boardmodule/boardmodel"

	"gorm.io/gorm"
)

type UserBoardRepo interface {
	FindByCondition(ctx context.Context, condition map[string]interface{}) (*boardmodel.UserBoard, error)
	Create(ctx context.Context, data *boardmodel.UserBoard) error
	Destroy(ctx context.Context, data *boardmodel.UserBoard) error
}

type UserBoardRepoMysql struct {
	db *gorm.DB
}

func NewUserBoardRepoMysql(db *gorm.DB) *UserBoardRepoMysql {
	return &UserBoardRepoMysql{db: db}
}

func (repo *UserBoardRepoMysql) FindByCondition(ctx context.Context, condition map[string]interface{}) (*boardmodel.UserBoard, error) {
	var userBoard boardmodel.UserBoard
	if err := repo.db.Where(condition).First(&userBoard).Error; err != nil {
		return nil, err
	}

	return &userBoard, nil
}

func (repo *UserBoardRepoMysql) Create(ctx context.Context, data *boardmodel.UserBoard) error {
	return repo.db.Create(data).Error
}

func (repo *UserBoardRepoMysql) Destroy(ctx context.Context, data *boardmodel.UserBoard) error {
	return repo.db.Where("board_id = ? AND user_id = ?", data.BoardId, data.UserId).Delete(data).Error
}
