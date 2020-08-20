package controller

import (
	"go-ginapp/dto"
	"go-ginapp/service"

	"github.com/gin-gonic/gin"
)

type loginController struct {
	loginService service.LoginService
	jWtService   service.JWTService
}

//LoginController is ...
type LoginController interface {
	Login(c *gin.Context) string
}

//LoginHandler is ...
func LoginHandler(loginService service.LoginService,
	jWtService service.JWTService) LoginController {
	return &loginController{
		loginService: loginService,
		jWtService:   jWtService,
	}
}

func (controller *loginController) Login(ctx *gin.Context) string {
	var credential dto.LoginCredentials
	err := ctx.ShouldBind(&credential)
	if err != nil {
		return "no data found"
	}
	isUserAuthenticated := controller.loginService.LoginUser(credential.Username, credential.Password)
	if isUserAuthenticated {
		return controller.jWtService.CreateToken(credential.Username, true)

	}
	return ""
}
