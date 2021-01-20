package controllers

import (
	"github.com/guilherme-brandao/to-go-list/models"
	"github.com/guilherme-brandao/to-go-list/services" 
	"github.com/gin-gonic/gin"
)

type TaskController interface {
	Save(ctx *gin.Context) error
	FindAll() []models.Task
}

type controller struct {
	service services.TaskService
}

func New(service services.TaskService) TaskController {
	return &controller{
		service: service,
	}
}

func (c *controller) Save(ctx *gin.Context) error {
	var task models.Task
	err := ctx.ShouldBindJSON(&task)
	if err != nil {
		return err
	}
	c.service.Save(task)
	return nil
}

func (c *controller) FindAll() []models.Task {
	return c.service.FindAll()
}
