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
		filter *boardmodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]boardmodel.Board, error)
	Create(ctx context.Context, boardCreate *boardmodel.BoardCreate) error
	// UpdateById(ctx context.Context, boardId int, boardUpdate *boardmodel.BoardUpdate) error
	// DeactiveById(ctx context.Context, boardId int) error
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
	filter *boardmodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]boardmodel.Board, error) {
	return service.repo.GetListByCondition(ctx, nil, filter, paging, moreKeys...)
}

func (service *BoardDefaultService) Create(ctx context.Context, boardCreate *boardmodel.BoardCreate) error {
	boardCreate.Status = true
	return service.repo.Create(ctx, boardCreate)
}
