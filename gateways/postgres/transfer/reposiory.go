package transfer

import "github.com/jackc/pgx/v5"

type Repository struct {
	db pgx.Conn
}

func NewRepository(db pgx.Conn) (Repository, error) {
	return Repository{db}, nil
}
