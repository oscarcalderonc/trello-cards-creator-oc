package main

import (
	"fmt"
	"github.com/Netflix/go-env"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"trello-cards-creator-oc/models"
	"trello-cards-creator-oc/routes"
)

func main() {
	var environment models.Environment
	_, err := env.UnmarshalFromEnviron(&environment)
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()

	appRoutes := routes.Routes{}

	appRoutes.InitializeRoutes(router)

	err = router.Run(":" + environment.ServerPort)

	if err != nil {
		fmt.Println("Shutting down...")
		os.Exit(2)
	}
}
