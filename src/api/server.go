package main

import (
	"go-ginapp/controller"
	"go-ginapp/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	var loginService service.LoginService = service.NewLoginService()
	var jwtService service.JWTService = service.JWTAuthService()
	var loginController controller.LoginController = controller.LoginHandler(loginService, jwtService)

	server := gin.New()

	server.POST("/login", func(c *gin.Context) {
		token := loginController.Login(c)
		if token != "" {
			c.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			c.JSON(http.StatusUnauthorized, nil)
		}
	})
	port := "5000"
	server.Run(":" + port)

}
