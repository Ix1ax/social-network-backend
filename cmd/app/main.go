package main

import (
	"fmt"
	"log"

	"github.com/ix1ax/social-network-backend/internal/common/config"
	"github.com/ix1ax/social-network-backend/internal/common/db"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading env files")
	}

	cfg := config.Load()

	database, err := db.NewConnection(cfg)

	if err  != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	if err := db.RunMigrations(cfg); err != nil {
		log.Fatal("Error running migrations: ", err) 
	}

	
	/*
	 * defer откладывает выполнение функции до момента выхода из текущей
	 * т.е как только функция main() завершится, закроется и подключение к БД
	 */
	defer database.Close()

	fmt.Println("Successfully connected to database!")
	fmt.Println("Successfully applied migrations!")
	fmt.Printf("Server will start on port: %s\n", cfg.ServerPort)

}
