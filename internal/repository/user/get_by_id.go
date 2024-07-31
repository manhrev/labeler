package user

import (
	"context"

	"github.com/manhrev/labeler/pkg/db"
)

func (r *Repository) GetUserByID(ctx context.Context, id int64) (db.User, error) {
	return r.queries.GetUserById(ctx, id)
}
