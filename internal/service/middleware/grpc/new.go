package grpc

import (
	"github.com/manhrev/labeler/configs"
	"go.uber.org/zap"
)

type Service struct {
	logger *zap.SugaredLogger
}

func NewService(configs *configs.Config, logger *zap.SugaredLogger) *Service {
	return &Service{
		logger: logger,
	}
}
