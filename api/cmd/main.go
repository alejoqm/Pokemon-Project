package main

import (
    "log"
    "net"
    "os"
    "os/signal"
    "syscall"

    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"  

    "stava/batch/internal/client"
    rpc "stava/batch/internal/rpc"
    "stava/batch/internal/service"
)



func main() {
   
    apiClient := client.NewServiceClient("https://pokeapi.co/api/v2")

    pokemonService := service.NewPokemonService(apiClient)
    
    grpcServer := grpc.NewServer()
    rpc.RegisterPokemonServiceServer(grpcServer, rpc.NewPokemonServer(pokemonService))
    
    reflection.Register(grpcServer)
    
    lis, err := net.Listen("tcp", ":50051")

    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    
    log.Println("Starting gRPC server on :50051")
    
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    
    go func() {
        if err := grpcServer.Serve(lis); err != nil {
            log.Fatalf("failed to serve gRPC: %v", err)
        }
    }()
    
    log.Println("gRPC server started on :50051")
    
    // Keep main alive or add graceful shutdown
    select {}
}