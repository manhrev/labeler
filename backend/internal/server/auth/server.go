package auth

import (
	"go.uber.org/zap"

	"github.com/manhrev/labeler/internal/repository"
	"github.com/manhrev/labeler/internal/service/token"
	"github.com/manhrev/labeler/pkg/api/go/auth/authconnect"
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
