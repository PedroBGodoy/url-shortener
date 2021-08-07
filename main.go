package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/pedrobgodoy/url-shortener/infrastructure/grpc/server"
	"github.com/pedrobgodoy/url-shortener/usecase"
)

func main() {
	loadEnvironmentVariables()

	shortenLinkUC := usecase.NewShortenLink()

	grpcServer := server.NewGRPCServer(shortenLinkUC)
	address := fmt.Sprintf("0.0.0.0:%s", os.Getenv("GRPC_PORT"))
	log.Printf("Rodando servidor gRPC %s\n", address)
	grpcServer.Serve(address)
}

func loadEnvironmentVariables() {
	godotenv.Load()
}
