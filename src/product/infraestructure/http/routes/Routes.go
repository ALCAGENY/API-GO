package routes

import (
	"api-go/src/product/infraestructure/http"
	"api-go/src/shared/middlewares"

	"github.com/gin-gonic/gin"
)

func ProductRouter(router *gin.RouterGroup) {
	middlewareAuth := middlewares.JWTAuthMiddleware()
	createController := http.SetUpCreateController()
	getByIdController := http.SetUpGetIDController()
	deleteProductController := http.SetUpDeleteController()

	router.Use(middlewareAuth)

	router.POST("/", createController.Run)
	router.GET("/:id", getByIdController.Run)
	router.DELETE("/:id", deleteProductController.Run)
}