package handlers

import "github.com/gin-gonic/gin"

type HealthCheck struct{}

func (hc *HealthCheck) Healthz(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "Ok",
	})
}
