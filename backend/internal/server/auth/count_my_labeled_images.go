package auth

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/manhrev/labeler/internal/const/header"
	"github.com/manhrev/labeler/internal/util"
	"github.com/manhrev/labeler/pkg/api/go/auth/rpc"
)

func (s *Server) CountMyLabeledImages(
	ctx context.Context, in *connect.Request[rpc.CountMyLabeledImagesRequest],
) (*connect.Response[rpc.CountMyLabeledImagesResponse], error) {
	count, err := s.repo.Queries.CountImagesByLabelerID(ctx, pgtype.Int8{
		Int64: util.MustParseInt64(in.Header().Get(header.UserID)),
		Valid: true,
	})
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("cannot count images: %v", err))
	}

	return connect.NewResponse(&rpc.CountMyLabeledImagesResponse{
		Count: count,
	}), nil
}
