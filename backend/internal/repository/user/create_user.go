package user

import (
	"context"

	"github.com/manhrev/labeler/pkg/db"
)

func (r *Repository) CreateUser(ctx context.Context, username, password, displayname string, email string) (db.User, error) {
	return r.queries.CreateUser(ctx, db.CreateUserParams{
		Username:    username,
		Password:    password,
		DisplayName: displayname,
		Email:       email,
	})
}
