package service

import (
    "context"
    "log"
    "fmt"
    "errors"
    "stava/batch/internal/client"
)

type PokemonBaseInfo struct {
    Name string;
}

type PokemonDetails struct {
    ID int
    Name string 
    Height int
    Weight int
    Types []string
    DamageRelations DamageRelations
    Sprites Sprites
}

type DamageRelations struct {
    NoDamageTo []string
    HalfDamageTo []string
    DoubleDamageTo []string
    NoDamageFrom []string
    HalfDamageFrom []string
    DoubleDamageFrom []string
}

type Sprites struct {
    FrontDefault string
    BackDefault string
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
    
    if pokemonName == "" {
		return PokemonDetails{}, errors.New("pokemonName must not be empty")
	}
    
    if (err != nil) {
        return PokemonDetails{}, err
    }
    
    pokemonDetails := PokemonDetails{
        ID:     response.ID,
        Name:   response.Name,
        Height: response.Height,
        Weight: response.Weight,
        Types:  make([]string, 0, len(response.Types)),
        Sprites: Sprites{
            FrontDefault: response.Sprites.FrontDefault,
            BackDefault: response.Sprites.BackDefault,
        },
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
        
        pokemonDetails.DamageRelations.NoDamageTo = append(pokemonDetails.DamageRelations.NoDamageTo, 
                                           extractTypes(dataType.DamageRelations.NoDamageTo)...)
        pokemonDetails.DamageRelations.NoDamageTo = dedupe(pokemonDetails.DamageRelations.NoDamageTo)
        
        
        pokemonDetails.DamageRelations.HalfDamageTo = append(pokemonDetails.DamageRelations.HalfDamageTo, 
                                           extractTypes(dataType.DamageRelations.HalfDamageTo)...)
        pokemonDetails.DamageRelations.HalfDamageTo = dedupe(pokemonDetails.DamageRelations.HalfDamageTo)
        
        
        pokemonDetails.DamageRelations.DoubleDamageTo = append(pokemonDetails.DamageRelations.DoubleDamageTo, 
                                           extractTypes(dataType.DamageRelations.DoubleDamageTo)...)
        pokemonDetails.DamageRelations.DoubleDamageTo = dedupe(pokemonDetails.DamageRelations.DoubleDamageTo)
        
        
        pokemonDetails.DamageRelations.NoDamageFrom = append(pokemonDetails.DamageRelations.NoDamageFrom, 
                                           extractTypes(dataType.DamageRelations.NoDamageFrom)...)
        pokemonDetails.DamageRelations.NoDamageFrom = dedupe(pokemonDetails.DamageRelations.NoDamageFrom)
        
        
        pokemonDetails.DamageRelations.HalfDamageFrom = append(pokemonDetails.DamageRelations.HalfDamageFrom, 
                                           extractTypes(dataType.DamageRelations.HalfDamageFrom)...)
        pokemonDetails.DamageRelations.HalfDamageFrom = dedupe(pokemonDetails.DamageRelations.HalfDamageFrom)
        
        pokemonDetails.DamageRelations.DoubleDamageFrom = append(pokemonDetails.DamageRelations.DoubleDamageFrom, 
                                           extractTypes(dataType.DamageRelations.DoubleDamageFrom)...)
        pokemonDetails.DamageRelations.DoubleDamageFrom = dedupe(pokemonDetails.DamageRelations.DoubleDamageFrom)
        
    }
    
    return pokemonDetails, nil;
}

func (s *PokemonService) GetPokemonsByType(ctx context.Context, pokemonType string, total int) ([]PokemonBaseInfo, error) {

    if pokemonType == "" {
		return nil, errors.New("pokemonType must not be empty")
	}
    
    response, err := s.client.GetTypeInfo(ctx, pokemonType)
    
    if (err != nil) {
        return nil, err
    }
    
    results := make([]PokemonBaseInfo, 0, min(total, len(response.Pokemon)))
    
    limit := min(int(total), len(response.Pokemon))
	for i := 0; i < limit; i++ {
		p := response.Pokemon[i]
		results = append(results, PokemonBaseInfo{
			Name: p.Pokemon.Name,
		})
	}
    fmt.Print(results)
    return results,  nil
}

func extractTypes(items []struct{ Name string `json:"name"` }) []string {
	out := make([]string, 0, len(items))
	for _, item := range items {
		out = append(out, item.Name)
	}
	return out
}

func dedupe(in []string) []string {
	seen := make(map[string]struct{}, len(in))
	out := in[:0]
	for _, v := range in {
		if _, ok := seen[v]; ok {
			continue
		}
		seen[v] = struct{}{}
		out = append(out, v)
	}
	return out
}