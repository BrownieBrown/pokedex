package cli

import "github.com/BrownieBrown/pokedex/internal/models"

func CommandHelp(cfg *models.Config) error {
	println("Welcome to Pokedex!")
	println("Usage:")
	println()
	commands := GetCommands()
	for _, command := range commands {
		println(command.Name + ": " + command.Description)
	}
	println()

	return nil
}
