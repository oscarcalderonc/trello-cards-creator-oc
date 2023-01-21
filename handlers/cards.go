package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"trello-cards-creator-oc/models"
	"trello-cards-creator-oc/utils"
)

type Cards struct {
	TrelloClient *utils.TrelloClient
}

func (cd *Cards) CreateCard(c *gin.Context) {

	cardType := models.CardType(c.Param("cardType"))
	var (
		err     error
		newCard *utils.TrelloCard
	)

	if err = cardType.IsValid(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	switch cardType {
	case models.Issue:
		var issueCard models.IssueCard
		if err := c.ShouldBindJSON(&issueCard); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		newCard, err = cd.TrelloClient.CreateIssueCard(&issueCard)
	case models.Bug:
		var bugCard models.BugCard
		if err := c.ShouldBindJSON(&bugCard); err != nil {
			//algo := err.(validator.ValidationErrors)[0].
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		newCard, err = cd.TrelloClient.CreateBugCard(&bugCard)
	case models.Task:
		var taskCard models.TaskCard
		if err := c.ShouldBindJSON(&taskCard); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		newCard, err = cd.TrelloClient.CreateTaskCard(&taskCard)
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"details": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": newCard})
}
