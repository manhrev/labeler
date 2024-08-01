package user

import (
	"context"

	"github.com/manhrev/labeler/pkg/db"
)

func (r *Repository) GetUserByUsername(ctx context.Context, username string) (db.User, error) {
	return r.queries.GetUserByUsername(ctx, username)
}
