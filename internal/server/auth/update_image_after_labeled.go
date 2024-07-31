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

func (s *Server) UpdateImageAfterLabeled(
	ctx context.Context, in *connect.Request[rpc.UpdateImageAfterLabeledRequest],
) (*connect.Response[rpc.UpdateImageAfterLabeledResponse], error) {
	var category db.Category
	var backgroundType db.BackgroundType

	switch in.Msg.GetCategory() {
	case model.Category_C_BASKETBALL_COMPETITION:
		category = db.CategoryCBASKETBALLCOMPETITION
	case model.Category_C_BASKETBALL_COMPETITOR:
		category = db.CategoryCBASKETBALLCOMPETITOR
	case model.Category_C_FOOTBALL_COMPETITION:
		category = db.CategoryCFOOTBALLCOMPETITION
	case model.Category_C_FOOTBALL_COMPETITOR:
		category = db.CategoryCFOOTBALLCOMPETITOR
	}

	switch in.Msg.GetBackgroundType() {
	case model.BackgroundType_BT_FULL:
		backgroundType = db.BackgroundTypeBTFULL
	case model.BackgroundType_BT_OUTSIDE:
		backgroundType = db.BackgroundTypeBTOUTSIDE
	case model.BackgroundType_BT_NONE:
		backgroundType = db.BackgroundTypeBTNONE
	}

	_, err := s.repo.Queries.UpdateImageAfterLabeled(ctx, db.UpdateImageAfterLabeledParams{
		ID:             util.MustParseInt64(in.Msg.GetId()),
		Category:       category,
		UrlSelected:    pgtype.Int2{Valid: true, Int16: int16(util.MustParseInt(in.Msg.GetUrlSelected()))},
		LabelerID:      pgtype.Int8{Int64: util.MustParseInt64(in.Header().Get(header.UserID)), Valid: true},
		BackgroundType: db.NullBackgroundType{BackgroundType: backgroundType, Valid: true},
	})
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("cannot rollback image  due to: %v", err))
	}

	return connect.NewResponse(&rpc.UpdateImageAfterLabeledResponse{}), nil
}
