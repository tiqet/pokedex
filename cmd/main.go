package main

import (
	"bufio"
	"fmt"
	"github.com/tiqet/pokedex/internal/commands"
	"github.com/tiqet/pokedex/internal/config"
	"os"
	"strings"
)

func main() {
	c := config.NewConfig()
	go c.Cache.ReapLoop()

	comds := commands.NewCommands()
	scanner := bufio.NewScanner(os.Stdin)
	var param string
	fmt.Printf("pokedex> ")
	for scanner.Scan() {
		scannedInput := scanner.Text()
		args := strings.Split(scannedInput, " ")
		if len(args) > 1 {
			param = args[1]
		}
		if cmd, ok := comds[args[0]]; ok {
			err := cmd.Callback(c, param)
			if err != nil {
				fmt.Printf("Error executing command '%s': %v\n", cmd.Name, err)
			}
		}
		fmt.Printf("pokedex> ")
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "shouldn't see an error scanning a string")
	}
}
