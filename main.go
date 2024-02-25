package main

import (
	"time"
	"github.com/ljudina/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient           pokeapi.Client
	nextLocationAreaURL     *string
	previousLocationAreaURL *string
    caugthPokemons map[string]pokeapi.Pokemon
}


func main() {
    cfg := config {
        pokeapiClient: pokeapi.NewClient(time.Hour),
        caugthPokemons: make(map[string]pokeapi.Pokemon),
    }
	startRepl(&cfg)
}
