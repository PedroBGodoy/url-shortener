package usecase

import (
	"os"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/pedrobgodoy/url-shortener/domain"
)

type ShortenUseCase struct {
	bitlinkRepository domain.BitlinkRepository
}

func NewShortenUseCase(bitlinkRepository domain.BitlinkRepository) ShortenUseCase {
	return ShortenUseCase{bitlinkRepository: bitlinkRepository}
}

func (s *ShortenUseCase) BitLink(longUrl string) (*domain.BitLink, error) {
	id, err := gonanoid.New(7)
	if err != nil {
		return nil, err
	}

	shorten := domain.NewBitlink(id, longUrl, os.Getenv("SHORTEN_DOMAIN"))

	err = s.bitlinkRepository.SaveBitlink(*shorten)
	if err != nil {
		return nil, err
	}

	return shorten, nil
}
