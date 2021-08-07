package usecase

type ShortenLink struct {
}

func NewShortenLink() ShortenLink {
	return ShortenLink{}
}

func (s *ShortenLink) Shorten(longUrl string) (string, error) {
	return "Hello UseCase 2", nil
}
