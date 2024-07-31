package token

import (
	"strings"

	"github.com/codolabs/fushon/internal/model/auth"
	"github.com/golang-jwt/jwt/v5"
)

func (s *Service) Validate(tokenString string) (*auth.Claims, error) {
	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
	claims := &auth.Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (any, error) {
		return s.jwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, nil
		}
		return nil, err
	}

	if !token.Valid {
		return nil, nil
	}

	return claims, nil
}