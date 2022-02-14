package main

import (
	"hienviluong125/trello-clone-be/component"
	"hienviluong125/trello-clone-be/middleware"
	"hienviluong125/trello-clone-be/modules/boardmodule/boardhandler"
	"hienviluong125/trello-clone-be/modules/boardmodule/boardmodel"
	"hienviluong125/trello-clone-be/modules/boardmodule/boardrepo"
	"hienviluong125/trello-clone-be/modules/boardmodule/boardservice"
	"hienviluong125/trello-clone-be/modules/userboardmodule/userboardhandler"
	"hienviluong125/trello-clone-be/modules/userboardmodule/userboardrepo"
	"hienviluong125/trello-clone-be/modules/userboardmodule/userboardservice"
	"hienviluong125/trello-clone-be/modules/usermodule/userhandler"
	"hienviluong125/trello-clone-be/modules/usermodule/usermodel"
	"hienviluong125/trello-clone-be/modules/usermodule/userrepo"
	"hienviluong125/trello-clone-be/modules/usermodule/userservice"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&usermodel.User{})
	db.AutoMigrate(&boardmodel.Board{})
	db.Debug()

	if err != nil {
		panic(err)
	}

	log.Println("Connected to database")

	appContext := component.NewAppContext(db, os.Getenv("JWT_SECRET"), 8)

	r := gin.Default()
	r.Use(middleware.Recover(appContext))
	r.GET("/", Home)
	runService(r, appContext)

	r.Run(":8080")
}

func Home(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func runService(r *gin.Engine, appContext component.AppContext) {
	db := appContext.GetDbConnection()

	// user resources
	// another way to init a module
	// userModule := usermodule.NewUserModule(appContext)
	// userModule.RunModule(r)
	userRepo := userrepo.NewUserRepoMysql(db)
	userService := userservice.NewUserDefaultService(userRepo, appContext)
	userHandler := userhandler.NewUserHandler(userService)
	r.POST("/signup", userHandler.Signup)
	r.POST("/login", userHandler.Login)
	r.POST("/users/keep_login", userHandler.KeepLogin)
	r.GET("/users/profile", middleware.Authenticate(appContext), middleware.Authorize(appContext, []string{"member", "admin"}), userHandler.Profile)

	// board resources
	boardRepo := boardrepo.NewBoardRepoMysql(db)
	boardService := boardservice.NewBoardDefaultService(boardRepo, appContext)
	boardHandler := boardhandler.NewBoardHandler(boardService)
	r.GET("/boards", middleware.Authenticate(appContext), middleware.Authorize(appContext, []string{"member", "admin"}), boardHandler.Index)
	r.POST("/boards", middleware.Authenticate(appContext), middleware.Authorize(appContext, []string{"member", "admin"}), boardHandler.Create)
	r.PUT("/boards/:id", middleware.Authenticate(appContext), middleware.Authorize(appContext, []string{"member", "admin"}), boardHandler.Update)
	r.DELETE("/boards/:id", middleware.Authenticate(appContext), middleware.Authorize(appContext, []string{"member", "admin"}), boardHandler.Destroy)

	// user board resources
	userBoardRepo := userboardrepo.NewUserBoardRepoMysql(db)
	userBoardService := userboardservice.NewUserBoardDefaultService(userBoardRepo, appContext)
	userBoardHandler := userboardhandler.NewUserBoardHandler(userBoardService)
	r.POST("/boards/:id/user_boards", middleware.Authenticate(appContext), middleware.Authorize(appContext, []string{"member", "admin"}), userBoardHandler.Create)
}
