package auth

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/manhrev/labeler/internal/const/header"
	"github.com/manhrev/labeler/internal/util"
	"github.com/manhrev/labeler/pkg/api/go/auth/rpc"
	"github.com/manhrev/labeler/pkg/db"
)

func (s *Server) RollbackLabeledImage(
	ctx context.Context, in *connect.Request[rpc.RollbackLabeledImageRequest],
) (*connect.Response[rpc.RollbackLabeledImageResponse], error) {
	err := s.repo.Queries.RollbackImageLabeled(ctx, db.RollbackImageLabeledParams{
		ID:        util.MustParseInt64(in.Msg.GetId()),
		Category:  db.Category(in.Msg.GetCategory().String()),
		LabelerID: pgtype.Int8{Int64: util.MustParseInt64(in.Header().Get(header.UserID)), Valid: true},
	})

	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("cannot rollback image  due to: %v", err))
	}

	return connect.NewResponse(&rpc.RollbackLabeledImageResponse{}), nil
}
