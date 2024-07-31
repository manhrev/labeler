package gcpbucket

import (
	"time"
)

type localService struct{}

func (s *localService) GenerateSignedURL(
	bucketName, fileName, contentType string, expiredDuration time.Duration,
) (string, error) {
	return "http://localhost", nil
}

func (s *localService) GenerateSignedURLWithLimitContentType(
	bucketName, fileName, contentType string, expiredDuration time.Duration, allowedContentTypes map[string]bool,
	sizeLimit uint64,
) (string, map[string]string, error) {
	return "http://localhost", map[string]string{}, nil
}

func (s *localService) GenerateSignedURLWithLimit(
	bucketName, fileName, contentType string, expiredDuration time.Duration, allowedContentTypes map[string]bool,
	sizeLimit uint64,
) (string, map[string]string, error) {
	return "http://localhost", map[string]string{}, nil
}
