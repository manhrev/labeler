package auth

import (
	"context"
	"errors"

	"connectrpc.com/connect"
	"github.com/manhrev/labeler/pkg/api/go/auth/rpc"
)

func (s *Server) Login(
	ctx context.Context, req *connect.Request[rpc.LoginRequest],
) (*connect.Response[rpc.LoginResponse], error) {
	username := req.Msg.GetUsername()
	password := req.Msg.GetPassword()

	user, err := s.repo.User.GetUserByUsername(ctx, username)

	if err != nil {
		s.logger.Errorf("Login error: cannot query user: %v", err)
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	expectedPassword := user.Password
	if password != expectedPassword {
		s.logger.Errorf("Login error: password not match")
		return nil, connect.NewError(connect.CodeFailedPrecondition, errors.New("wrong credentials"))
	}

	tokenString, err := s.tokenService.Create(user.Username, user.ID)
	if err != nil {
		s.logger.Errorf("Login error: cannot create token for user")
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	res := connect.NewResponse(&rpc.LoginResponse{
		Token: tokenString,
	})

	return res, nil
}
