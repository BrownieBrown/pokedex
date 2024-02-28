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
	println()

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
	// Normalize the experience to a 1-100 scale.
	level := int(math.Log(float64(experience+1)) * 10)
	// Ensure level falls within the 1 to 100 range.
	if level < 1 {
		level = 1
	} else if level > 100 {
		level = 100
	}
	return level
}

func attemptCatch(level int) bool {
	// Define the maximum catch chance for the lowest level Pokémon.
	maxCatchChance := 90.0 // This means a level 1 Pokémon has a 90% catch chance.
	// Calculate the decrease in catch chance per level.
	catchChanceDecreasePerLevel := 0.9

	// Calculate the threshold that needs to be met or exceeded to catch the Pokémon.
	// The higher the level, the lower the threshold, making it harder to catch.
	catchThreshold := maxCatchChance - (catchChanceDecreasePerLevel * float64(level-1))

	// Generate a random roll to compare with the threshold.
	roll := rand.Float64() * 100 // Generates a number between 0.0 and 100.0

	fmt.Printf("Roll: %.2f, Needed: %.2f or higher\n", roll, catchThreshold)

	// Determine if the catch attempt is successful.
	// A successful catch now requires the roll to be higher than the catchThreshold.
	return roll >= catchThreshold
}
