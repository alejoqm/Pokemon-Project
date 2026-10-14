package rpc

import (
    "context"

    "stava/batch/internal/service"
)

type PokemonServer struct {
    UnimplementedPokemonServiceServer
    pokemonSvc *service.PokemonService
}

func NewPokemonServer(pokemonSvc *service.PokemonService) *PokemonServer {
    return &PokemonServer{
        pokemonSvc: pokemonSvc,
    }
}

func (s *PokemonServer) GetPokemon(ctx context.Context, req *GetPokemonRequest) (*GetPokemonResponse, error) {
    pokemon, err := s.pokemonSvc.GetPokemonInfo(ctx, req.Name)
    if err != nil {
        return nil, err
    }
    
    return &GetPokemonResponse{
        Name:  pokemon.Name,
        Id:    int32(pokemon.ID),
        Types: pokemon.Types,
        // Map other fields
    }, nil
}