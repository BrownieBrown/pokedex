package models

import "time"

type Config struct {
	BaseURL       string `json:"base_url"`
	CacheInterval time.Duration
	ClientTimeout time.Duration
}
