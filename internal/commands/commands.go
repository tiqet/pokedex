package commands

import (
	"encoding/json"
	"fmt"
	"github.com/tiqet/pokedex/internal/config"
	"github.com/tiqet/pokedex/internal/helpers"
	"github.com/tiqet/pokedex/internal/types"
	"math/rand"
	"os"
	"strings"
)

type cliCommand struct {
	Name        string
	Description string
	Callback    func(c *config.Config, areaName string) error
}

func NewCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    commandHelp,
		}, "exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    commandExit,
		}, "map": {
			Name:        "map",
			Description: "Explore the next 20 locations",
			Callback:    commandMap,
		}, "mapb": {
			Name:        "mapb",
			Description: "Explore the previous 20 locations",
			Callback:    commandMapb,
		}, "explore": {
			Name:        "explore",
			Description: "List of all the Pokémon in a given area",
			Callback:    commandExplore,
		}, "catch": {
			Name:        "catch",
			Description: "Catching Pokemon adds them to the user's Pokedex",
			Callback:    commandCatch,
		}, "inspect": {
			Name:        "inspect",
			Description: "Show pokemon's details",
			Callback:    commandInspect,
		}, "pokedex": {
			Name:        "pokedex",
			Description: "Print a list of all the names of the Pokemon the user has caught",
			Callback:    commandPokedex,
		},
	}
}

func commandHelp(c *config.Config, s string) error {
	fmt.Printf("\nWelcome to the Pokedex!\nUsage:\n\n")
	for Name, cmd := range NewCommands() {
		fmt.Printf("%s:%s%s\n", Name, strings.Repeat(" ", 8-len(Name)), cmd.Description)
	}
	fmt.Println()
	return nil
}

func commandExit(c *config.Config, s string) error {
	os.Exit(0)
	return nil
}

func commandMap(c *config.Config, s string) error {
	var (
		err       error
		bodyBytes []byte
		ok        bool
	)
	if bodyBytes, ok = c.Cache.Get(c.Lau.Next); ok {
	} else {
		bodyBytes, err = helpers.MakeRequest(c.Lau.Next)
		if err != nil {
			return err
		}
		c.Cache.Add(c.Lau.Next, bodyBytes)
	}
	return helpers.ProcessLocalArea(c, bodyBytes)
}

func commandMapb(c *config.Config, s string) error {
	var (
		err       error
		bodyBytes []byte
		ok        bool
	)
	if bodyBytes, ok = c.Cache.Get(c.Lau.Previous); ok {
	} else {
		bodyBytes, err = helpers.MakeRequest(c.Lau.Previous)
		if err != nil {
			return err
		}
		c.Cache.Add(c.Lau.Previous, bodyBytes)
	}
	return helpers.ProcessLocalArea(c, bodyBytes)
}

func commandExplore(c *config.Config, areaName string) error {
	var (
		err       error
		bodyBytes []byte
		ok        bool
	)
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", areaName)
	if bodyBytes, ok = c.Cache.Get(url); ok {
	} else {
		bodyBytes, err = helpers.MakeRequest(url)
		if err != nil {
			return err
		}
		c.Cache.Add(url, bodyBytes)
	}
	return helpers.ProcessLocalAreaPokemon(bodyBytes, areaName)
}

func commandCatch(c *config.Config, pokName string) error {
	var (
		err       error
		bodyBytes []byte
	)
	fmt.Printf("Throwing a Pokeball at %s...\n", pokName)
	if rand.Intn(c.GuessPoolSize) == 0 {
		fmt.Printf("%s was caught!\n", pokName)
		fmt.Println("You may now inspect it with the inspect command.")
		url := fmt.Sprintf("%s%s", c.PokBaseUrl, pokName)
		bodyBytes, err = helpers.MakeRequest(url)
		if err != nil {
			return err
		}
		var pokemon types.Pokemon
		err = json.Unmarshal(bodyBytes, &pokemon)
		if err != nil {
			return fmt.Errorf("error unmarshalling location area: %v", err)
		}
		c.Pokedex[pokName] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokName)
	}
	return nil
}

func commandInspect(c *config.Config, pokName string) error {
	if pok, ok := c.Pokedex[pokName]; ok {
		fmt.Printf("Name: %s\n", pok.Name)
		fmt.Printf("Height: %d\n", pok.Height)
		fmt.Printf("Weight: %d\n", pok.Weight)
		fmt.Println("Stats:")
		for _, stat := range pok.Stats {
			fmt.Printf("  - %s: %d\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, stat := range pok.Types {
			fmt.Printf("  - %s\n", stat.Type.Name)
		}
	} else {
		fmt.Println("you have not caught that pokemon")
	}
	return nil
}

func commandPokedex(c *config.Config, pokName string) error {
	if len(c.Pokedex) == 0 {
		fmt.Println("You have not caught any Pokémon yet.")
		return nil
	}
	fmt.Println("Your Pokedex:")
	for name, _ := range c.Pokedex {
		fmt.Printf("  - %s\n", name)
	}
	return nil
}
