package config

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func ConfigurationCors() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"*, localhost:5173, localhost:3000"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},  
        AllowHeaders:     []string{"Content-Type", "Authorization"},            
        AllowCredentials: true,
        MaxAge:           12 * time.Hour,
	})
}

