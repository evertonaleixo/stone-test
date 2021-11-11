package middlewares

import (
	"net/http"
	"strconv"

	"github.com/evertonaleixo/stone-test/api"
	"github.com/evertonaleixo/stone-test/services"
	"github.com/gin-gonic/gin"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := services.TokenValid(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}

		accessDetails, err := services.ExtractTokenMetadata(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}

		account, err := api.GetAccountById(strconv.Itoa(int(accessDetails.UserId)))
		if err != nil {
			c.JSON(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}
		c.Set("user", account)
		c.Next()
	}
}
