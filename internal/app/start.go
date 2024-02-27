package app

import (
	"bufio"
	"github.com/BrownieBrown/pokedex/internal/api/pokemon"
	"github.com/BrownieBrown/pokedex/internal/cli"
	"github.com/BrownieBrown/pokedex/internal/models"
	"github.com/BrownieBrown/pokedex/internal/utils"
	"os"
)

func Start() {
	var cfg models.Config
	cfg.BaseURL = "https://pokeapi.co/api/v2/location/"
	var page models.LocationPage
	var err error

	page, err = pokemon.GetLocations(&cfg.BaseURL)
	if err != nil {
		println(err.Error())
		return
	}
	cfg.Next = page.Next
	cfg.Previous = page.Previous

	scanner := bufio.NewScanner(os.Stdin)

	for {
		print("Pokedex > ")

		scanner.Scan()
		line := utils.CleanInput(scanner.Text())
		commandName := line[0]
		command, exists := cli.GetCommands()[commandName]

		if exists {
			err := command.Callback(&cfg)
			if err != nil {
				println(err.Error())
			}
		} else {
			println("Command not found")
			continue
		}
	}
}
