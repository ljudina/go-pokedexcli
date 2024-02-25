
package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error){
    endpoint := "/pokemon/" + pokemonName
    fullUrl := baseURL + endpoint

    cached, ok := c.cache.Get(fullUrl)
    if ok {
        fmt.Println("cache hit")
        pokemon := Pokemon{}
        err := json.Unmarshal(cached, &pokemon)
        if err != nil {
            return pokemon, err
        }
        return pokemon, nil
    }
    fmt.Println("cache missed")

    req, err := http.NewRequest("GET", fullUrl, nil)
    if err != nil {
        return Pokemon{}, err
    }
    resp, err := c.httpClient.Do(req)
    if err != nil {
        return Pokemon{}, err
    }

    defer resp.Body.Close()

    if resp.StatusCode > 399 {
        return Pokemon{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
    }
    data, err := io.ReadAll(resp.Body)
    if err != nil {
        return Pokemon{}, err
    }
    pokemonData := Pokemon{}
    err = json.Unmarshal(data, &pokemonData)
    if err != nil {
        return pokemonData, err
    }
    c.cache.Add(fullUrl, data)
    return pokemonData, nil
}
