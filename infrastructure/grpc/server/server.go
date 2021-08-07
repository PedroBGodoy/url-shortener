package server

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/pedrobgodoy/url-shortener/infrastructure/grpc/protofile/shortener/v1"
	"github.com/pedrobgodoy/url-shortener/infrastructure/grpc/service"
	"github.com/pedrobgodoy/url-shortener/usecase"
)

type GRPCServer struct {
	ShortenLinkUseCase usecase.ShortenLink
}

func NewGRPCServer(shortenLinkUseCase usecase.ShortenLink) GRPCServer {
	return GRPCServer{ShortenLinkUseCase: shortenLinkUseCase}
}

func (g *GRPCServer) Serve(address string) {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("could not listen tcp port")
	}

	shortenerService := service.NewShortenerService(g.ShortenLinkUseCase)

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	shortener.RegisterShortenerServiceServer(grpcServer, shortenerService)

	grpcServer.Serve(lis)
}
