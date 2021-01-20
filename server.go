package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guilherme-brandao/to-go-list/controllers"
	"github.com/guilherme-brandao/to-go-list/middlewares"
	"github.com/guilherme-brandao/to-go-list/services"
)

var (
	taskService services.TaskService  = services.New()
	taskControllers controllers.TaskController = controllers.New(taskService)
)

func main() {

	server := gin.New()

	server.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth())

	server.GET("/tasks", func(ctx *gin.Context) {
		ctx.JSON(200, taskControllers.FindAll())
	})

	server.POST("/tasks", func(ctx *gin.Context) {
		err := taskControllers.Save(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "Task included with success!"})
		}
	})

	server.Run(":8080")
}
