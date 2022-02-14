package userboardservice

import (
	"context"
	"errors"
	"hienviluong125/trello-clone-be/errorhandler"
	"hienviluong125/trello-clone-be/modules/boardmodule/boardmodel"
	"hienviluong125/trello-clone-be/modules/boardmodule/boardrepo"
	"hienviluong125/trello-clone-be/modules/userboardmodule/userboardrepo"
)

type UserBoardService interface {
	Create(ctx context.Context, boardId int, userId int) error
	Destroy(ctx context.Context, boardId int, userId int) error
}

type UserBoardDefaultService struct {
	repo      userboardrepo.UserBoardRepo
	boardRepo boardrepo.BoardRepo
}

func NewUserBoardDefaultService(repo userboardrepo.UserBoardRepo, boardRepo boardrepo.BoardRepo) *UserBoardDefaultService {
	return &UserBoardDefaultService{repo: repo, boardRepo: boardRepo}
}

func (service *UserBoardDefaultService) Create(ctx context.Context, boardId int, userId int) error {
	if _, err := service.boardRepo.FindByCondition(ctx, map[string]interface{}{
		"id": boardId,
	}); err != nil {
		return errorhandler.ErrCannotGetRecord("board", err)
	}

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

func (service *UserBoardDefaultService) Destroy(ctx context.Context, boardId int, userId int) error {
	userBoard, err := service.repo.FindByCondition(ctx, map[string]interface{}{
		"user_id":  userId,
		"board_id": boardId,
	})

	if err != nil {
		return errorhandler.ErrCannotGetRecord("user board", err)
	}

	if err := service.repo.Destroy(ctx, userBoard); err != nil {
		return errorhandler.ErrInternal(err)
	}

	return nil
}
