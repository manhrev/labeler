package auth

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	"github.com/manhrev/labeler/internal/const/header"
	"github.com/manhrev/labeler/internal/model/auth"
	"github.com/manhrev/labeler/internal/util"
	"github.com/manhrev/labeler/pkg/api/go/auth/rpc"
	"github.com/manhrev/labeler/pkg/db"
)

func (s *Server) GetMyLabeledImage(
	ctx context.Context, in *connect.Request[rpc.GetMyLabeledImageRequest],
) (*connect.Response[rpc.GetMyLabeledImageResponse], error) {
	image, err := s.repo.Queries.GetImageByID(ctx, db.GetImageByIDParams{
		ID:       util.MustParseInt64(in.Msg.GetId()),
		Category: db.Category(in.Msg.GetCategory().String()),
	})
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("cannot query image or image not found: %v", err))
	}

	if image.LabelerID.Int64 != util.MustParseInt64(in.Header().Get(header.UserID)) {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("image is not labeled by you"))
	}

	return connect.NewResponse(&rpc.GetMyLabeledImageResponse{
		Image: auth.ImageDbToPb(image),
	}), nil
}
