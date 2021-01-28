package server

import (
	"github.com/gin-gonic/gin"
	"github.com/guilherme-brandao/to-go-list/database"
	"github.com/guilherme-brandao/to-go-list/middlewares"
)

func Init() {
	database.Init()

	server := NewRouter()

	server.Use(gin.Recovery(), middlewares.Logger())

	server.Run(":8080")
}
