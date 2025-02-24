package http

import (
	"github.com/manhrev/labeler/configs"
	"github.com/manhrev/labeler/internal/repository/user"
	"github.com/manhrev/labeler/internal/service/token"
	"github.com/manhrev/labeler/pkg/db"
	"go.uber.org/zap"
)

type Service struct {
	logger             *zap.SugaredLogger
	tokenService       *token.Service
	authExludedPathMap map[string]string
	userRepo           *user.Repository
	userMap            map[int64]db.User
}

func NewService(configs *configs.Config, logger *zap.SugaredLogger, tokenService *token.Service, userRepo *user.Repository) *Service {
	return &Service{
		logger:             logger,
		tokenService:       tokenService,
		authExludedPathMap: configs.Server.ExcludedAuthPath,
		userRepo:           userRepo,
		userMap:            make(map[int64]db.User),
	}
}
