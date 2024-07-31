package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"connectrpc.com/connect"
	"connectrpc.com/grpcreflect"
	"github.com/codolabs/fushon/configs"
	"github.com/codolabs/fushon/internal/repository"
	"github.com/codolabs/fushon/internal/server/auth"
	"github.com/codolabs/fushon/internal/service/middleware"
	"github.com/codolabs/fushon/internal/service/token"
	"github.com/codolabs/fushon/pkg/api/go/auth/authconnect"
	"github.com/codolabs/fushon/pkg/db"
)

func main() {
	var (
		ctx          = context.Background()
		configs      = configs.GetConfigs()
		dbCfg        = configs.DB
		listenPort   = configs.Server.ListenPort
		tokenService = token.NewService(configs)
	)

	// init logger
	zapLogger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("Cannot init logger: %s", err.Error())
	}
	defer func() {
		if err = zapLogger.Sync(); err != nil {
			log.Fatalf("Cannot sync logger: %s", err.Error())
		}
	}()

	logger := zapLogger.Sugar()

	// init db connection
	conn, err := pgx.Connect(ctx, fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s %s",
		dbCfg.DBDomain, dbCfg.DBPort, dbCfg.DBUsername, dbCfg.DBName, dbCfg.DBPassword, dbCfg.DBSSLMode,
	))
	if err != nil {
		logger.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close(ctx)

	// init queries
	queries := db.New(conn)
	repository := repository.New(queries, conn)

	// init gcpbucket service
	// _, err = gcpbucket.NewService(ctx, configs)
	// if err != nil {
	// 	logger.Fatalf("Cannot init gcpbucket service: %s", err.Error())
	// }

	// add middlewares
	middleWare := middleware.NewService(configs, logger, tokenService, repository)
	httpMiddleWare := middleWare.Http
	grpcInterceptor := middleWare.Grpc
	mux := http.NewServeMux()

	// init auth server
	auth := auth.NewServer(logger, tokenService, repository)
	authPath, authHandler := authconnect.NewAuthServiceHandler(
		auth,
		connect.WithInterceptors(grpcInterceptor.ValidateRequestInterceptor()),
	)
	mux.Handle(
		authPath,
		httpMiddleWare.WithCORS(
			httpMiddleWare.WithAuth(authHandler),
		),
	)

	// reflector for services
	reflector := grpcreflect.NewStaticReflector(
		authconnect.AuthServiceName,
	)
	// if configs.Environment == environment.Local {
	mux.Handle(grpcreflect.NewHandlerV1(reflector))
	mux.Handle(grpcreflect.NewHandlerV1Alpha(reflector))
	// }

	fmt.Printf("\nServer started and listen on port: %d\n", listenPort)

	// serve
	err = http.ListenAndServe(
		fmt.Sprintf("0.0.0.0:%d", listenPort),
		h2c.NewHandler(mux, &http2.Server{}),
	)

	if err != nil {
		logger.Fatalf("Cannot serve: %s", err.Error())
	}
}
