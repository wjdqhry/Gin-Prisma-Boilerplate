package server

import (
	"Gin-Prisma-Boilerplate/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.CORSmiddleware())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	v1 := r.Group("/v1")
	{
		userGroup := v1.Group("/user", middleware.JWTmiddleware)
		{

			userGroup.GET("/")
		}
	}

	return r
}
