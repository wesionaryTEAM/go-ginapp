package app

import (
	"go-ginapp/controller"
	"go-ginapp/middlewares"
)

func route() {

	router.POST("/todo", middlewares.TokenAuthMiddleware(), controller.CreateTodo)
	router.POST("/login", controller.Login)
	router.POST("/logout", middlewares.TokenAuthMiddleware(), controller.LogOut)
}
