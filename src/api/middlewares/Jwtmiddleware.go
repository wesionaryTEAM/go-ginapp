package middlewares

import (
	"go-ginapp/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

//TokenAuthMiddleware is ...
func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := auth.TokenValid(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, "You need to be authorized to access this route")
			c.Abort()
			return
		}
		c.Next()
	}
}
