package models

type CliCommand struct {
	Name        string
	Description string
	Callback    func(cfg *Config) error
}
