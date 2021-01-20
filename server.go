package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guilherme-brandao/to-go-list/controllers"
	"github.com/guilherme-brandao/to-go-list/middlewares"
	"github.com/guilherme-brandao/to-go-list/services"
)

var (
	taskService  services.TaskService  = services.New()
	loginService services.LoginService = services.NewLoginService()
	jwtService   services.JWTService   = services.NewJWTService()

	taskControllers  controllers.TaskController  = controllers.New(taskService)
	loginControllers controllers.LoginController = controllers.NewLoginController(loginService, jwtService)
)

func main() {

	server := gin.New()

	server.Use(gin.Recovery(), middlewares.Logger())

	server.POST("/login", func(ctx *gin.Context) {
		token := loginControllers.Login(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
		}
	})

	apiRoutes := server.Group("/api", middlewares.AuthorizeJWT())
	{
		apiRoutes.GET("/tasks", func(ctx *gin.Context) {
			ctx.JSON(200, taskControllers.FindAll())
		})

		apiRoutes.POST("/tasks", func(ctx *gin.Context) {
			err := taskControllers.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Task included with success!"})
			}
		})
	}

	server.Run(":8080")
}
