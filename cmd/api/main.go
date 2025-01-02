package main

import (
	"fmt"
	"github.com/nadirakdag/todo-service/internal/config"
	"log"
)

func main() {

	log.Println("loading config")
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cfg)
}
