package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"github.com/pedrobgodoy/url-shortener/infrastructure/gateway"
	"github.com/pedrobgodoy/url-shortener/infrastructure/grpc/server"
	"github.com/pedrobgodoy/url-shortener/infrastructure/repository"
	"github.com/pedrobgodoy/url-shortener/usecase"
)

func main() {
	godotenv.Load()

	db := setupDb()
	bitlinkRepository := repository.NewBitlinkRepositoryDb(db)
	shortenLinkUC := usecase.NewShortenUseCase(bitlinkRepository)
	getBitlinkUC := usecase.NewGetBitlinkUseCase(bitlinkRepository)

	grpcServer := server.NewGRPCServer(
		shortenLinkUC,
		getBitlinkUC,
		os.Getenv("GRPC_PORT"),
	)
	go func() {
		grpcServer.Serve()
	}()

	gateway.Run(":"+os.Getenv("GRPC_PORT"), os.Getenv("HTTP_PORT"))
}

func setupDb() *sql.DB {
	db, err := sql.Open("postgres", os.Getenv("DB_STRING"))
	if err != nil {
		log.Fatal("error connection to database")
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("error connection to database: %s", err.Error())
	}
	return db
}
