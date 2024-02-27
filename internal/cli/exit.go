package cli

import (
	"github.com/BrownieBrown/pokedex/internal/models"
	"os"
)

func CommandExit(cfg *models.Config) error {
	os.Exit(0)
	return nil
}
