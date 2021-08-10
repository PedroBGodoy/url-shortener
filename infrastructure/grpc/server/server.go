package server

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/pedrobgodoy/url-shortener/infrastructure/grpc/gen/shortener/v1"
	"github.com/pedrobgodoy/url-shortener/infrastructure/grpc/service"
	"github.com/pedrobgodoy/url-shortener/usecase"
)

type GRPCServer struct {
	ShortenLinkUseCase usecase.ShortenUseCase
	GetBitlinkUseCase  usecase.GetBitlinkUseCase
	GrpcPort           string
}

func NewGRPCServer(
	shortenLinkUseCase usecase.ShortenUseCase,
	getBitlinkUseCase usecase.GetBitlinkUseCase,
	grpcPort string,
) GRPCServer {
	return GRPCServer{
		ShortenLinkUseCase: shortenLinkUseCase,
		GetBitlinkUseCase:  getBitlinkUseCase,
		GrpcPort:           grpcPort,
	}
}

func (g *GRPCServer) Serve() {
	address := "0.0.0.0:" + g.GrpcPort
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	shortenerService := service.NewShortenerService(
		g.ShortenLinkUseCase,
		g.GetBitlinkUseCase,
	)

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	shortener.RegisterShortenerServiceServer(grpcServer, shortenerService)

	log.Printf("Serving gRPC on %s\n", address)
	log.Fatalln(grpcServer.Serve(lis))
}
