package routes

import (
	"github.com/gin-gonic/gin"
	"trello-cards-creator-oc/handlers"
)

type Routes struct {
}

func (rt *Routes) InitializeRoutes(r *gin.Engine) {
	cards(r.Group("/cards"))
}

func cards(g *gin.RouterGroup) {
	cardsHandler := handlers.Cards{}
	g.POST("/:cardType", cardsHandler.CreateCard)
}
