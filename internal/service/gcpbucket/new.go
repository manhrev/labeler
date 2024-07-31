package gcpbucket

import (
	"context"
	"time"

	"cloud.google.com/go/storage"
	"github.com/codolabs/fushon/configs"
	"github.com/codolabs/fushon/internal/const/environment"
	"google.golang.org/api/option"
)

type Service interface {
	GenerateSignedURL(bucketName, fileName, contentType string, expiredDuration time.Duration) (string, error)
	GenerateSignedURLWithLimitContentType(
		bucketName, fileName, contentType string, expiredDuration time.Duration, allowedContentTypes map[string]bool, sizeLimit uint64,
	) (string, map[string]string, error)
	GenerateSignedURLWithLimit(
		bucketName, fileName, contentType string, expiredDuration time.Duration, allowedContentTypes map[string]bool, sizeLimit uint64,
	) (string, map[string]string, error)
}

type service struct {
	client *storage.Client
}

func NewService(ctx context.Context, configs *configs.Config) (Service, error) {
	if configs.Environment == environment.Local {
		return &localService{}, nil
	}

	client, err := storage.NewClient(ctx, option.WithCredentialsJSON(configs.GCPBucket.JsonCredentials))
	if err != nil {
		return nil, err
	}

	return &service{client: client}, nil
}
