package middleware

import (
	"github.com/manhrev/labeler/configs"
	"github.com/manhrev/labeler/internal/repository"
	"github.com/manhrev/labeler/internal/service/middleware/grpc"
	"github.com/manhrev/labeler/internal/service/middleware/http"
	"github.com/manhrev/labeler/internal/service/token"
	"go.uber.org/zap"
)

type Service struct {
	Http *http.Service
	Grpc *grpc.Service
}

func NewService(configs *configs.Config, logger *zap.SugaredLogger, tokenService *token.Service, repo *repository.Repository) *Service {
	return &Service{
		Http: http.NewService(configs, logger, tokenService, repo.User),
		Grpc: grpc.NewService(configs, logger),
	}
}
