package main

import (
	"errors"
	"fmt"
)

func callbackPokedex(cfg *config, args ...string) error {
    fmt.Println("Pokemon in Pokedex:")
    if len(cfg.caugthPokemons) == 0 {
        return errors.New("no pokemon in pokedex")
    }
    for _,pokemon := range cfg.caugthPokemons {
        fmt.Printf(" - %s\n", pokemon.Name)
    }
    return nil
}
