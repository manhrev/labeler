package configs

import (
	"os"

	"github.com/codolabs/fushon/pkg/api/go/auth/authconnect"
)

func GetConfigs() *Config {
	sslMode := os.Getenv("DB_SSL_MODE")
	dbSSLMode := ""
	if sslMode != "" {
		dbSSLMode = "sslmode=" + sslMode
	}
	return &Config{
		Environment: os.Getenv("ENVIRONMENT"),
		Server: &Server{
			ListenPort: 8080,
			ExcludedAuthPath: map[string]string{
				"Login":  authconnect.AuthServiceLoginProcedure,
				"SignUp": authconnect.AuthServiceSignUpProcedure,
			},
			JWTKey:                            os.Getenv("JWT_KEY"),
			JWTTokenExpirationDurationMinutes: 86400,
		},
		DB: &Database{
			DBUsername: os.Getenv("DB_USERNAME"),
			DBPassword: os.Getenv("DB_PASSWORD"),
			DBDomain:   os.Getenv("DB_HOST"),
			DBPort:     os.Getenv("DB_PORT"),
			DBName:     os.Getenv("DB_NAME"),
			DBSSLMode:  dbSSLMode,
		},
	}

}
