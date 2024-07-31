package grpc

import (
	"github.com/codolabs/fushon/configs"
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
