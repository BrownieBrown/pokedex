package cli

import (
	"github.com/BrownieBrown/pokedex/internal/api/pokemon"
	"os"
)

func CommandExit(client *pokemon.Client, input string) error {
	os.Exit(0)
	return nil
}
