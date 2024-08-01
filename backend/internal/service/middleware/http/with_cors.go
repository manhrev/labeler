package http

import (
	"net/http"

	connectcors "connectrpc.com/cors"
	"github.com/manhrev/labeler/internal/const/header"
	"github.com/rs/cors"
)

func (s Service) WithCORS(h http.Handler) http.Handler {
	allowedHeader := append(connectcors.AllowedHeaders(), header.Authorization)
	middleware := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: connectcors.AllowedMethods(),
		AllowedHeaders: allowedHeader,
		ExposedHeaders: connectcors.ExposedHeaders(),
	})
	return middleware.Handler(h)
}
