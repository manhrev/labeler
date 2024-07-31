package gcpbucket

import (
	"fmt"
	"net/http"
	"time"

	"cloud.google.com/go/storage"
)

func (s *service) GenerateSignedURL(
	bucketName, fileName, contentType string, expiredDuration time.Duration,
) (string, error) {
	return s.client.Bucket(bucketName).
		SignedURL(
			fileName,
			&storage.SignedURLOptions{
				Method:      http.MethodPut,
				Expires:     time.Now().Add(expiredDuration),
				Scheme:      storage.SigningSchemeV4,
				ContentType: contentType,
			},
		)
}

func (s *service) GenerateSignedURLWithLimitContentType(
	bucketName, fileName, contentType string, expiredDuration time.Duration, allowedContentTypes map[string]bool,
	sizeLimit uint64,
) (string, map[string]string, error) {
	if ok := allowedContentTypes[contentType]; !ok {
		return "", nil, fmt.Errorf("content-type(%s) is not allow", contentType)
	}

	if sizeLimit == 0 {
		// Default limit 5MB per file
		sizeLimit = 5_000_000
	}

	token, err := s.client.Bucket(bucketName).
		SignedURL(
			fileName,
			&storage.SignedURLOptions{
				Method:      http.MethodPut,
				Expires:     time.Now().Add(expiredDuration),
				Scheme:      storage.SigningSchemeV4,
				ContentType: contentType,
				//Headers: []string{
				//	fmt.Sprintf("X-Goog-Content-Length-Range:0,%d", sizeLimit),
				//	fmt.Sprintf("Content-Type:%s", contentType),
				//},
			},
		)
	if err != nil {
		return "", nil, err
	}

	return token, map[string]string{
		"X-Goog-Content-Length-Range": fmt.Sprintf("0,%d", sizeLimit),
		"Content-Type":                contentType,
	}, nil
}

func (s *service) GenerateSignedURLWithLimit(
	bucketName, fileName, contentType string, expiredDuration time.Duration, allowedContentTypes map[string]bool,
	sizeLimit uint64,
) (string, map[string]string, error) {
	if ok := allowedContentTypes[contentType]; !ok {
		return "", nil, fmt.Errorf("content-type(%s) is not allow", contentType)
	}

	if sizeLimit == 0 {
		// Default limit 5MB per file
		sizeLimit = 5_000_000
	}

	token, err := s.client.Bucket(bucketName).
		SignedURL(
			fileName,
			&storage.SignedURLOptions{
				Method:  http.MethodPut,
				Expires: time.Now().Add(expiredDuration),
				Scheme:  storage.SigningSchemeV4,
				Headers: []string{
					fmt.Sprintf("X-Goog-Content-Length-Range:0,%d", sizeLimit),
					fmt.Sprintf("Content-Type:%s", contentType),
				},
			},
		)
	if err != nil {
		return "", nil, err
	}

	return token, map[string]string{
		"X-Goog-Content-Length-Range": fmt.Sprintf("0,%d", sizeLimit),
		"Content-Type":                contentType,
	}, nil
}
