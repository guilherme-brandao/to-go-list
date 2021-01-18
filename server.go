package main

import (
	"github.com/gin-gonic/gin"
	"github.com/guilherme-brandao/to-go-list/controllers"
	"github.com/guilherme-brandao/to-go-list/services"
)

var (
	taskService service.TaskService  = service.New()
	taskControllers controllers.TaskController = controllers.New(taskService)
)

func main() {

	server := gin.Default()

	server.GET("/tasks", func(ctx *gin.Context) {
		ctx.JSON(200, taskControllers.FindAll())
	})

	server.POST("/tasks", func(ctx *gin.Context) {

		ctx.JSON(200, taskControllers.Save(ctx))
	})

	server.Run(":8080")
}
