package server

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/guilherme-brandao/to-go-list/controllers"
	"github.com/guilherme-brandao/to-go-list/middlewares"
	"github.com/guilherme-brandao/to-go-list/services"
)

var (
	listService  services.ListService  = services.New()
	loginService services.LoginService = services.NewLoginService()
	jwtService   services.JWTService   = services.NewJWTService()

	taskControllers  controllers.TaskController  = controllers.New(listService)
	loginControllers controllers.LoginController = controllers.NewLoginController(loginService, jwtService)
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	health := new(controllers.HealthController)

	router.GET("/health", health.Status)

	router.POST("/login", func(ctx *gin.Context) {
		token := loginControllers.Login(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
		}
	})

	apiRoutes := router.Group("/api", middlewares.AuthorizeJWT())
	{
		apiRoutes.GET("/lists", func(ctx *gin.Context) {
			ctx.JSON(200, taskControllers.FindAll())
		})

		apiRoutes.POST("/lists", func(ctx *gin.Context) {
			err := taskControllers.NewList(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Task included with success!"})
			}
		})

		apiRoutes.POST("/task/:idList", func(ctx *gin.Context) {
			err := taskControllers.NewTask(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Task included with success!"})
			}
		})

		apiRoutes.GET("/lists/:id", func(ctx *gin.Context) {
			ctx.JSON(200, taskControllers.GetList(ctx))
		})

		apiRoutes.POST("/task-delete/:idList/:idTask", func(ctx *gin.Context) {
			ctx.JSON(200, taskControllers.DeleteTask(ctx))
		})
	}

	return router

}