package repository

import (
	"database/sql"
	"errors"
	"log"

	"github.com/pedrobgodoy/url-shortener/domain"
)

type BitlinkRepositoryDb struct {
	db *sql.DB
}

func NewBitlinkRepositoryDb(db *sql.DB) *BitlinkRepositoryDb {
	return &BitlinkRepositoryDb{db: db}
}

func (s *BitlinkRepositoryDb) SaveBitlink(bitlink domain.BitLink) error {
	stmt, err := s.db.Prepare("INSERT INTO bitlink (id, long_url, created_at, domain) VALUES ($1, $2, $3, $4)")
	if err != nil {
		log.Println(err.Error())
		return err
	}

	_, err = stmt.Exec(
		bitlink.Id,
		bitlink.LongUrl,
		bitlink.CreatedAt,
		bitlink.Domain,
	)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	err = stmt.Close()
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func (s *BitlinkRepositoryDb) GetBitlink(id string) (domain.BitLink, error) {
	var b domain.BitLink
	stmt, err := s.db.Prepare("SELECT id, long_url, domain, created_at FROM bitlink WHERE id = $1")
	if err != nil {
		log.Println(err.Error())
		return b, err
	}
	if err = stmt.QueryRow(id).Scan(&b.Id, &b.LongUrl, &b.Domain, &b.CreatedAt); err != nil {
		return b, errors.New("bitlink does not exists")
	}
	b.Link = b.Domain + "/" + b.Id
	return b, nil
}
