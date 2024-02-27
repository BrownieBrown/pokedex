package pokemon

import (
	"encoding/json"
	"github.com/BrownieBrown/pokedex/internal/models"
	"io"
	"log"
	"net/http"
)

func GetLocations(url *string) (models.LocationPage, error) {
	if url == nil {
		log.Printf("No URL provided")
		return models.LocationPage{}, nil
	}

	res, err := http.Get(*url)
	if err != nil {
		log.Fatal(err)
		return models.LocationPage{}, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("Failed to read response body: %v", err)
		return models.LocationPage{}, err
	}

	var page models.LocationPage
	err = json.Unmarshal(body, &page)
	if err != nil {
		log.Printf("Failed to unmarshal JSON: %v", err)
		return models.LocationPage{}, err
	}

	return page, err
}
