package cli

import (
	"fmt"
	"github.com/BrownieBrown/pokedex/internal/api/pokemon"
	"github.com/BrownieBrown/pokedex/internal/models"
)

func CommandInspect(client *pokemon.Client, input string) error {
	pkm, exists := client.Pokedex.Get(input)
	if !exists {
		println("You haven't caught " + input + " yet!")
		return nil
	}

	printPokemon(pkm)
	return nil
}

func printPokemon(pkm models.Pokemon) {
	println("Name: " + pkm.Name)
	fmt.Printf("Height: %d m\n", pkm.Height)
	fmt.Printf("Weight: %d kg\n", pkm.Weight)
	println("Stats:")
	for _, stat := range pkm.Stats {
		fmt.Printf("- %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	println("Types:")
	for _, t := range pkm.Types {
		println("- " + t.Type.Name)
	}
}
