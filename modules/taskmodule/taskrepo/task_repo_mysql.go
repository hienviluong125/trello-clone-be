package taskrepo

import (
	"context"
	"hienviluong125/trello-clone-be/common"
	"hienviluong125/trello-clone-be/modules/taskmodule/taskmodel"
	"strings"

	"gorm.io/gorm"
)

type TaskRepo interface {
	GetListByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		filter *taskmodel.Filter,
		paging *common.Paging, moreKeys ...string,
	) ([]taskmodel.Task, error)
	Create(ctx context.Context, data *taskmodel.TaskCreate) error
	FindByCondition(ctx context.Context, condition map[string]interface{}, moreKeys ...string) (*taskmodel.Task, error)
	UpdateById(ctx context.Context, id int, params *taskmodel.TaskUpdate) error
}

type TaskRepoMysql struct {
	db *gorm.DB
}

func NewTaskRepoMysql(db *gorm.DB) *TaskRepoMysql {
	return &TaskRepoMysql{db: db}
}

func (repo *TaskRepoMysql) Create(ctx context.Context, data *taskmodel.TaskCreate) error {
	return repo.db.Create(data).Error
}

func (repo *TaskRepoMysql) FindByCondition(ctx context.Context, condition map[string]interface{}, moreKeys ...string) (*taskmodel.Task, error) {
	db := repo.db
	var task taskmodel.Task

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if err := db.Where("status IS TRUE").Where(condition).First(&task).Error; err != nil {
		return nil, err
	}

	return &task, nil
}

func (repo *TaskRepoMysql) UpdateById(ctx context.Context, id int, params *taskmodel.TaskUpdate) error {
	return repo.db.Where("id = ?", id).Updates(params).Error
}

func (repo *TaskRepoMysql) GetListByCondition(
	ctx context.Context,
	conditions map[string]interface{},
	filter *taskmodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]taskmodel.Task, error) {
	var result []taskmodel.Task
	db := repo.db

	db = db.Table(taskmodel.Task{}.TableName()).Where("status IS TRUE").Where(conditions)

	if v := filter; v != nil {
		if v.Title != "" {
			db = db.Where("LOWER(title) LIKE ?", "%"+strings.ToLower(v.Title)+"%")
		}

		if v.Body != "" {
			db = db.Where("LOWER(body) LIKE ?", "%"+strings.ToLower(v.Body)+"%")
		}

		if v.ListId != 0 {
			db = db.Where("LOWER(list_id) = ?", v.ListId)
		}

		if v.AssigneeId != 0 {
			db = db.Where("LOWER(assignee_id) = ?", v.AssigneeId)
		}
	}

	if err := db.Table(taskmodel.Task{}.TableName()).Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if err := db.
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Order("index asc").
		Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}
