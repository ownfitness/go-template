package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthCheck struct {
	Status string `json:"status"`
}

func Health(c *gin.Context) {
	c.JSON(http.StatusOK, HealthCheck{
		Status: "ok",
	})
}
