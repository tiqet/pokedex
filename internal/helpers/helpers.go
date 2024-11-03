package helpers

import (
	"encoding/json"
	"fmt"
	"github.com/tiqet/pokedex/internal/config"
	"github.com/tiqet/pokedex/internal/types"
	"io"
	"net/http"
)

func ProcessLocalArea(c *config.Config, bodyBytes []byte) error {
	var la types.LocationArea
	err := json.Unmarshal(bodyBytes, &la)
	if err != nil {
		return fmt.Errorf("error unmarshalling location area: %v", err)
	}

	for _, result := range la.Results {
		fmt.Printf("%s\n", result.Name)
	}

	if la.Next != "" {
		c.Lau.Next = la.Next
	}
	if la.Previous != "" {
		c.Lau.Previous = la.Previous
	}
	return nil
}

func ProcessLocalAreaPokemon(bodyBytes []byte, areaName string) error {
	var lap types.LocationAreaPokemon
	err := json.Unmarshal(bodyBytes, &lap)
	if err != nil {
		return fmt.Errorf("error unmarshalling location area: %v", err)
	}
	fmt.Printf("Exploring %s...\n", areaName)
	fmt.Println("Found Pokemon:")
	for _, encounter := range lap.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}
	return nil
}

func MakeRequest(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return bodyBytes, nil
}
