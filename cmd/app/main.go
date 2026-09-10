package main

import (
	"fmt"
	"log"

	"github.com/ix1ax/social-network-backend/internal/common/config"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading env files")
	}

	cfg := config.Load()

	fmt.Printf("Server will start on port: %s\n", cfg.ServerPort)
	fmt.Printf("DB Host: %s\n", cfg.DbHost)

}
