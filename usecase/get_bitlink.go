package usecase

import (
	"github.com/pedrobgodoy/url-shortener/domain"
)

type GetBitlinkUseCase struct {
	bitlinkRepository domain.BitlinkRepository
}

func NewGetBitlinkUseCase(bitlinkRepository domain.BitlinkRepository) GetBitlinkUseCase {
	return GetBitlinkUseCase{bitlinkRepository: bitlinkRepository}
}

func (s *GetBitlinkUseCase) GetBitlink(id string) (*domain.BitLink, error) {
	bitlink, err := s.bitlinkRepository.GetBitlink(id)
	if err != nil {
		return nil, err
	}

	return &bitlink, nil
}
