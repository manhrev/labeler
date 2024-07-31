package auth

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	"github.com/codolabs/fushon/internal/const/header"
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

	category := model.Category_C_NONE
	switch image.Category {
	case db.CategoryCBASKETBALLCOMPETITION:
		category = model.Category_C_BASKETBALL_COMPETITION
	case db.CategoryCBASKETBALLCOMPETITOR:
		category = model.Category_C_BASKETBALL_COMPETITOR
	case db.CategoryCFOOTBALLCOMPETITION:
		category = model.Category_C_FOOTBALL_COMPETITION
	case db.CategoryCFOOTBALLCOMPETITOR:
		category = model.Category_C_FOOTBALL_COMPETITOR
	}

	backgroundType := model.BackgroundType_BT_NONE
	if image.BackgroundType.Valid {
		switch image.BackgroundType.BackgroundType {
		case db.BackgroundTypeBTFULL:
			backgroundType = model.BackgroundType_BT_FULL
		case db.BackgroundTypeBTOUTSIDE:
			backgroundType = model.BackgroundType_BT_OUTSIDE
		case db.BackgroundTypeBTNONE:
			backgroundType = model.BackgroundType_BT_NONE
		}
	}

	urlSelected := ""
	if image.UrlSelected.Valid {
		urlSelected = util.IntToString(int(image.UrlSelected.Int16))
	}

	return connect.NewResponse(&rpc.GetImageToLabelResponse{
		Image: &model.Image{
			Id:             image.ID,
			Category:       category,
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
