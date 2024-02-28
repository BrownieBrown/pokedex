package pokemon

import (
	"encoding/json"
	"github.com/BrownieBrown/pokedex/internal/models"
	"io"
	"log"
	"net/http"
)

func (c *Client) GetPokemon(name string) (models.Pokemon, error) {
	url := c.baseURL + "/pokemon/" + name

	if value, ok := c.cache.Get(url); ok {
		var pokemon models.Pokemon
		err := json.Unmarshal(value, &pokemon)
		if err != nil {
			return models.Pokemon{}, err
		}
		return pokemon, err
	}

	res, err := http.Get(url)
	if err != nil {
		return models.Pokemon{}, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("Failed to read response body: %v", err)
		return models.Pokemon{}, err
	}

	var pokemon models.Pokemon
	err = json.Unmarshal(body, &pokemon)
	if err != nil {
		log.Printf("Failed to unmarshal JSON: %v", err)
		return models.Pokemon{}, err
	}

	c.cache.Add(url, body)

	return pokemon, err
}
