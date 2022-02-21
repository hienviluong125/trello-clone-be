package usermodule

import (
	"hienviluong125/trello-clone-be/component"
	"hienviluong125/trello-clone-be/middleware"
	"hienviluong125/trello-clone-be/modules/usermodule/userhandler"
	"hienviluong125/trello-clone-be/modules/usermodule/userrepo"
	"hienviluong125/trello-clone-be/modules/usermodule/userservice"

	"github.com/gin-gonic/gin"
)

type UserModule struct {
	appContext component.AppContext
}

func NewUserModule(appContext component.AppContext) *UserModule {
	return &UserModule{appContext}
}

func (module *UserModule) RunModule(r *gin.Engine) {
	db := module.appContext.GetDbConnection()
	userRepo := userrepo.NewUserRepoMysql(db)
	userService := userservice.NewUserDefaultService(userRepo, module.appContext)
	userHandler := userhandler.NewUserHandler(userService)

	r.POST("/signup", userHandler.Signup)
	r.POST("/login", userHandler.Login)
	r.POST("/users/keep_login", userHandler.KeepLogin)
	r.GET("/users/profile", middleware.Authenticate(module.appContext), middleware.Authorize(module.appContext, []string{"member", "admin"}), userHandler.Profile)
}
