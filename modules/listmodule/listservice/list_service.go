package listservice

import (
	"context"
	"errors"
	"hienviluong125/trello-clone-be/common"
	"hienviluong125/trello-clone-be/errorhandler"
	"hienviluong125/trello-clone-be/modules/listmodule/listmodel"
	"hienviluong125/trello-clone-be/modules/listmodule/listrepo"
)

type ListService interface {
	ListByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		filter *listmodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]listmodel.List, error)
	Create(ctx context.Context, boardCreate *listmodel.ListCreate) error
	FindByCondition(ctx context.Context, conditions map[string]interface{}) (*listmodel.List, error)
	UpdateById(ctx context.Context, listId int, listUpdate *listmodel.ListUpdate) error
	DeactiveById(ctx context.Context, listId int) error
}

type ListDefaultService struct {
	repo listrepo.ListRepo
}

func NewListDefaultService(repo listrepo.ListRepo) *ListDefaultService {
	return &ListDefaultService{repo: repo}
}

func (service *ListDefaultService) ListByCondition(
	ctx context.Context,
	conditions map[string]interface{},
	filter *listmodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]listmodel.List, error) {
	return service.repo.GetListByCondition(ctx, conditions, filter, paging, moreKeys...)
}

func (service *ListDefaultService) Create(ctx context.Context, listCreate *listmodel.ListCreate) error {
	existingList, _ := service.repo.FindByCondition(ctx, map[string]interface{}{
		"index":    listCreate.Index,
		"board_id": listCreate.BoardId,
	})

	if existingList != nil {
		return errorhandler.ErrRecordExisted("list", errors.New("index of list should be a unique number"))
	}

	defaultStatus := true
	listCreate.Status = &defaultStatus

	return service.repo.Create(ctx, listCreate)
}

func (service *ListDefaultService) FindByCondition(ctx context.Context, conditions map[string]interface{}) (*listmodel.List, error) {
	return service.repo.FindByCondition(ctx, conditions)
}

func (service *ListDefaultService) UpdateById(ctx context.Context, listId int, listUpdate *listmodel.ListUpdate) error {
	return service.repo.UpdateById(ctx, listId, listUpdate)
}

func (service *ListDefaultService) DeactiveById(ctx context.Context, listId int) error {
	status := false
	softDestroyParams := &listmodel.ListUpdate{Status: &status}
	return service.repo.UpdateById(ctx, listId, softDestroyParams)
}
