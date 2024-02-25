package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		cleaned := cleanInput(input)
		if len(cleaned) == 0 {
			continue
		}
		commandName := cleaned[0]
        args := []string{}
        if len(cleaned) > 1 {
            args = cleaned[1:]
        }
        availableCommands := getCommands()

        command, ok := availableCommands[commandName]

        if !ok {
			fmt.Println("invalid command")
            continue
        }

        err := command.callback(cfg, args...)
        if err != nil {
            fmt.Println(err)
        }
	}
}

type cliCommand struct {
	name        string
	description string
    callback func(*config, ...string) error
}

func getCommands() map[string] cliCommand {
    return map[string]cliCommand {
        "help": {
            name: "help",
            description: "Prints the help menu",
            callback: callbackHelp,
        },
        "map": {
            name: "map",
            description: "List location areas",
            callback: callbackMap,
        },
        "mapb": {
            name: "mapb",
            description: "List location areas backwards",
            callback: callbackMapBack,
        },
        "explore": {
            name: "explore {location_area}",
            description: "Explore location for pokemons",
            callback: callbackExplore,
        },
        "catch": {
            name: "pokemon {pokemon_name}",
            description: "Try to catch pokemon",
            callback: callbackCatch,
        },
        "inspect": {
            name: "inspect {pokemon_name}",
            description: "Inspect pokemon",
            callback: callbackInspect,
        },
        "exit": {
            name: "exit",
            description: "Turns off the Pokedex",
            callback: callbackExit,
        },
        "pokedex": {
            name: "pokedex",
            description: "View all pokemons in pokedex",
            callback: callbackPokedex,
        },
    }
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
