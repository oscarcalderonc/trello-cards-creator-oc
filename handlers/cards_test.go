package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
	"trello-cards-creator-oc/models"
	"trello-cards-creator-oc/utils"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

type FakeTrelloApi struct {
	mock.Mock
}

func (ft *FakeTrelloApi) CreateIssueCard(c *models.IssueCard) (*utils.TrelloCard, error) {
	args := ft.Called(c)
	return args.Get(0).(*utils.TrelloCard), args.Error(1)
}

func (ft *FakeTrelloApi) CreateBugCard(c *models.BugCard) (*utils.TrelloCard, error) {
	args := ft.Called(c)
	return args.Get(0).(*utils.TrelloCard), args.Error(1)
}

func (ft *FakeTrelloApi) CreateTaskCard(c *models.TaskCard) (*utils.TrelloCard, error) {
	args := ft.Called(c)
	return args.Get(0).(*utils.TrelloCard), args.Error(1)
}

type createCardTest struct {
	description  string
	cardType     string
	payload      interface{}
	expectedCode int
}

func TestCreateIssueCard(t *testing.T) {
	testObj := new(FakeTrelloApi)
	cardsHandler := Cards{
		TrelloApi: testObj,
	}

	testObj.On("CreateIssueCard", mock.Anything).Return(&utils.TrelloCard{}, nil)

	g := SetUpRouter()
	g.POST("/cards/:cardType", cardsHandler.CreateCard)

	tests := []createCardTest{
		{
			description:  "invalid card type",
			cardType:     faker.Word(),
			payload:      map[string]interface{}{},
			expectedCode: 400,
		},
		{
			description:  "missing required fields",
			cardType:     "issue",
			payload:      map[string]interface{}{"title": faker.Sentence()},
			expectedCode: 400,
		},
		{
			description:  "successful creation",
			cardType:     "issue",
			payload:      models.IssueCard{Title: faker.Sentence(), Description: faker.Sentence()},
			expectedCode: 200,
		},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			body, _ := json.Marshal(tc.payload)

			req, _ := http.NewRequest("POST", "/cards/"+tc.cardType, bytes.NewReader(body))
			w := httptest.NewRecorder()
			g.ServeHTTP(w, req)

			assert.Equal(t, tc.expectedCode, w.Code)
		})
	}
}

func TestCreateBugCard(t *testing.T) {
	testObj := new(FakeTrelloApi)
	cardsHandler := Cards{
		TrelloApi: testObj,
	}

	testObj.On("CreateBugCard", mock.Anything).Return(&utils.TrelloCard{}, nil)

	g := SetUpRouter()
	g.POST("/cards/:cardType", cardsHandler.CreateCard)

	tests := []createCardTest{
		{
			description:  "invalid card type",
			cardType:     faker.Word(),
			payload:      map[string]interface{}{},
			expectedCode: 400,
		},
		{
			description:  "missing required fields",
			cardType:     "bug",
			payload:      map[string]interface{}{"title": faker.Sentence()},
			expectedCode: 400,
		},
		{
			description:  "successful creation",
			cardType:     "bug",
			payload:      models.BugCard{Description: faker.Sentence()},
			expectedCode: 200,
		},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			body, _ := json.Marshal(tc.payload)

			req, _ := http.NewRequest("POST", "/cards/"+tc.cardType, bytes.NewReader(body))
			w := httptest.NewRecorder()
			g.ServeHTTP(w, req)

			assert.Equal(t, tc.expectedCode, w.Code)
		})
	}
}

func TestCreateTaskCard(t *testing.T) {
	testObj := new(FakeTrelloApi)
	cardsHandler := Cards{
		TrelloApi: testObj,
	}

	testObj.On("CreateTaskCard", mock.Anything).Return(&utils.TrelloCard{}, nil)

	g := SetUpRouter()
	g.POST("/cards/:cardType", cardsHandler.CreateCard)

	tests := []createCardTest{
		{
			description:  "invalid card type",
			cardType:     faker.Word(),
			payload:      map[string]interface{}{},
			expectedCode: 400,
		},
		{
			description:  "missing required fields",
			cardType:     "bug",
			payload:      map[string]interface{}{"title": faker.Sentence()},
			expectedCode: 400,
		},
		{
			description:  "incorrect category",
			cardType:     "bug",
			payload:      map[string]interface{}{"title": faker.Sentence(), "category": faker.Word()},
			expectedCode: 400,
		},
		{
			description:  "successful creation",
			cardType:     "task",
			payload:      map[string]interface{}{"title": faker.Sentence(), "category": "test"},
			expectedCode: 200,
		},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			body, _ := json.Marshal(tc.payload)

			req, _ := http.NewRequest("POST", "/cards/"+tc.cardType, bytes.NewReader(body))
			w := httptest.NewRecorder()
			g.ServeHTTP(w, req)

			assert.Equal(t, tc.expectedCode, w.Code)
		})
	}
}
