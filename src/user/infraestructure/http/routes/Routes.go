package routes

import "github.com/gin-gonic/gin"


func UserRoutes(router *gin.RouterGroup){
	AuthRoutes(router)
	ProtectedRoutes(router)
}