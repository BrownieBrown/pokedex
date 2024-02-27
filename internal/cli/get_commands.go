package cli

import "github.com/BrownieBrown/pokedex/internal/models"

func GetCommands() map[string]models.CliCommand {
	return map[string]models.CliCommand{
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
	}
}
