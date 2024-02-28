package main

import (
	"github.com/BrownieBrown/pokedex/internal/api/pokemon"
	"github.com/BrownieBrown/pokedex/internal/app"
	"github.com/BrownieBrown/pokedex/internal/utils"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg, err := utils.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	pokeClient := pokemon.NewClient(cfg)

	app.Start(pokeClient)
}
