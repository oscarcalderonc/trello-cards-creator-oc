package models

type TrelloApi struct {
	TrelloBaseUrl   string `env:"TRELLO_BASE_URL,default=https://api.trello.com"`
	TrelloApiKey    string `env:"TRELLO_API_KEY"`
	TrelloUserToken string `env:"TRELLO_USER_TOKEN"`
}

type TrelloIds struct {
	TrelloIssuesListId  string `env:"TRELLO_ISSUES_LIST_ID"`
	TrelloGeneralListId string `env:"TRELLO_GENERAL_LIST_ID"`

	TrelloBoardId string `env:"TRELLO_BOARD_ID"`

	TrelloBugLabelId string `env:"TRELLO_BUG_LABEL_ID"`

	TrelloMaintenanceLabelId string `env:"TRELLO_MAINTENANCE_LABEL_ID"`
	TrelloResearchLabelId    string `env:"TRELLO_RESEARCH_LABEL_ID"`
	TrelloTestLabelId        string `env:"TRELLO_TEST_LABEL_ID"`
}

type Environment struct {
	ServerPort string `env:"SERVER_PORT,default=3030"`

	TrelloApi TrelloApi
	TrelloIds TrelloIds
}
