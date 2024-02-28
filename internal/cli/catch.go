package cli

import (
	"fmt"
	"github.com/BrownieBrown/pokedex/internal/api/pokemon"
	"github.com/BrownieBrown/pokedex/internal/models"
	"math"
	"math/rand"
)

func CommandCatch(client *pokemon.Client, input string) error {
	pkm, err := client.GetPokemon(input)
	if err != nil {
		return err
	}

	fmt.Printf("A wild level %d %s appeared!\n", calculateLevel(pkm.BaseExperience), pkm.Name)

	_, caught := client.Pokedex.Get(pkm.Name)
	if caught {
		println("You've already caught " + pkm.Name + "!")
		return nil
	}

	println("Attempting to catch " + pkm.Name + "...")

	if !catch(pkm) {
		println("Failed to catch " + pkm.Name + "...")
		return nil
	}

	println("Caught " + pkm.Name + "!")

	client.Pokedex.Add(pkm)
	return nil
}

func catch(pkm models.Pokemon) bool {
	experience := pkm.BaseExperience
	level := calculateLevel(experience)

	return attemptCatch(level)
}

func calculateLevel(experience int) int {
	level := int(math.Log(float64(experience+1)) * 10)
	if level < 1 {
		level = 1
	} else if level > 100 {
		level = 100
	}

	return level
}

func attemptCatch(level int) bool {
	maxCatchChance := 95.0
	catchChanceDecreasePerLevel := 0.5
	minCatchChance := 10.0

	catchThreshold := maxCatchChance - (catchChanceDecreasePerLevel * float64(level-1))
	if catchThreshold < minCatchChance {
		catchThreshold = minCatchChance
	}

	roll := rand.Float64() * 100
	fmt.Printf("Roll: %.2f, Needed: %.2f or higher\n", roll, catchThreshold)

	return roll >= catchThreshold
}
