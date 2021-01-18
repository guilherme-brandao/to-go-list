package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/guilherme-brandao/to-go-list/models"
	"github.com/guilherme-brandao/to-go-list/services"
)

type TaskController interface {
	Save(ctx *gin.Context) models.Task
	FindAll() []models.Task
}

type controller struct {
	service service.TaskService
}

func New(service service.TaskService) TaskController {
	return &controller{
		service: service,
	}
}

func (c *controller) Save(ctx *gin.Context) models.Task {
	var task models.Task
	ctx.Bind(&task)
	c.service.Save(task)
	return task
}

func (c *controller) FindAll() []models.Task {
	return c.service.FindAll()
}
