package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

// Represents database server and credentials
type Config struct {
	Server   string
	Database string
}

// ReadMoviesConfig and parse the configuration file
func (c *Config) ReadMoviesConfig() {
	if _, err := toml.DecodeFile("movies_config.toml", &c); err != nil {
		log.Fatal(err)
	}
}

// ReadPeopleConfig and parse the configuration file
func (c *Config) ReadPeopleConfig() {
	if _, err := toml.DecodeFile("people_config.toml", &c); err != nil {
		log.Fatal(err)
	}
}