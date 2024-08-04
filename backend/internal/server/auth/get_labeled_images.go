package auth

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	"github.com/manhrev/labeler/internal/model/auth"
	"github.com/manhrev/labeler/internal/util"
	"github.com/manhrev/labeler/pkg/api/go/auth/rpc"
	"github.com/manhrev/labeler/pkg/db"
)

func (s *Server) GetLabeledImages(
	ctx context.Context, in *connect.Request[rpc.GetLabeledImagesRequest],
) (*connect.Response[rpc.GetLabeledImagesResponse], error) {
	var (
		page = in.Msg.GetPage()
		size = in.Msg.GetSize()
	)

	images, err := s.repo.Queries.FindImagesByFilters(ctx, db.FindImagesByFiltersParams{
		Limit:     int32(size),
		Offset:    int32(page * size),
		Category:  db.Category(in.Msg.GetCategory()),
		LabelerID: util.MustParseInt64(in.Msg.GetLabelerId()),
	})
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("cannot query images: %v", err))
	}

	return connect.NewResponse(&rpc.GetLabeledImagesResponse{
		Images: auth.ImagesDbToPbArray(images),
	}), nil
}
