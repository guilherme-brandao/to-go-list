package controllers

import (
	"github.com/guilherme-brandao/to-go-list/models"
	"github.com/guilherme-brandao/to-go-list/services" 
	"github.com/gin-gonic/gin"
)

type TaskController interface {
	NewTask(ctx *gin.Context) error
	NewList(ctx *gin.Context) error
	FindAll() []models.List
	GetList(ctx *gin.Context) models.List
	DeleteTask(ctx *gin.Context) error
	DeleteList(ctx *gin.Context) error
}

type controller struct {
	service services.ListService
}

func New(service services.ListService) TaskController {
	return &controller{
		service: service,
	}
}

func (c *controller) NewTask(ctx *gin.Context) error {
	idList := ctx.Param("idList")
	var task models.Task
	err := ctx.ShouldBindJSON(&task)
	if err != nil {
		return err
	}
	c.service.NewTask(idList, task)
	return nil
}

func (c *controller) NewList(ctx *gin.Context) error {
	var list models.List
	err := ctx.ShouldBindJSON(&list)
	if err != nil {
		return err
	}
	c.service.NewList(list)
	return nil
}

func (c *controller) FindAll() []models.List {
	return c.service.FindAll()
}

func (c *controller) GetList(ctx *gin.Context) models.List {
	idList := ctx.Param("id")

	return c.service.GetList(idList)
}

func (c *controller) DeleteTask(ctx *gin.Context) error {
	idTask := ctx.Param("idTask")
	idList := ctx.Param("idList")

	return c.service.DeleteTask(idList, idTask)
}

func (c *controller) DeleteList(ctx *gin.Context) error {
	idList := ctx.Param("idList")

	return c.service.DeleteList(idList)
}
