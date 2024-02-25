package main

import (
	"errors"
	"fmt"
)
func callbackMap(cfg *config, args ...string) error {

    locationAreas, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAreaURL)
    if err != nil {
        return err
    }
    fmt.Println("Location areas:")
    for _, area := range locationAreas.Results {
        fmt.Printf(" - %s\n", area.Name)
    }
    cfg.nextLocationAreaURL = locationAreas.Next
    cfg.previousLocationAreaURL = locationAreas.Previous
    return nil
}
func callbackMapBack(cfg *config, args ...string) error {
    if cfg.previousLocationAreaURL == nil {
        return errors.New("you're on the first page")
    }
    locationAreas, err := cfg.pokeapiClient.ListLocationAreas(cfg.previousLocationAreaURL)
    if err != nil {
        return err
    }
    fmt.Println("Location areas:")
    for _, area := range locationAreas.Results {
        fmt.Printf(" - %s\n", area.Name)
    }
    cfg.nextLocationAreaURL = locationAreas.Next
    cfg.previousLocationAreaURL = locationAreas.Previous
    return nil
}
