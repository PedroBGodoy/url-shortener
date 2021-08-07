package main

import (
	"os"

	"github.com/joho/godotenv"

	"github.com/pedrobgodoy/url-shortener/infrastructure/gateway"
	"github.com/pedrobgodoy/url-shortener/infrastructure/grpc/server"
	"github.com/pedrobgodoy/url-shortener/usecase"
)

func main() {
	godotenv.Load()

	shortenLinkUC := usecase.NewShortenLink()

	grpcServer := server.NewGRPCServer(shortenLinkUC, os.Getenv("GRPC_PORT"))
	go func() {
		grpcServer.Serve()
	}()

	gateway.Run(":"+os.Getenv("GRPC_PORT"), os.Getenv("HTTP_PORT"))
}
