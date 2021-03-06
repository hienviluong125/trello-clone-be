package boardhandler

import (
	"hienviluong125/trello-clone-be/common"
	"hienviluong125/trello-clone-be/errorhandler"
	"hienviluong125/trello-clone-be/modules/boardmodule/boardmodel"
	boardservice "hienviluong125/trello-clone-be/modules/boardmodule/boardservice"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BoardHandler struct {
	service boardservice.BoardService
}

func NewBoardHandler(service boardservice.BoardService) *BoardHandler {
	return &BoardHandler{service: service}
}

func (handler *BoardHandler) Index(c *gin.Context) {
	var filter boardmodel.Filter

	if err := c.ShouldBind(&filter); err != nil {
		panic(err)
	}

	var paging common.Paging

	if err := c.ShouldBind(&paging); err != nil {
		panic(err)
	}

	paging.FullFill()

	currentUser := c.MustGet(common.CurrentUser).(common.Requester)

	boards, err := handler.service.ListByCondition(c.Request.Context(), map[string]interface{}{"owner_id": currentUser.GetUserId()}, &filter, &paging, "Owner", "Members")

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"data":   boards,
		"page":   paging,
		"filter": filter,
	})
}

func (handler *BoardHandler) Create(c *gin.Context) {
	var boardCreate *boardmodel.BoardCreate

	if err := c.ShouldBind(&boardCreate); err != nil {
		panic(err)
	}

	if err := boardCreate.Validate(); err != nil {
		panic(errorhandler.ErrBadRequest(err))
	}

	currentUser := c.MustGet(common.CurrentUser).(common.Requester)
	boardCreate.OwnerId = currentUser.GetUserId()

	if err := handler.service.Create(c.Request.Context(), boardCreate); err != nil {
		panic(err)
	}

	c.Status(http.StatusOK)
}

func (handler *BoardHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		panic(err)
	}

	var boardUpdate *boardmodel.BoardUpdate

	if err := c.ShouldBind(&boardUpdate); err != nil {
		panic(err)
	}

	if err := handler.AuthorizeBoard(c); err != nil {
		panic(err)
	}

	if err := handler.service.UpdateById(c.Request.Context(), id, boardUpdate); err != nil {
		panic(err)
	}

	c.Status(http.StatusOK)
}

func (handler *BoardHandler) Destroy(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		panic(err)
	}

	if err := handler.AuthorizeBoard(c); err != nil {
		panic(err)
	}

	if err := handler.service.DeactiveById(c.Request.Context(), id); err != nil {
		panic(err)
	}

	c.Status(http.StatusOK)
}

func (handler *BoardHandler) AuthorizeBoard(c *gin.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return err
	}

	currentUser := c.MustGet(common.CurrentUser).(common.Requester)
	board, err := handler.service.FindByCondition(c.Request.Context(), map[string]interface{}{"id": id, "owner_id": currentUser.GetUserId()})

	if err != nil {
		return errorhandler.ErrCannotGetRecord("board", err)
	}

	if board == nil {
		return errorhandler.ErrCannotGetRecord("board", nil)
	}

	return nil
}
