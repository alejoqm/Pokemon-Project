package client;

import (
	"context"
	"fmt"
    "io"
	"encoding/json"
	"net/http"
	"time"
)

type Pokemon struct {
    ID     int    `json:"id"`
    Name   string `json:"name"`
    Height int    `json:"height"`
    Weight int    `json:"weight"`
    Types  []struct {
        Type struct {
            Name string `json:"name"`
        } `json:"type"`
    } `json:"types"`
}

type PokemonTypeInfo struct {
    ID int `json:"id"`
    Name string `json:"name"`
    DamageRelations struct {
        NoDamageTo []struct {
            Name string `json:"name"`
        } `json:"no_damage_to"`
        HalfDamageTo []struct {
            Name string `json:"name"`
        } `json:"half_damage_to"`
        DoubleDamageTo []struct {
            Name string `json:"name"`
        } `json:"double_damage_to"`
        NoDamageFrom []struct {
            Name string `json:"name"`
        } `json:"no_damage_from"`
        HalfDamageFrom []struct {
            Name string `json:"name"`
        } `json:"half_damage_from"`
        DoubleDamageFrom []struct {
            Name string `json:"name"`
        } `json:"double_damage_from"`
    } `json:"damage_relations"`
}


type ServiceClient struct {
    HTTPClient *http.Client
	BaseURL    string
}

func NewServiceClient(baseURL string) *ServiceClient {
	return &ServiceClient{
		HTTPClient: &http.Client{Timeout: 5 * time.Second},
		BaseURL:    baseURL,
	}
}

func (c *ServiceClient) GetPokemonInfo(ctx context.Context, pokemonName string) (Pokemon, error) {
	var pokemon Pokemon

	req, err := http.NewRequestWithContext(ctx, "GET", c.BaseURL+"/pokemon/" + pokemonName, nil)
	if err != nil {
		return pokemon, err
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return pokemon, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return pokemon, fmt.Errorf("failed with status: %d", resp.StatusCode)
	}
    body, err := io.ReadAll(resp.Body)


    if (err != nil) {
        return pokemon, err
    }
    
    if err := json.Unmarshal(body, &pokemon); err != nil {
        return pokemon, err
    }
	
	return pokemon, nil
}


func (c *ServiceClient) GetTypeInfo(ctx context.Context, pokemonType string) (PokemonTypeInfo, error) {
	var pokemonTypeInfo PokemonTypeInfo

	req, err := http.NewRequestWithContext(ctx, "GET", c.BaseURL + "/type/" + pokemonType, nil)
	
	if (err != nil) {
		return pokemonTypeInfo, err
	}
	
	resp, err := c.HTTPClient.Do(req)
	if (err != nil) {
		return pokemonTypeInfo, err
	}
	
	defer resp.Body.Close()
	
	if resp.StatusCode != http.StatusOK {
		return pokemonTypeInfo, fmt.Errorf("failed with status: %d", resp.StatusCode)
	}
    body, err := io.ReadAll(resp.Body)

	
	if err := json.Unmarshal(body, &pokemonTypeInfo); err != nil {
        return pokemonTypeInfo, err
    }
	
	return pokemonTypeInfo, nil
	
}