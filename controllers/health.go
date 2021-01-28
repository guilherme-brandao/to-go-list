package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HealthController struct{}

func (health HealthController) Status(c *gin.Context) {
	c.String(http.StatusOK, "Up!")
}