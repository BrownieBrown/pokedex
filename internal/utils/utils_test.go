package utils

import (
	"github.com/BrownieBrown/pokedex/internal/models"
	"os"
	"testing"
	"time"
)

func TestCleanInput(t *testing.T) {
	got := CleanInput("Test Input")
	want := []string{"test", "input"}

	assertCleanInput(t, got, want)
}

func TestLoadConfig(t *testing.T) {
	baseURL := "http://pokeapi.co/api/v2"
	cacheInterval := "5s"
	clientTimeout := "10s"

	setEnv(t, "BASE_URL", baseURL)
	setEnv(t, "CACHE_INTERVAL", cacheInterval)
	setEnv(t, "CLIENT_TIMEOUT", clientTimeout)
	defer unsetEnv(t, "BASE_URL")
	defer unsetEnv(t, "CACHE_INTERVAL")
	defer unsetEnv(t, "CLIENT_TIMEOUT")

	got, err := LoadConfig()
	if err != nil {
		t.Fatalf("LoadConfig() returned an error: %v", err)
	}

	want := models.Config{
		BaseURL:       baseURL,
		CacheInterval: 5 * time.Second,
		ClientTimeout: 10 * time.Second,
	}

	assertConfig(t, got, want)
}

func TestLoadConfigMissingBaseURL(t *testing.T) {
	cacheInterval := "5s"
	clientTimeout := "10s"

	setEnv(t, "CACHE_INTERVAL", cacheInterval)
	setEnv(t, "CLIENT_TIMEOUT", clientTimeout)
	defer unsetEnv(t, "CACHE_INTERVAL")
	defer unsetEnv(t, "CLIENT_TIMEOUT")

	_, err := LoadConfig()
	if err == nil {
		t.Error("LoadConfig() did not return an error")
	}
}

func TestLoadConfigMissingCacheInterval(t *testing.T) {
	baseURL := "http://pokeapi.co/api/v2"
	clientTimeout := "10s"

	setEnv(t, "BASE_URL", baseURL)
	setEnv(t, "CLIENT_TIMEOUT", clientTimeout)
	defer unsetEnv(t, "BASE_URL")
	defer unsetEnv(t, "CLIENT_TIMEOUT")

	_, err := LoadConfig()
	if err == nil {
		t.Error("LoadConfig() did not return an error")
	}
}

func TestLoadConfigMissingClientTimeout(t *testing.T) {
	baseURL := "http://pokeapi.co/api/v2"
	cacheInterval := "5s"

	setEnv(t, "BASE_URL", baseURL)
	setEnv(t, "CACHE_INTERVAL", cacheInterval)
	defer unsetEnv(t, "BASE_URL")
	defer unsetEnv(t, "CACHE_INTERVAL")

	_, err := LoadConfig()
	if err == nil {
		t.Error("LoadConfig() did not return an error")
	}
}

func assertCleanInput(t *testing.T, got, want []string) {
	t.Helper()
	if len(got) != len(want) {
		t.Errorf("got %v want %v", got, want)
	}
	for i, v := range want {
		if got[i] != v {
			t.Errorf("got %v want %v", got[i], v)
		}
	}

}
func assertConfig(t *testing.T, got, want models.Config) {
	t.Helper()
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func setEnv(t *testing.T, key, value string) {
	t.Helper()
	err := os.Setenv(key, value)
	if err != nil {
		t.Fatalf("could not set environment variable: %v", err)
	}
}

func unsetEnv(t *testing.T, key string) {
	t.Helper()
	err := os.Unsetenv(key)
	if err != nil {
		t.Fatalf("could not unset environment variable: %v", err)
	}
}
