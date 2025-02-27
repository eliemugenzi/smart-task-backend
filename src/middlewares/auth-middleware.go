package middlewares

import (
	"net/http"
	"smart-task-backend/src/utils"

	"github.com/gin-gonic/gin"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := utils.GetTokenString(context)

		if tokenString == "" {
			context.AbortWithStatusJSON(
				http.StatusUnauthorized,
				utils.GetResponse(http.StatusUnauthorized, "Unauthorized access", nil))

			return
		}

		token, err := utils.ValidateToken(tokenString)

		if token == nil || !token.Valid {
			context.AbortWithStatusJSON(http.StatusUnauthorized, utils.GetResponse(http.StatusUnauthorized, err.Error(), nil))

			return
		}

		userId, err := utils.GetUserIdFromToken(tokenString)

		if err != nil {
			context.AbortWithStatusJSON(http.StatusUnauthorized, utils.GetResponse(http.StatusUnauthorized, err.Error(), nil))
		}

		context.Set("user_id", userId)
	}
}
