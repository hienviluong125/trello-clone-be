package main

import (
	"hienviluong125/trello-clone-be/component"
	"hienviluong125/trello-clone-be/middleware"
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

	if err != nil {
		panic(err)
	}

	log.Println("Connected to database")

	appContext := component.NewAppContext(db, os.Getenv("JWT_SECRET"), 8)

	r := gin.Default()
	r.Use(middleware.Recover(appContext))

	userRepo := userrepo.NewUserRepoMysql(db)
	userService := userservice.NewUserDefaultService(userRepo, appContext)
	userHandler := userhandler.NewUserHandler(userService)

	r.GET("/", Home)
	r.POST("/signup", userHandler.Signup)
	r.POST("/login", userHandler.Login)
	r.POST("/users/keep_login", userHandler.KeepLogin)

	authorized := r.Group("/users")
	authorized.Use(middleware.Authenticate(appContext))
	{
		authorized.GET("/profile", middleware.Authorize(appContext, []string{"member"}), userHandler.Profile)
	}

	r.Run(":8080")
}

func Home(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
