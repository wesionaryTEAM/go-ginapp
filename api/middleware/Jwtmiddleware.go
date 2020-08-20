package middleware

import (
	"fmt"
	"go-ginapp/service"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//AuthorizeJWT  is ... check validity of JWT in incoming request
func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BearerSchema = "Bearer"
		authHeader := c.GetHeader("Authorization")
		tokenString := authHeader[len(BearerSchema):]
		token, err := service.JWTAuthService().ValidateToken(tokenString)
		// if token valid then claim
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			fmt.Println(claims)
		} else {

			c.AbortWithStatus(http.StatusUnauthorized)
			fmt.Println("Malformed Token:", err)
		}

	}
}
