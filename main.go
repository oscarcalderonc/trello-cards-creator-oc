package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"trello-cards-creator-oc/routes"
)

func main() {
	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "Ok",
		})
	})

	appRoutes := routes.Routes{}

	appRoutes.InitializeRoutes(router)

	err := router.Run(":3030")

	if err != nil {
		fmt.Println("Shutting down...")
		os.Exit(2)
	}
}
