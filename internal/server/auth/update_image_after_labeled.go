package auth

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	"github.com/codolabs/fushon/internal/const/header"
	"github.com/codolabs/fushon/internal/util"
	"github.com/codolabs/fushon/pkg/api/go/auth/rpc"
	"github.com/codolabs/fushon/pkg/db"
	"github.com/jackc/pgx/v5/pgtype"
)

func (s *Server) UpdateImageAfterLabeled(
	ctx context.Context, in *connect.Request[rpc.UpdateImageAfterLabeledRequest],
) (*connect.Response[rpc.UpdateImageAfterLabeledResponse], error) {

	_, err := s.repo.Queries.UpdateImageAfterLabeled(ctx, db.UpdateImageAfterLabeledParams{
		ID:       util.MustParseInt64(in.Msg.GetId()),
		Category: db.Category(in.Msg.GetCategory().String()),
		UrlSelected: pgtype.Int2{
			Int16: int16(util.MustParseInt(in.Msg.GetUrlSelected())),
			Valid: true},
		LabelerID: pgtype.Int8{
			Int64: util.MustParseInt64(in.Header().Get(header.UserID)),
			Valid: true},
		BackgroundType: db.NullBackgroundType{
			BackgroundType: db.BackgroundType(in.Msg.GetBackgroundType().String()),
			Valid:          true},
	})
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("cannot update image  due to: %v", err))
	}

	return connect.NewResponse(&rpc.UpdateImageAfterLabeledResponse{}), nil
}
