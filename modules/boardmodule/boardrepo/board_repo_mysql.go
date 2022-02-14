package boardrepo

import (
	"context"
	"hienviluong125/trello-clone-be/common"
	"hienviluong125/trello-clone-be/modules/boardmodule/boardmodel"
	"strings"

	"gorm.io/gorm"
)

type BoardRepo interface {
	GetListByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		filter *boardmodel.Filter,
		paging *common.Paging, moreKeys ...string,
	) ([]boardmodel.Board, error)
	Create(ctx context.Context, data *boardmodel.BoardCreate) error
	FindByCondition(ctx context.Context, condition map[string]interface{}, moreKeys ...string) (*boardmodel.Board, error)
	UpdateById(ctx context.Context, id int, params *boardmodel.BoardUpdate) error
}

type BoardRepoMysql struct {
	db *gorm.DB
}

func NewBoardRepoMysql(db *gorm.DB) *BoardRepoMysql {
	return &BoardRepoMysql{db: db}
}

func (repo *BoardRepoMysql) Create(ctx context.Context, data *boardmodel.BoardCreate) error {
	return repo.db.Create(data).Error
}

func (repo *BoardRepoMysql) FindByCondition(ctx context.Context, condition map[string]interface{}, moreKeys ...string) (*boardmodel.Board, error) {
	db := repo.db
	var board boardmodel.Board

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if err := db.Where("status IS TRUE").Where(condition).First(&board).Error; err != nil {
		return nil, err
	}

	return &board, nil
}

func (repo *BoardRepoMysql) UpdateById(ctx context.Context, id int, params *boardmodel.BoardUpdate) error {
	return repo.db.Where("id = ?", id).Updates(params).Error
}

func (repo *BoardRepoMysql) GetListByCondition(
	ctx context.Context,
	conditions map[string]interface{},
	filter *boardmodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]boardmodel.Board, error) {
	var result []boardmodel.Board
	db := repo.db

	db = db.Table(boardmodel.Board{}.TableName()).Where("status IS TRUE").Where(conditions)

	if v := filter; v != nil {
		if v.Name != "" {
			db = db.Where("LOWER(name) LIKE ?", "%"+strings.ToLower(v.Name)+"%")
		}
	}

	if err := db.Table(boardmodel.Board{}.TableName()).Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if err := db.
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Order("id desc").
		Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}
