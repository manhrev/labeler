package middleware

import (
	"github.com/codolabs/fushon/configs"
	"github.com/codolabs/fushon/internal/repository"
	"github.com/codolabs/fushon/internal/service/middleware/grpc"
	"github.com/codolabs/fushon/internal/service/middleware/http"
	"github.com/codolabs/fushon/internal/service/token"
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
