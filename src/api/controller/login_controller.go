package controller

import (
	"go-ginapp/auth"
	"go-ginapp/model"
	"go-ginapp/service"
	"log"

	"net/http"

	"github.com/gin-gonic/gin"
)

//Login is ...
func Login(c *gin.Context) {
	var u model.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	//check if the user exist:
	user, err := model.Model.GetUserByEmail(u.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}
	//since after the user logged out, we destroyed that record in the database so that same jwt token can't be used twice. We need to create the token again
	authData, err := model.Model.CreateAuth(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	var authD auth.AuthDetails
	authD.UserId = authData.UserID
	authD.AuthUuid = authData.AuthUUID

	token, loginErr := service.Authorize.SignIn(authD)
	if loginErr != nil {
		c.JSON(http.StatusForbidden, "Please try to login later")
		return
	}
	c.JSON(http.StatusOK, token)
}

//LogOut is ...
func LogOut(c *gin.Context) {
	au, err := auth.ExtractTokenAuth(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	Err := model.Model.DeleteAuth(au)
	if Err != nil {
		log.Println(Err)
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	c.JSON(http.StatusOK, "Successfully logged out")
}
