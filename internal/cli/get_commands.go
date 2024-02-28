package cli

import (
	"github.com/BrownieBrown/pokedex/internal/api/pokemon"
)

type Command struct {
	Name        string
	Description string
	Callback    func(client *pokemon.Client, input string) error
}

func GetCommands() map[string]Command {
	return map[string]Command{
		"help": {
			Name:        "help",
			Description: "Displays the help message",
			Callback:    CommandHelp,
		},
		"exit": {
			Name:        "exit",
			Description: "Exits the program",
			Callback:    CommandExit,
		},
		"map": {
			Name:        "map",
			Description: "Displays the next 20 locations",
			Callback:    CommandMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Displays the previous 20 locations",
			Callback:    CommandMapb,
		},
		"explore": {
			Name:        "explore",
			Description: "Display the next given location",
			Callback:    CommandExplore,
		},
		"catch": {
			Name:        "catch",
			Description: "Catch a Pokémon",
			Callback:    CommandCatch,
		},
		"inspect": {
			Name:        "inspect",
			Description: "Inspect a Pokémon",
			Callback:    CommandInspect,
		},
	}
}
