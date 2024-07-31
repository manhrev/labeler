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

func (s *Server) RollbackLabeledImage(
	ctx context.Context, in *connect.Request[rpc.RollbackLabeledImageRequest],
) (*connect.Response[rpc.RollbackLabeledImageResponse], error) {
	var category db.Category

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

	_, err := s.repo.Queries.UpdateImageUrlSelectedByLabelerID(ctx, db.UpdateImageUrlSelectedByLabelerIDParams{
		ID:          util.MustParseInt64(in.Msg.GetId()),
		Category:    category,
		UrlSelected: pgtype.Int2{Valid: false},
		LabelerID:   pgtype.Int8{Int64: util.MustParseInt64(in.Header().Get(header.UserID)), Valid: true},
	})
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("cannot rollback image  due to: %v", err))
	}

	return connect.NewResponse(&rpc.RollbackLabeledImageResponse{}), nil
}
