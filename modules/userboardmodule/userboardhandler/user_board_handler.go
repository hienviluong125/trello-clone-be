package userboardhandler

import (
	"hienviluong125/trello-clone-be/modules/userboardmodule/userboardservice"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserBoardHandler struct {
	service userboardservice.UserBoardService
}

func NewUserBoardHandler(service userboardservice.UserBoardService) *UserBoardHandler {
	return &UserBoardHandler{service: service}
}

func (handler *UserBoardHandler) Create(c *gin.Context) {
	boardId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		panic(err)
	}

	var userBoardCreate map[string]int

	if err := c.ShouldBind(&userBoardCreate); err != nil {
		panic(err)
	}

	// currentUser := c.MustGet(common.CurrentUser).(common.Requester)
	// board, err := handler.service.FindByCondition(c.Request.Context(), map[string]interface{}{"id": id, "owner_id": currentUser.GetUserId()})

	// if err != nil {
	// 	panic(errorhandler.ErrCannotGetRecord("board", err))
	// }

	// if board == nil {
	// 	panic(errorhandler.ErrCannotGetRecord("board", nil))
	// }

	if err := handler.service.Create(c.Request.Context(), boardId, userBoardCreate["user_id"]); err != nil {
		panic(err)
	}

	c.Status(http.StatusOK)
}
