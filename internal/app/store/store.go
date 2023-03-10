package store

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Store struct {
	config *Config
	db     *sql.DB
}

func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

func (s *Store) Open() error {

	var err error

	s.db, err = sql.Open("postgres", s.config.DatabaseURL)

	if err != nil {

		return err

	}

	if err := s.db.Ping(); err != nil {

		return err

	}

	return nil

}

func (s *Store) Close() {
	s.db.Close()
}
