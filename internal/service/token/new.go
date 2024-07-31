package token

import "github.com/codolabs/fushon/configs"

type Service struct {
	jwtKey             []byte
	expirationDuration int
}

func NewService(configs *configs.Config) *Service {
	serverConfigs := configs.Server

	return &Service{
		jwtKey:             []byte(serverConfigs.JWTKey),
		expirationDuration: serverConfigs.JWTTokenExpirationDurationMinutes,
	}
}
