package app

import (
	"bufio"
	"github.com/BrownieBrown/pokedex/internal/api/pokemon"
	"github.com/BrownieBrown/pokedex/internal/cli"
	"github.com/BrownieBrown/pokedex/internal/utils"
	"os"
)

func Start(client *pokemon.Client) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		print("Pokedex > ")

		scanner.Scan()
		line := utils.CleanInput(scanner.Text())
		if len(line) < 1 {
			println("Please enter a command")
			continue
		}
		commandName := line[0]
		command, exists := cli.GetCommands()[commandName]

		var location string
		if len(line) > 1 {
			location = line[1]
		}

		if exists {
			err := command.Callback(client, location)
			if err != nil {
				println(err.Error())
			}
		} else {
			println("Command not found")
			continue
		}
	}
}
