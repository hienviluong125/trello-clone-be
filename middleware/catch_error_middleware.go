package middleware

import (
	"hienviluong125/trello-clone-be/component"
	"hienviluong125/trello-clone-be/errorhandler"

	"github.com/gin-gonic/gin"
)

func Recover(ac component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.Header("Content-type", "application/json")

				if appErr, ok := err.(*errorhandler.AppError); ok {
					c.AbortWithStatusJSON(appErr.StatusCode, appErr)
					panic(err)
				}

				appErr := errorhandler.ErrInternal(err.(error))
				c.AbortWithStatusJSON(appErr.StatusCode, appErr)
				panic(err)
			}

		}()

		c.Next()
	}
}
