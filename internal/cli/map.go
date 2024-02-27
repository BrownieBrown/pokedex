package cli

import (
	"github.com/BrownieBrown/pokedex/internal/api/pokemon"
	"github.com/BrownieBrown/pokedex/internal/models"
)

func CommandMap(cfg *models.Config) error {
	if cfg.Previous == nil {
		page, err := pokemon.GetLocations(&cfg.BaseURL)
		if err != nil {
			return err
		}
		printLocations(page)
	}

	if cfg.Next == nil {
		println("No next page")
		return nil
	}

	nextPage, err := pokemon.GetLocations(cfg.Next)
	if err != nil {
		return err
	}

	printLocations(nextPage)

	cfg.Next = nextPage.Next
	cfg.Previous = nextPage.Previous

	return nil
}

func CommandMapb(cfg *models.Config) error {
	if cfg.Previous == nil {
		println("No previous page")
		return nil
	}

	previousPage, err := pokemon.GetLocations(cfg.Previous)
	if err != nil {
		return err
	}

	printLocations(previousPage)

	cfg.Previous = previousPage.Previous
	cfg.Next = previousPage.Next

	return nil
}

func printLocations(page models.LocationPage) {
	for _, location := range page.Results {
		println(location.Name)
	}
}
