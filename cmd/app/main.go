package main

import (
	"log"

	"github.com/ix1ax/social-network-backend/internal/app"
)

func main() {

	application, err := app.New()

	if err != nil {
		log.Fatal("failed to initialize application: ", err)
	}

	if err := application.Run(); err != nil {
		log.Fatal("failed to run application: ", err)
	}

}
