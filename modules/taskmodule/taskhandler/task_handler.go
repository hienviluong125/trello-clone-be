package taskhandler

import (
	"hienviluong125/trello-clone-be/common"
	"hienviluong125/trello-clone-be/errorhandler"
	"hienviluong125/trello-clone-be/modules/taskmodule/taskmodel"
	"hienviluong125/trello-clone-be/modules/taskmodule/taskservice"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	service taskservice.TaskService
}

func NewTaskHandler(service taskservice.TaskService) *TaskHandler {
	return &TaskHandler{service: service}
}

// func (handler *TaskHandler) Index(c *gin.Context) {
// 	var filter taskmodel.Filter

// 	if err := c.ShouldBind(&filter); err != nil {
// 		panic(err)
// 	}

// 	var paging common.Paging

// 	if err := c.ShouldBind(&paging); err != nil {
// 		panic(err)
// 	}

// 	paging.FullFill()

// 	currentUser := c.MustGet(common.CurrentUser).(common.Requester)

// 	boards, err := handler.service.ListByCondition(c.Request.Context(), map[string]interface{}{"owner_id": currentUser.GetUserId()}, &filter, &paging, "Owner", "Members")

// 	if err != nil {
// 		panic(err)
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"data":   boards,
// 		"page":   paging,
// 		"filter": filter,
// 	})
// }

func (handler *TaskHandler) Create(c *gin.Context) {
	var taskCreate *taskmodel.TaskCreate

	if err := c.ShouldBind(&taskCreate); err != nil {
		panic(err)
	}

	currentUser := c.MustGet(common.CurrentUser).(common.Requester)
	currentUserId := currentUser.GetUserId()
	taskCreate.ReportedById = &currentUserId

	if err := taskCreate.Validate(); err != nil {
		panic(errorhandler.ErrBadRequest(err))
	}

	if err := handler.service.Create(c.Request.Context(), taskCreate); err != nil {
		panic(err)
	}

	c.Status(http.StatusOK)
}

func (handler *TaskHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		panic(err)
	}

	var taskUpdate *taskmodel.TaskUpdate

	if err := c.ShouldBind(&taskUpdate); err != nil {
		panic(err)
	}

	if err := handler.service.UpdateById(c.Request.Context(), id, taskUpdate); err != nil {
		panic(err)
	}

	c.Status(http.StatusOK)
}

func (handler *TaskHandler) Destroy(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		panic(err)
	}

	if err := handler.service.DeactiveById(c.Request.Context(), id); err != nil {
		panic(err)
	}

	c.Status(http.StatusOK)
}
