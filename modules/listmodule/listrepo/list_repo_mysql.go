package listrepo

import (
	"context"
	"hienviluong125/trello-clone-be/common"
	"hienviluong125/trello-clone-be/modules/listmodule/listmodel"
	"strings"

	"gorm.io/gorm"
)

type ListRepo interface {
	GetListByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		filter *listmodel.Filter,
		paging *common.Paging, moreKeys ...string,
	) ([]listmodel.List, error)
	Create(ctx context.Context, data *listmodel.ListCreate) error
	FindByCondition(ctx context.Context, condition map[string]interface{}, moreKeys ...string) (*listmodel.List, error)
	UpdateById(ctx context.Context, id int, params *listmodel.ListUpdate) error
}

type ListRepoMysql struct {
	db *gorm.DB
}

func NewListRepoMysql(db *gorm.DB) *ListRepoMysql {
	return &ListRepoMysql{db: db}
}

func (repo *ListRepoMysql) Create(ctx context.Context, data *listmodel.ListCreate) error {
	return repo.db.Create(data).Error
}

func (repo *ListRepoMysql) FindByCondition(ctx context.Context, condition map[string]interface{}, moreKeys ...string) (*listmodel.List, error) {
	db := repo.db
	var board listmodel.List

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if err := db.Where("status IS TRUE").Where(condition).First(&board).Error; err != nil {
		return nil, err
	}

	return &board, nil
}

func (repo *ListRepoMysql) UpdateById(ctx context.Context, id int, params *listmodel.ListUpdate) error {
	return repo.db.Where("id = ?", id).Updates(params).Error
}

func (repo *ListRepoMysql) GetListByCondition(
	ctx context.Context,
	conditions map[string]interface{},
	filter *listmodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]listmodel.List, error) {
	var result []listmodel.List
	db := repo.db

	db = db.Table(listmodel.List{}.TableName()).Where("status IS TRUE").Where(conditions)

	if v := filter; v != nil {
		if v.Name != "" {
			db = db.Where("LOWER(name) LIKE ?", "%"+strings.ToLower(v.Name)+"%")
		}
	}

	if err := db.Table(listmodel.List{}.TableName()).Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if err := db.
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Order("index ASC").
		Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}
