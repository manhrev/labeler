package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/manhrev/labeler/internal/repository/user"
	"github.com/manhrev/labeler/pkg/db"
)

type Repository struct {
	Queries *db.Queries
	User    *user.Repository
	DB      *pgxpool.Pool
}

func New(queries *db.Queries, db *pgxpool.Pool) *Repository {
	return &Repository{
		Queries: queries,
		User:    user.NewRepository(queries),
		DB:      db,
	}
}
