package userboardservice

import (
	"context"
	"errors"
	"hienviluong125/trello-clone-be/component"
	"hienviluong125/trello-clone-be/errorhandler"
	"hienviluong125/trello-clone-be/modules/boardmodule/boardmodel"
	"hienviluong125/trello-clone-be/modules/userboardmodule/userboardrepo"
)

type UserBoardService interface {
	Create(ctx context.Context, boardId int, userId int) error
}

type UserBoardDefaultService struct {
	repo       userboardrepo.UserBoardRepo
	appContext component.AppContext
}

func NewUserBoardDefaultService(repo userboardrepo.UserBoardRepo, appContext component.AppContext) *UserBoardDefaultService {
	return &UserBoardDefaultService{repo: repo, appContext: appContext}
}

func (service *UserBoardDefaultService) Create(ctx context.Context, boardId int, userId int) error {
	userBoard, _ := service.repo.FindByCondition(ctx, map[string]interface{}{
		"user_id":  userId,
		"board_id": boardId,
	})

	if userBoard != nil {
		return errorhandler.ErrRecordExisted("user board", errors.New("this user already was a member of this board"))
	}

	userBoardCreate := &boardmodel.UserBoard{
		BoardId: boardId,
		UserId:  userId,
	}

	return service.repo.Create(ctx, userBoardCreate)
}
