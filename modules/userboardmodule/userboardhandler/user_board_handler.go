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

	if err := handler.service.Create(c.Request.Context(), boardId, userBoardCreate["user_id"]); err != nil {
		panic(err)
	}

	c.Status(http.StatusOK)
}

func (handler *UserBoardHandler) Destroy(c *gin.Context) {
	boardId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		panic(err)
	}

	userId, err := strconv.Atoi(c.Param("user_id"))

	if err != nil {
		panic(err)
	}

	if err := handler.service.Destroy(c.Request.Context(), boardId, userId); err != nil {
		panic(err)
	}

	c.Status(http.StatusOK)
}
