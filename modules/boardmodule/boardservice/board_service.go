package boardservice

import (
	"context"
	"hienviluong125/trello-clone-be/common"
	"hienviluong125/trello-clone-be/component"
	"hienviluong125/trello-clone-be/modules/boardmodule/boardmodel"
	"hienviluong125/trello-clone-be/modules/boardmodule/boardrepo"
)

type BoardService interface {
	ListByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		filter *boardmodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]boardmodel.Board, error)
	Create(ctx context.Context, boardCreate *boardmodel.BoardCreate) error
	FindByCondition(ctx context.Context, conditions map[string]interface{}) (*boardmodel.Board, error)
	UpdateById(ctx context.Context, boardId int, boardUpdate *boardmodel.BoardUpdate) error
	DeactiveById(ctx context.Context, boardId int) error
	// AddMember(ctx context.Context, boardId int, userId int) error
}

type BoardDefaultService struct {
	repo       boardrepo.BoardRepo
	appContext component.AppContext
}

func NewBoardDefaultService(repo boardrepo.BoardRepo, appContext component.AppContext) *BoardDefaultService {
	return &BoardDefaultService{repo: repo, appContext: appContext}
}

func (service *BoardDefaultService) ListByCondition(
	ctx context.Context,
	conditions map[string]interface{},
	filter *boardmodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]boardmodel.Board, error) {
	return service.repo.GetListByCondition(ctx, conditions, filter, paging, moreKeys...)
}

func (service *BoardDefaultService) Create(ctx context.Context, boardCreate *boardmodel.BoardCreate) error {
	boardCreate.Status = true
	return service.repo.Create(ctx, boardCreate)
}

func (service *BoardDefaultService) FindByCondition(ctx context.Context, conditions map[string]interface{}) (*boardmodel.Board, error) {
	return service.repo.FindByCondition(ctx, conditions)
}

func (service *BoardDefaultService) UpdateById(ctx context.Context, boardId int, boardUpdate *boardmodel.BoardUpdate) error {
	return service.repo.UpdateById(ctx, boardId, boardUpdate)
}

func (service *BoardDefaultService) DeactiveById(ctx context.Context, boardId int) error {
	status := false
	softDestroyParams := &boardmodel.BoardUpdate{Status: &status}
	return service.repo.UpdateById(ctx, boardId, softDestroyParams)
}
