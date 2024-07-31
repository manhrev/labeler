package auth

import (
	"context"
	"errors"
	"strconv"

	"connectrpc.com/connect"
	"github.com/manhrev/labeler/internal/const/header"
	"github.com/manhrev/labeler/pkg/api/go/auth/model"
	"github.com/manhrev/labeler/pkg/api/go/auth/rpc"
)

func (s *Server) Me(
	ctx context.Context, req *connect.Request[rpc.MeRequest],
) (*connect.Response[rpc.MeResponse], error) {
	userIdStr := req.Header().Get(header.UserID)
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		s.logger.Errorf("Login error: cannot parse user id from header: %v", err)
		return nil, connect.NewError(connect.CodeInternal, errors.New("cannot get user id"))
	}
	user, err := s.repo.User.GetUserByID(ctx, int64(userId))
	if err != nil {
		s.logger.Errorf("Login error: cannot query user: %v", err)
		return nil, connect.NewError(connect.CodeInternal, errors.New("cannot query user or user not found"))
	}

	return connect.NewResponse(&rpc.MeResponse{
		Info: &model.UserInfo{
			Id:          user.ID,
			DisplayName: user.DisplayName,
			Username:    user.Username,
			Email:       user.Email,
		},
	}), nil
}
