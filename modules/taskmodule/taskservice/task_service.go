package taskservice

import (
	"context"
	"errors"
	"hienviluong125/trello-clone-be/common"
	"hienviluong125/trello-clone-be/errorhandler"
	"hienviluong125/trello-clone-be/modules/taskmodule/taskmodel"
	"hienviluong125/trello-clone-be/modules/taskmodule/taskrepo"
)

type TaskService interface {
	ListByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		filter *taskmodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]taskmodel.Task, error)
	Create(ctx context.Context, taskCreate *taskmodel.TaskCreate) error
	FindByCondition(ctx context.Context, conditions map[string]interface{}) (*taskmodel.Task, error)
	UpdateById(ctx context.Context, taskId int, boardUpdate *taskmodel.TaskUpdate) error
	DeactiveById(ctx context.Context, taskId int) error
	SwapIndexOfTwoTask(ctx context.Context, fromTaskId int, toTaskId int) error
}

type TaskDefaultService struct {
	repo taskrepo.TaskRepo
}

func NewTaskDefaultService(repo taskrepo.TaskRepo) *TaskDefaultService {
	return &TaskDefaultService{repo: repo}
}

func (service *TaskDefaultService) ListByCondition(
	ctx context.Context,
	conditions map[string]interface{},
	filter *taskmodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]taskmodel.Task, error) {
	return service.repo.GetListByCondition(ctx, conditions, filter, paging, moreKeys...)
}

func (service *TaskDefaultService) Create(ctx context.Context, taskCreate *taskmodel.TaskCreate) error {
	if existingTask, _ := service.FindByCondition(ctx, map[string]interface{}{
		"index": taskCreate.Index,
	}); existingTask != nil {
		return errorhandler.ErrRecordExisted("task", errors.New("index of task should be a uniq number"))
	}

	defaultStatus := true
	taskCreate.Status = &defaultStatus
	return service.repo.Create(ctx, taskCreate)
}

func (service *TaskDefaultService) FindByCondition(ctx context.Context, conditions map[string]interface{}) (*taskmodel.Task, error) {
	return service.repo.FindByCondition(ctx, conditions)
}

func (service *TaskDefaultService) UpdateById(ctx context.Context, taskId int, boardUpdate *taskmodel.TaskUpdate) error {
	return service.repo.UpdateById(ctx, taskId, boardUpdate)
}

func (service *TaskDefaultService) DeactiveById(ctx context.Context, taskId int) error {
	status := false
	softDestroyParams := &taskmodel.TaskUpdate{Status: &status}
	return service.repo.UpdateById(ctx, taskId, softDestroyParams)
}

func (service *TaskDefaultService) SwapIndexOfTwoTask(ctx context.Context, fromTaskId int, toTaskId int) error {
	return service.repo.SwapIndexOfTwoTask(ctx, fromTaskId, toTaskId)
}
