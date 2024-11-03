package config

import (
	"github.com/tiqet/pokedex/internal/pokecache"
	"github.com/tiqet/pokedex/internal/types"
	"time"
)

type Config struct {
	Lau           types.LocAreaUrls
	Cache         pokecache.Cache
	PokBaseUrl    string
	GuessPoolSize int
	Pokedex       map[string]types.Pokemon
}

func NewConfig() *Config {
	return &Config{
		Lau:           types.DefaultLocAreaUrls(),
		Cache:         pokecache.NewCache(5 * time.Second),
		PokBaseUrl:    "https://pokeapi.co/api/v2/pokemon/",
		GuessPoolSize: 10,
		Pokedex:       make(map[string]types.Pokemon),
	}
}
