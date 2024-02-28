package pokemon

import (
	"encoding/json"
	"github.com/BrownieBrown/pokedex/internal/models"
	"io"
	"log"
	"net/http"
)

func (c *Client) GetLocations(pageURL *string) (models.LocationPage, error) {
	url := c.baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if value, ok := c.cache.Get(url); ok {
		var page models.LocationPage
		err := json.Unmarshal(value, &page)
		if err != nil {
			return models.LocationPage{}, err
		}
		return page, err
	}

	res, err := http.Get(url)
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

	c.cache.Add(url, body)

	return page, err
}

func (c *Client) GetArea(location string) (models.LocationArea, error) {
	url := c.baseURL + "/location-area/" + location
	if value, ok := c.cache.Get(url); ok {
		var area models.LocationArea
		err := json.Unmarshal(value, &area)
		if err != nil {
			return models.LocationArea{}, err
		}
		return area, err
	}

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return models.LocationArea{}, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("Failed to read response body: %v", err)
		return models.LocationArea{}, err
	}

	var area models.LocationArea
	err = json.Unmarshal(body, &area)
	if err != nil {
		log.Printf("Failed to unmarshal JSON: %v", err)
		return models.LocationArea{}, err
	}

	c.cache.Add(url, body)

	return area, err
}
