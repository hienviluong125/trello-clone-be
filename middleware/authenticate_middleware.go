package middleware

import (
	"errors"
	"fmt"
	"hienviluong125/trello-clone-be/common"
	"hienviluong125/trello-clone-be/component"
	"hienviluong125/trello-clone-be/errorhandler"
	"hienviluong125/trello-clone-be/modules/usermodule"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Authenticate(ac component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		bearToken := c.Request.Header["Authorization"]

		if len(bearToken) < 1 {
			panic(errorhandler.ErrUnauthorized(errors.New("invalid credentials")))
		}

		bearTokenArr := strings.Split(bearToken[0], " ")

		if len(bearTokenArr) != 2 {
			panic(errorhandler.ErrUnauthorized(errors.New("invalid credentials")))
		}

		if bearTokenArr[0] != "Bearer" {
			panic(errorhandler.ErrUnauthorized(errors.New("invalid credentials")))
		}

		var defaultTokenProvider common.TokenProvider = common.NewDefaultTokenProvider(ac.GetSecretKey())
		jwtToken, err := defaultTokenProvider.ParseAccessToken(bearTokenArr[1])

		if err != nil {
			panic(errorhandler.ErrUnauthorized(err))
		}

		claims, ok := jwtToken.Claims.(jwt.MapClaims)

		if !ok {
			panic(errorhandler.ErrUnauthorized(errors.New("invalid credentials")))
		}

		userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)

		if err != nil {
			panic(errorhandler.ErrUnauthorized(err))
		}

		userRepo := usermodule.NewUserRepoMysql(ac.GetDbConnection())
		authenticatedUser, err := userRepo.FindByCondition(c.Request.Context(), map[string]interface{}{"id": userId})

		if err != nil {
			panic(errorhandler.ErrUnauthorized(err))
		}

		c.Set(common.CurrentUser, authenticatedUser)
		c.Next()
	}
}
