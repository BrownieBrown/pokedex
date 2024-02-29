package cli

import "github.com/BrownieBrown/pokedex/internal/api/pokemon"

func CommandPokedex(client *pokemon.Client, input string) error {
	pkms := client.Pokedex.GetAll()
	println("Your Pokedex:")
	for _, pkm := range pkms {
		println("- " + pkm)
	}
	return nil
}
