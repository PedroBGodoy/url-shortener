package service

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/pedrobgodoy/url-shortener/infrastructure/grpc/protofile/shortener/v1"
	"github.com/pedrobgodoy/url-shortener/usecase"
)

type ShortenerService struct {
	ShortenLinkUseCase usecase.ShortenLink
	shortener.UnimplementedShortenerServiceServer
}

func NewShortenerService(shortenLinkUseCase usecase.ShortenLink) *ShortenerService {
	return &ShortenerService{ShortenLinkUseCase: shortenLinkUseCase}
}

func (s *ShortenerService) Shorten(
	ctx context.Context,
	in *shortener.CreateShortenRequest,
) (*shortener.CreateShortenResponse, error) {
	link, err := s.ShortenLinkUseCase.Shorten(in.GetLongUrl())
	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	res := &shortener.CreateShortenResponse{
		LongUrl: in.GetLongUrl(),
		Link:    link,
	}

	return res, nil
}
