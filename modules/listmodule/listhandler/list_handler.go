package listhandler

import (
	"errors"
	"hienviluong125/trello-clone-be/common"
	"hienviluong125/trello-clone-be/errorhandler"
	"hienviluong125/trello-clone-be/modules/boardmodule/boardservice"
	"hienviluong125/trello-clone-be/modules/listmodule/listmodel"
	"hienviluong125/trello-clone-be/modules/listmodule/listservice"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ListHandler struct {
	service      listservice.ListService
	boardService boardservice.BoardService
}

func NewListHandler(service listservice.ListService, boardService boardservice.BoardService) *ListHandler {
	return &ListHandler{service: service, boardService: boardService}
}

func (handler *ListHandler) AuthorizeBoardOwner(c *gin.Context) error {
	boardId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return err
	}

	currentUser := c.MustGet(common.CurrentUser).(common.Requester)

	if _, err := handler.boardService.FindByCondition(c.Request.Context(), map[string]interface{}{
		"id":       boardId,
		"owner_id": currentUser.GetUserId(),
	}); err != nil {
		return err
	}

	return nil
}

func (handler *ListHandler) Index(c *gin.Context) {
	if err := handler.AuthorizeBoardOwner(c); err != nil {
		panic(err)
	}

	boardId, _ := strconv.Atoi(c.Param("id"))

	var filter listmodel.Filter

	if err := c.ShouldBind(&filter); err != nil {
		panic(err)
	}

	var paging common.Paging

	if err := c.ShouldBind(&paging); err != nil {
		panic(err)
	}

	paging.FullFill()

	lists, err := handler.service.ListByCondition(c.Request.Context(), map[string]interface{}{"board_id": boardId}, &filter, &paging, "Tasks")

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"data":   lists,
		"page":   paging,
		"filter": filter,
	})
}

func (handler *ListHandler) Create(c *gin.Context) {
	if err := handler.AuthorizeBoardOwner(c); err != nil {
		panic(err)
	}

	boardId, _ := strconv.Atoi(c.Param("id"))

	var listCreate *listmodel.ListCreate

	if err := c.ShouldBind(&listCreate); err != nil {
		panic(err)
	}

	listCreate.BoardId = &boardId

	if err := listCreate.Validate(); err != nil {
		panic(errorhandler.ErrBadRequest(err))
	}

	if err := handler.service.Create(c.Request.Context(), listCreate); err != nil {
		panic(err)
	}

	c.Status(http.StatusOK)
}

func (handler *ListHandler) Update(c *gin.Context) {
	if err := handler.AuthorizeBoardOwner(c); err != nil {
		panic(err)
	}

	var listUpdate *listmodel.ListUpdate

	if err := c.ShouldBind(&listUpdate); err != nil {
		panic(err)
	}

	listId, err := strconv.Atoi(c.Param("list_id"))

	if err != nil {
		panic(err)
	}

	if err := listUpdate.Validate(); err != nil {
		panic(errorhandler.ErrBadRequest(err))
	}

	if err := handler.service.UpdateById(c.Request.Context(), listId, listUpdate); err != nil {
		panic(err)
	}

	c.Status(http.StatusOK)
}

func (handler *ListHandler) Destroy(c *gin.Context) {
	if err := handler.AuthorizeBoardOwner(c); err != nil {
		panic(err)
	}

	listId, err := strconv.Atoi(c.Param("list_id"))

	if err != nil {
		panic(err)
	}

	if err := handler.service.DeactiveById(c.Request.Context(), listId); err != nil {
		panic(err)
	}

	c.Status(http.StatusOK)
}

func (handler *ListHandler) SwapTwoList(c *gin.Context) {
	if err := handler.AuthorizeBoardOwner(c); err != nil {
		panic(err)
	}

	var swapTwoListParams map[string]*int

	if err := c.ShouldBind(&swapTwoListParams); err != nil {
		panic(err)
	}

	if _, ok := swapTwoListParams["fromListId"]; !ok {
		panic(errorhandler.ErrBadRequest(errors.New("fromListId parameter can't be blank")))
	}

	if _, ok := swapTwoListParams["fromListIndex"]; !ok {
		panic(errorhandler.ErrBadRequest(errors.New("fromLisstIndex parameter can't be blank")))
	}

	if _, ok := swapTwoListParams["toListId"]; !ok {
		panic(errorhandler.ErrBadRequest(errors.New("toListId parameter can't be blank")))
	}

	if _, ok := swapTwoListParams["toListIndex"]; !ok {
		panic(errorhandler.ErrBadRequest(errors.New("toListIndex parameter can't be blank")))
	}

	if err := handler.service.SwapIndexOfTwoList(
		c.Request.Context(),
		*swapTwoListParams["fromListId"],
		*swapTwoListParams["fromListIndex"],
		*swapTwoListParams["toListId"],
		*swapTwoListParams["toListIndex"],
	); err != nil {
		panic(err)
	}

	c.Status(http.StatusOK)
}
