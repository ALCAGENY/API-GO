package routes

import (
	"api-go/src/user/infraestructure/http"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.RouterGroup) {
	createController := http.SetUpCreateController()
	authController := http.SetUpAuthController()

	router.POST("/", createController.Run)
	router.POST("/auth", authController.Run)
}