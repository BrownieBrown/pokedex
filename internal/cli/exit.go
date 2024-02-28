package cli

import (
	"github.com/BrownieBrown/pokedex/internal/api/pokemon"
	"os"
)

func CommandExit(client *pokemon.Client) error {
	os.Exit(0)
	return nil
}
