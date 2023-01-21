package routes

import (
	"github.com/gin-gonic/gin"
	"trello-cards-creator-oc/handlers"
	"trello-cards-creator-oc/models"
	"trello-cards-creator-oc/utils"
)

type Routes struct {
	Env *models.Environment
}

func (rt *Routes) InitializeRoutes(r *gin.Engine) {
	cards(r.Group("/cards"), rt.Env)
	healthcheck(r.Group("/health"))
}

func cards(g *gin.RouterGroup, env *models.Environment) {
	trelloClient, _ := utils.NewTrelloClient(env)
	cardsHandler := handlers.Cards{
		TrelloApi: trelloClient,
	}
	g.POST("/:cardType", cardsHandler.CreateCard)
}

func healthcheck(g *gin.RouterGroup) {
	healthcheckHandler := handlers.HealthCheck{}
	g.GET("/", healthcheckHandler.Healthz)
}
