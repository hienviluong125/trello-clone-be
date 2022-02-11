package usermodule

import (
	"hienviluong125/trello-clone-be/common"
	"hienviluong125/trello-clone-be/errorhandler"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service UserService
}

func NewUserHandler(service UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (handler *UserHandler) Signup(c *gin.Context) {
	var userCreate UserCreate

	if err := c.ShouldBind(&userCreate); err != nil {
		panic(err)
	}

	if err := handler.service.Signup(c.Request.Context(), &userCreate); err != nil {
		panic(err)
	}

	c.Status(http.StatusOK)
}

func (handler *UserHandler) Login(c *gin.Context) {
	var userLogin UserLogin

	if err := c.ShouldBind(&userLogin); err != nil {
		panic(err)
	}

	accessToken, refreshToken, err := handler.service.Login(c.Request.Context(), &userLogin)

	if err != nil {
		panic(err)
	}

	c.SetCookie(common.RefreshToken, *refreshToken, 60*60*24*7, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"accessToken": accessToken})
}

func (handler *UserHandler) Profile(c *gin.Context) {
	currentUser := c.MustGet(common.CurrentUser).(common.Requester)
	c.JSON(http.StatusOK, currentUser)
}

func (handler *UserHandler) KeepLogin(c *gin.Context) {
	rfTokenCookie, err := c.Cookie(common.RefreshToken)

	if err != nil {
		panic(errorhandler.ErrBadRequest(err))
	}

	accessToken, refreshToken, err := handler.service.RefreshCredentials(c.Request.Context(), rfTokenCookie)

	if err != nil {
		panic(err)
	}

	c.SetCookie(common.RefreshToken, *refreshToken, 60*60*24*7, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"accessToken": accessToken})
}
