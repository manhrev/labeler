package auth

import (
	"context"
	"errors"
	"fmt"

	"connectrpc.com/connect"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/manhrev/labeler/internal/const/header"
	"github.com/manhrev/labeler/internal/model/auth"
	"github.com/manhrev/labeler/internal/util"
	"github.com/manhrev/labeler/pkg/api/go/auth/rpc"
	"github.com/manhrev/labeler/pkg/db"
)

func (s *Server) GetImageToLabel(
	ctx context.Context, in *connect.Request[rpc.GetImageToLabelRequest],
) (*connect.Response[rpc.GetImageToLabelResponse], error) {
	image, err := s.repo.Queries.GetImageToLabel(ctx, pgtype.Int8{
		Int64: util.MustParseInt64(in.Header().Get(header.UserID)),
		Valid: true,
	})
	if err != nil {
		// s.logger.Errorf("Cannot query image or image not found: %v", err)
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, connect.NewError(connect.CodeNotFound, fmt.Errorf("no image to label: %v", err))
		}
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("cannot query image: %v", err))
	}

	if err = s.repo.Queries.UpdateImageLabelerID(ctx, db.UpdateImageLabelerIDParams{
		ID:       image.ID,
		Category: image.Category,
		LabelerID: pgtype.Int8{
			Int64: util.MustParseInt64(in.Header().Get(header.UserID)),
			Valid: true,
		},
	}); err != nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("cannot update image labeler id: %v", err))
	}

	return connect.NewResponse(&rpc.GetImageToLabelResponse{
		Image: auth.ImageDbToPb(image),
	}), nil
}
