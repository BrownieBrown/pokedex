package cli

import (
	"github.com/BrownieBrown/pokedex/internal/api/pokemon"
	"github.com/BrownieBrown/pokedex/internal/models"
)

func CommandExplore(client *pokemon.Client, input string) error {
	area, err := client.GetArea(input)
	if err != nil {
		return err
	}

	println("Exploring " + area.Name + "...")
	println("Pokémon in this area:")
	printPokemonsForArea(area)
	return nil
}

func printPokemonsForArea(area models.LocationArea) {
	for _, pokemonEncounter := range area.PokemonEncounters {
		println("- " + pokemonEncounter.Pokemon.Name)
	}
}
