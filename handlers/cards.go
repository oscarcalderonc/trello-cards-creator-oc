package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Cards struct {
}

func (cd *Cards) CreateCard(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"result": "success",
	})
}
