package token

import (
	"strconv"
	"time"

	"github.com/codolabs/fushon/internal/model/auth"
	"github.com/golang-jwt/jwt/v5"
)

func (s *Service) Create(username string, userID int64) (string, error) {
	minutes := s.expirationDuration
	expirationTime := time.Now().Add(time.Duration(minutes) * time.Minute)
	claims := auth.Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   strconv.Itoa(int(userID)),
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(s.jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
