package user

import (
	"context"

	"github.com/codolabs/fushon/pkg/db"
)

func (r *Repository) GetUserByUsername(ctx context.Context, username string) (db.User, error) {
	return r.queries.GetUserByUsername(ctx, username)
}
