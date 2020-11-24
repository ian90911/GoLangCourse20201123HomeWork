package main

import (
	"hw/controllers"

	"github.com/gin-gonic/gin"
)

func initRoutes(router *gin.Engine) {

	router.GET("/role", controllers.GetAllRole)

	router.GET("/role/:id", controllers.GetOneRole)

	router.POST("/role", controllers.PostRole)

	router.PUT("/role/:id", controllers.PutRole)

	router.DELETE("/role/:id", controllers.DeleteRole)
}
