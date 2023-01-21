package routes

import (
	"github.com/gin-gonic/gin"
	"trello-cards-creator-oc/handlers"
)

type Routes struct {
}

func (rt *Routes) InitializeRoutes(r *gin.Engine) {
	cards(r.Group("/cards"))
	healthcheck(r.Group("/health"))
}

func cards(g *gin.RouterGroup) {
	cardsHandler := handlers.Cards{}
	g.POST("/:cardType", cardsHandler.CreateCard)
}

func healthcheck(g *gin.RouterGroup) {
	healthcheckHandler := handlers.HealthCheck{}
	g.GET("/", healthcheckHandler.Healthz)
}
