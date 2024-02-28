package cli

import (
	"github.com/BrownieBrown/pokedex/internal/api/pokemon"
	"github.com/BrownieBrown/pokedex/internal/models"
)

func CommandMap(client *pokemon.Client, input string) error {
	locations, err := getNextLocations(client)
	if err != nil {
		return err
	}

	printLocations(locations)
	return nil
}

func CommandMapb(client *pokemon.Client, input string) error {
	locations, err := getPreviousLocations(client)
	if err != nil {
		return err
	}

	printLocations(locations)

	return nil
}

func printLocations(page models.LocationPage) {
	for _, location := range page.Results {
		println(location.Name)
	}
}

func getNextLocations(client *pokemon.Client) (models.LocationPage, error) {
	url := client.Pagination.GetNext()
	locations, err := client.GetLocations(url)
	if err != nil {
		return models.LocationPage{}, err
	}

	client.Pagination.Update(locations.Next, locations.Previous)

	return locations, nil
}

func getPreviousLocations(client *pokemon.Client) (models.LocationPage, error) {
	url := client.Pagination.GetPrevious()
	locations, err := client.GetLocations(url)
	if err != nil {
		return models.LocationPage{}, err
	}

	client.Pagination.Update(locations.Next, locations.Previous)

	return locations, nil
}
