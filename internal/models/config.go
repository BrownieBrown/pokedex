package models

type Config struct {
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	BaseURL  string  `json:"base_url"`
}
