package domain

import "time"

type BitlinkRepository interface {
	SaveBitlink(shorten BitLink) error
	GetBitlink(id string) (BitLink, error)
}

type BitLink struct {
	Id        string
	Link      string
	LongUrl   string
	Domain    string
	CreatedAt time.Time
}

func NewBitlink(id string, long_url string, domain string) *BitLink {
	return &BitLink{
		Id:        id,
		LongUrl:   long_url,
		Link:      domain + "/" + id,
		CreatedAt: time.Now(),
		Domain:    domain,
	}
}

func (b *BitLink) GetLink() string {
	return b.Domain + "/" + b.Id
}
