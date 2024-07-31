package repository

import (
	"github.com/codolabs/fushon/internal/repository/user"
	"github.com/codolabs/fushon/pkg/db"
	"github.com/jackc/pgx/v5"
)

type Repository struct {
	Queries *db.Queries
	User    *user.Repository
	DB      *pgx.Conn
}

func New(queries *db.Queries, db *pgx.Conn) *Repository {
	return &Repository{
		Queries: queries,
		User:    user.NewRepository(queries),
		DB:      db,
	}
}
