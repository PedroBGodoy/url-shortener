package service

import (
	"context"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/pedrobgodoy/url-shortener/infrastructure/grpc/gen/shortener/v1"
	"github.com/pedrobgodoy/url-shortener/usecase"
)

type ShortenerService struct {
	ShortenLinkUseCase usecase.ShortenUseCase
	GetBitlinkUseCase  usecase.GetBitlinkUseCase
	shortener.UnimplementedShortenerServiceServer
}

func NewShortenerService(
	shortenLinkUseCase usecase.ShortenUseCase,
	getBitlinkUseCase usecase.GetBitlinkUseCase,
) *ShortenerService {
	return &ShortenerService{
		ShortenLinkUseCase: shortenLinkUseCase,
		GetBitlinkUseCase:  getBitlinkUseCase,
	}
}

func (s *ShortenerService) Shorten(
	ctx context.Context,
	in *shortener.ShortenRequest,
) (*shortener.Bitlink, error) {
	bitlink, err := s.ShortenLinkUseCase.BitLink(in.GetLongUrl())
	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	res := &shortener.Bitlink{
		BitlinkId: bitlink.Id,
		Link:      bitlink.Link,
		LongUrl:   bitlink.LongUrl,
		CreatedAt: bitlink.CreatedAt.Format(time.RFC3339),
	}

	return res, nil
}

func (s *ShortenerService) GetBitlink(
	ctx context.Context,
	in *shortener.GetBitlinkRequest,
) (*shortener.Bitlink, error) {
	bitlink, err := s.GetBitlinkUseCase.GetBitlink(in.GetBitlinkId())
	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	res := &shortener.Bitlink{
		BitlinkId: bitlink.Id,
		Link:      bitlink.Link,
		LongUrl:   bitlink.LongUrl,
		CreatedAt: bitlink.CreatedAt.Format(time.RFC3339),
	}

	return res, nil
}
