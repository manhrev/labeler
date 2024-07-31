package http

import (
	"github.com/codolabs/fushon/configs"
	"github.com/codolabs/fushon/internal/repository/user"
	"github.com/codolabs/fushon/internal/service/token"
	"go.uber.org/zap"
)

type Service struct {
	logger             *zap.SugaredLogger
	tokenService       *token.Service
	authExludedPathMap map[string]string
	userRepo           *user.Repository
}

func NewService(configs *configs.Config, logger *zap.SugaredLogger, tokenService *token.Service, userRepo *user.Repository) *Service {
	return &Service{
		logger:             logger,
		tokenService:       tokenService,
		authExludedPathMap: configs.Server.ExcludedAuthPath,
		userRepo:           userRepo,
	}
}
