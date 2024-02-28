package utils

import (
	"errors"
	"github.com/BrownieBrown/pokedex/internal/models"
	"os"
	"strings"
	"time"
)

func CleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func LoadConfig() (models.Config, error) {
	var cfg models.Config
	baseURL := os.Getenv("BASE_URL")
	if baseURL == "" {
		return cfg, errors.New("BASE_URL is not set")
	}

	cacheInterval := os.Getenv("CACHE_INTERVAL")
	if cacheInterval == "" {
		return cfg, errors.New("CACHE_INTERVAL is not set")
	}
	clientTimeout := os.Getenv("CLIENT_TIMEOUT")
	if clientTimeout == "" {
		return cfg, errors.New("CLIENT_TIMEOUT is not set")
	}

	parsedCacheInterval, err := time.ParseDuration(cacheInterval)
	if err != nil {
		return cfg, err
	}
	parsedClientTimeout, err := time.ParseDuration(clientTimeout)
	if err != nil {
		return cfg, err
	}

	cfg.BaseURL = baseURL
	cfg.CacheInterval = parsedCacheInterval
	cfg.ClientTimeout = parsedClientTimeout

	return cfg, nil
}
