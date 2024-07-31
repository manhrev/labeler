package auth

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	"github.com/codolabs/fushon/internal/const/header"
	"github.com/codolabs/fushon/internal/model/auth"
	"github.com/codolabs/fushon/internal/util"
	"github.com/codolabs/fushon/pkg/api/go/auth/model"
	"github.com/codolabs/fushon/pkg/api/go/auth/rpc"
	"github.com/codolabs/fushon/pkg/db"
	"github.com/jackc/pgx/v5/pgtype"
)

func (s *Server) GetImageToLabel(
	ctx context.Context, in *connect.Request[rpc.GetImageToLabelRequest],
) (*connect.Response[rpc.GetImageToLabelResponse], error) {
	image, err := s.repo.Queries.GetImageToLabel(ctx)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("cannot query image or image not found: %v", err))
	}

	if _, err := s.repo.Queries.UpdateImageLabelerID(ctx, db.UpdateImageLabelerIDParams{
		ID:       image.ID,
		Category: image.Category,
		LabelerID: pgtype.Int8{
			Int64: util.MustParseInt64(in.Header().Get(header.UserID)),
			Valid: true,
		},
	}); err != nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("cannot update image labeler id: %v", err))
	}

	backgroundType := model.BackgroundType_BT_NONE
	if image.BackgroundType.Valid {
		backgroundType = auth.BackgroundTypeDbToPb(image.BackgroundType.BackgroundType)
	}

	urlSelected := ""
	if image.UrlSelected.Valid {
		urlSelected = util.IntToString(int(image.UrlSelected.Int16))
	}

	return connect.NewResponse(&rpc.GetImageToLabelResponse{
		Image: &model.Image{
			Id:             image.ID,
			Category:       auth.CategoryDbToPb(image.Category),
			BackgroundType: backgroundType,
			LabelerId:      util.IntToString(int(image.LabelerID.Int64)),
			Name:           image.Name,
			DisplayName:    image.DisplayName.String,
			Url1:           image.Url1,
			Url2:           image.Url2,
			Url3:           image.Url3,
			UrlSelected:    urlSelected,
		},
	}), nil
}
