package main

import (
	"fmt"
	"github.com/nadirakdag/todo-service/internal/config"
	"github.com/nadirakdag/todo-service/pkg/database"
	"log"
)

func main() {

	log.Println("loading config")
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.NewPostgresDB(cfg.Database)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(db)
}
