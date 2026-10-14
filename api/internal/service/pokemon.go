package service

import (
    "context"
    "fmt"
    "log"
    "stava/batch/internal/client"
)

type PokemonDetails struct {
    ID int
    Name string 
    Height int
    Weight int
    Types []string
    NoDamageTo []string
}

type PokemonService struct {
    client *client.ServiceClient
}

func NewPokemonService(client *client.ServiceClient) *PokemonService {
    return &PokemonService{
        client: client,
    }
}

func (s *PokemonService) GetPokemonInfo(ctx context.Context, pokemonName string) (PokemonDetails, error) {
    response, err := s.client.GetPokemonInfo(ctx, pokemonName)
    
    if (err != nil) {
        return PokemonDetails{}, err
    }
    
    pokemonDetails := PokemonDetails{
        ID:     response.ID,
        Name:   response.Name,
        Height: response.Height,
        Weight: response.Weight,
        Types:  make([]string, 0, len(response.Types)),
    }
    
    
    for _, value := range response.Types {
        pokemonDetails.Types = append(
            pokemonDetails.Types,
            value.Type.Name,
        )
        dataType, err := s.client.GetTypeInfo(ctx, value.Type.Name)
        
        if (err != nil) {
            log.Fatal(err)
        }
        
        fmt.Println(dataType)
        
    }
    
    return pokemonDetails, nil;
}