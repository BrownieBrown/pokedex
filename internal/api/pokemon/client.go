package pokemon

import (
	"github.com/BrownieBrown/pokedex/internal/cache"
	"github.com/BrownieBrown/pokedex/internal/models"
	"net/http"
)

type Client struct {
	cache      *cache.Cache
	httpClient http.Client
	baseURL    string
	Pagination *PaginationManager
	Pokedex    *Pokedex
}

func NewClient(cfg models.Config) *Client {
	c := Client{
		cache.NewCache(cfg.CacheInterval),
		http.Client{
			Timeout: cfg.ClientTimeout,
		},
		cfg.BaseURL,
		NewPaginationManager(),
		NewPokedex(),
	}
	return &c
}
