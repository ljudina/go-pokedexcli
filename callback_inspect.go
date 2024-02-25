package main

import (
	"errors"
	"fmt"
)

func callbackInspect(cfg *config, args ...string) error {
    if len(args) != 1 {
        return errors.New("no pokemon name provided")
    }
    pokemonName := args[0]
    pokemon, ok := cfg.caugthPokemons[pokemonName]
    if !ok {
        return fmt.Errorf("pokemon %s is not caught", pokemonName)
    }
    fmt.Printf("Name: %s\n", pokemon.Name)
    fmt.Printf("Height: %v\n", pokemon.Height)
    fmt.Printf("Weight: %v\n", pokemon.Weight)
    fmt.Printf("Stats:\n")
    for _, stat := range pokemon.Stats {
        fmt.Printf(" - %s: %v\n", stat.Stat.Name, stat.BaseStat)
    }
    fmt.Printf("Types:\n")
    for _, typ := range pokemon.Types {
        fmt.Printf(" - %s\n", typ.Type.Name)
    }
    return nil
}
