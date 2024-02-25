package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageUrl *string) (LocationAreasResponse, error)  {
    endpoint := "/location-area"  
    fullUrl := baseURL + endpoint
    if pageUrl != nil {
        fullUrl = *pageUrl
    }

    cached, ok := c.cache.Get(fullUrl)
    if ok {
        fmt.Println("cache hit")
        locationAreas := LocationAreasResponse{}
        err := json.Unmarshal(cached, &locationAreas)
        if err != nil {
            return locationAreas, err
        }
        return locationAreas, nil
    }
    fmt.Println("cache missed")

    req, err := http.NewRequest("GET", fullUrl, nil)
    if err != nil {
        return LocationAreasResponse{}, err
    }
    resp, err := c.httpClient.Do(req)
    if err != nil {
        return LocationAreasResponse{}, err
    }

    defer resp.Body.Close()

    if resp.StatusCode > 399 {
        return LocationAreasResponse{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
    }
    data, err := io.ReadAll(resp.Body)
    if err != nil {
        return LocationAreasResponse{}, err
    }
    locationAreas := LocationAreasResponse{}
    err = json.Unmarshal(data, &locationAreas)
    if err != nil {
        return locationAreas, err
    }
    c.cache.Add(fullUrl, data)
    return locationAreas, nil
}

func (c *Client) GetLocationArea(locationAreaName string) (LocationArea, error){
    endpoint := "/location-area/" + locationAreaName
    fullUrl := baseURL + endpoint

    cached, ok := c.cache.Get(fullUrl)
    if ok {
        fmt.Println("cache hit")
        locationArea := LocationArea{}
        err := json.Unmarshal(cached, &locationArea)
        if err != nil {
            return locationArea, err
        }
        return locationArea, nil
    }
    fmt.Println("cache missed")

    req, err := http.NewRequest("GET", fullUrl, nil)
    if err != nil {
        return LocationArea{}, err
    }
    resp, err := c.httpClient.Do(req)
    if err != nil {
        return LocationArea{}, err
    }

    defer resp.Body.Close()

    if resp.StatusCode > 399 {
        return LocationArea{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
    }
    data, err := io.ReadAll(resp.Body)
    if err != nil {
        return LocationArea{}, err
    }
    locationArea := LocationArea{}
    err = json.Unmarshal(data, &locationArea)
    if err != nil {
        return locationArea, err
    }
    c.cache.Add(fullUrl, data)
    return locationArea, nil
}
