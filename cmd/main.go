package main

import (
	"log"

	"github.com/pr1sonmike/project-layout/internal/config"
)

func main() {
	cfg, err := config.NewConfig("config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(cfg)
}
