package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-faker/faker/v4"
	"io"
	"math/rand"
	"net/http"
	"trello-cards-creator-oc/models"
)

type TrelloMember struct {
	Id       string `json:"id"`
	FullName string `json:"fullName"`
	Username string `json:"username"`
}

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
		Name:        issue.Title,
		Description: issue.Description,
		List:        tc.trelloIds.TrelloIssuesListId,
		Position:    cardPosition,
	}

	return tc.createTrelloCard(&trelloCard)
}

func (tc *TrelloClient) CreateBugCard(bug *models.BugCard) error {
	cardName := fmt.Sprintf("bug-%s-%v", faker.Word(), faker.CCNumber())

	members, _ := tc.getBoardMembers()

	randomMemberIdx := rand.Intn(len(members))
	if randomMemberIdx > 0 {
		randomMemberIdx--
	}

	trelloCard := TrelloCard{
		Name:        cardName,
		Description: bug.Description,
		List:        tc.trelloIds.TrelloGeneralListId,
		Position:    cardPosition,
		Members:     []string{members[randomMemberIdx].Id},
		Labels:      []string{tc.trelloIds.TrelloBugLabelId},
	}
	return tc.createTrelloCard(&trelloCard)
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

	return tc.createTrelloCard(&trelloCard)
}

func (tc *TrelloClient) createTrelloCard(card *TrelloCard) error {
	payload, _ := json.Marshal(card)

	url := fmt.Sprintf("%s/1/cards?%s", tc.baseUrl, tc.authParams)

	res, err := http.Post(url, "application/json", bytes.NewReader(payload))
	fmt.Println(res.Body)

	return err
}

func (tc *TrelloClient) getBoardMembers() ([]TrelloMember, error) {

	url := fmt.Sprintf("%s/1/boards/%s/members?%s", tc.baseUrl, tc.trelloIds.TrelloBoardId, tc.authParams)

	res, err := http.Get(url)

	fmt.Println(res.Body)

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)

	var members []TrelloMember

	err = json.Unmarshal(body, &members)

	return members, err
}
