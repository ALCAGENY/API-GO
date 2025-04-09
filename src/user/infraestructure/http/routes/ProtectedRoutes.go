package routes

import (
	"api-go/src/shared/middlewares"
	"api-go/src/user/infraestructure/http"

	"github.com/gin-gonic/gin"
)

func ProtectedRoutes(router *gin.RouterGroup){
	middlewareAuth := middlewares.JWTAuthMiddleware()
	getByIdController := http.SetUpGetIDController()
	deleteUserController := http.SetUpDeleteController()

	router.Use(middlewareAuth)
	router.GET("/:id", getByIdController.Run)
	router.DELETE("/:id", deleteUserController.Run)
}