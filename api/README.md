# Pokemon-Project
A gRPC-based Pokémon API that fetches and enriches Pokémon data from PokeAPI, returning base stats, sprites, and type damage relations.

## Description
This project implements a backend service that acts as an intermediary between clients and the public PokeAPI. It combines data from multiple endpoints into a single unified response, including base Pokémon stats, sprite images, and comprehensive type-based damage relations (immunities, resistances, and weaknesses). Built in Go using gRPC for efficient communication and structured API contracts.

## Getting Started
Dependencies
Go 1.18+
protoc (Protocol Buffers compiler)
Go protobuf plugins:
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

## Installing
Clone the repository and download dependencies:

git clone https://github.com/alejoqm/Pokemon-Project.git
cd Pokemon-Project/api
go mod download

### How to run the project
Start the gRPC server from the api directory:

go run ./cmd

The server listens on:

127.0.0.1:50051

### Regenerate Proto Files
After modifying pokemon.proto, regenerate the Go code:

cd api/proto

protoc --go_out=../internal/rpc --go_opt=paths=source_relative --go-grpc_out=../internal/rpc --go-grpc_opt=paths=source_relative pokemon.proto

## Testing
Sample request using grpcurl:

```
grpcurl -plaintext -d '{"name": "pikachu"}' localhost:50051 pokemon.GetPokemon
```
Sample response:

```
{
  "id": 25,
  "name": "pikachu",
  "types": [
    "electric"
  ],
  "sprites": {
    "frontDefault": "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/25.png",
    "backDefault": "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/back/25.png"
  },
  "damageRelations": {
    "noDamageTo": [
      "ground"
    ],
    "halfDamageTo": [
      "grass",
      "electric",
      "dragon"
    ],
    "doubleDamageTo": [
      "flying",
      "water"
    ],
    "noDamageFrom": [],
    "halfDamageFrom": [
      "flying",
      "steel",
      "electric"
    ],
    "doubleDamageFrom": [
      "ground"
    ]
  }
}
```

### Help
Console locked (process still running)
Find and kill the running process.

Git Bash:

ps aux | grep "go run.*pokemon" | grep -v grep
kill <PID>

PowerShell:

Get-CimInstance Win32_Process | Where-Object {$_.CommandLine -like "*Pokemon*"}
Stop-Process -Id <PID> -Force

Port already in use
Find the process using port 50051:

netstat -ano | findstr :50051

Kill it by PID:

Stop-Process -Id 24204 -Force

### Project Structure
Pokemon-Project/
└── api/
    ├── cmd/              # Application entry point
    ├── internal/
    │   ├── client/       # PokeAPI HTTP client
    │   ├── rpc/          # Generated gRPC code
    │   └── service/      # Business logic
    ├── proto/            # .proto definitions
    ├── go.mod
    ├── go.sum
    └── README.md

### Authors
Alejandro Quintero @alejoqm

Version History
0.1
Initial Release
