package utils

import (
	"fmt"
	"github.com/go-faker/faker/v4"
	"trello-cards-creator-oc/models"
)

type TrelloClient struct {
	trelloApi  *models.TrelloApi
	trelloIds  *models.TrelloIds
	baseUrl    string
	authParams string
}

var cardPosition = "top"

func NewTrelloClient(env *models.Environment) (*TrelloClient, error) {
	return &TrelloClient{
		baseUrl:    env.TrelloApi.TrelloBaseUrl,
		authParams: fmt.Sprintf("key=%s&token=%s", env.TrelloApi.TrelloApiKey, env.TrelloApi.TrelloUserToken),
		trelloApi:  &env.TrelloApi,
		trelloIds:  &env.TrelloIds,
	}, nil
}

type TrelloCard struct {
	Name        string   `json:"name"`
	Description string   `json:"desc"`
	Position    string   `json:"pos"`
	Labels      []string `json:"idLabels"`
	Members     []string `json:"idMembers"`
	List        string   `json:"idList"`
}

func (tc *TrelloClient) CreateIssueCard(issue *models.IssueCard) error {
	trelloCard := TrelloCard{
		Name:     issue.Title,
		List:     tc.trelloIds.TrelloIssuesListId,
		Position: cardPosition,
	}
	fmt.Println(trelloCard)
	return nil
}

func (tc *TrelloClient) CreateBugCard(bug *models.BugCard) error {
	cardName := fmt.Sprintf("bug-%s-%v", faker.Word(), faker.CCNumber())

	trelloCard := TrelloCard{
		Name:        cardName,
		Description: bug.Description,
		List:        tc.trelloIds.TrelloGeneralListId,
		Position:    cardPosition,
		Labels:      []string{tc.trelloIds.TrelloBugLabelId},
	}

	fmt.Println(trelloCard)
	return nil
}

func (tc *TrelloClient) CreateTaskCard(task *models.TaskCard) error {

	var categoryLabel string

	switch task.Category {
	case models.Maintenance:
		categoryLabel = tc.trelloIds.TrelloMaintenanceLabelId
	case models.Research:
		categoryLabel = tc.trelloIds.TrelloResearchLabelId
	case models.Test:
		categoryLabel = tc.trelloIds.TrelloTestLabelId
	}
	trelloCard := TrelloCard{
		Name:     task.Title,
		List:     tc.trelloIds.TrelloGeneralListId,
		Position: cardPosition,
		Labels:   []string{categoryLabel},
	}

	fmt.Println(trelloCard)
	return nil
}
