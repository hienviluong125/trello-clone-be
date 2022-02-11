package middleware

import (
	"errors"
	"hienviluong125/trello-clone-be/common"
	"hienviluong125/trello-clone-be/component"
	"hienviluong125/trello-clone-be/errorhandler"

	"github.com/gin-gonic/gin"
)

func Authorize(ac component.AppContext, roles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		currentUser := c.MustGet(common.CurrentUser).(common.Requester)
		currentUserRole := currentUser.GetRole()

		for _, role := range roles {
			if role == currentUserRole {
				c.Next()
				return
			}
		}

		panic(errorhandler.ErrUnauthorized(errors.New("permission denied")))
	}
}
