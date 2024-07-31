package auth

import (
	"context"
	"strings"

	"connectrpc.com/connect"
	"github.com/codolabs/fushon/pkg/api/go/auth/model"
	"github.com/codolabs/fushon/pkg/api/go/auth/rpc"
)

func (s *Server) SignUp(
	ctx context.Context, req *connect.Request[rpc.SignUpRequest],
) (*connect.Response[rpc.SignUpResponse], error) {
	username := strings.TrimSpace(req.Msg.GetUsername())
	pwd := strings.TrimSpace(req.Msg.GetPassword())

	newUser, err := s.repo.User.CreateUser(ctx, username, pwd, req.Msg.GetDisplayName(), req.Msg.GetEmail())
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(&rpc.SignUpResponse{
		Info: &model.UserInfo{
			Id:          newUser.ID,
			DisplayName: newUser.DisplayName,
			Username:    newUser.Username,
			Email:       newUser.Email,
		},
	}), nil
}
