package auth

import (
	"go.uber.org/zap"

	"github.com/codolabs/fushon/internal/repository"
	"github.com/codolabs/fushon/internal/service/token"
	"github.com/codolabs/fushon/pkg/api/go/auth/authconnect"
)

type Server struct {
	authconnect.UnimplementedAuthServiceHandler
	logger *zap.SugaredLogger
	repo   *repository.Repository

	tokenService *token.Service
}

func NewServer(logger *zap.SugaredLogger, tokenService *token.Service, repository *repository.Repository) authconnect.AuthServiceHandler {
	return &Server{
		logger:       logger,
		tokenService: tokenService,
		repo:         repository,
	}
}
