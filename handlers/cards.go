package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"trello-cards-creator-oc/models"
)

type Cards struct {
}

func (cd *Cards) CreateCard(c *gin.Context) {

	cardType := models.CardType(c.Param("cardType"))

	if err := cardType.IsValid(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO format EOF errors
	// TODO customize error messages for fields

	switch cardType {
	case models.Issue:
		var issueCard models.IssueCard
		if err := c.ShouldBindJSON(&issueCard); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Println(issueCard)
	case models.Bug:
		var bugCard models.BugCard
		if err := c.ShouldBindJSON(&bugCard); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Println(bugCard)
	case models.Task:
		var taskCard models.TaskCard
		if err := c.ShouldBindJSON(&taskCard); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		fmt.Println(taskCard)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "success",
	})
}
