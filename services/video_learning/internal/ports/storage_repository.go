package ports

import (
	"context"
	"io"
	"time"
)

type StorageRepository interface {
	UploadFile(ctx context.Context, bucket string, objectName string, reader io.Reader, objectSize int64, contentType string) error
	DeleteFile(ctx context.Context, bucket string, objectName string) error
	GeneratePresignedURL(ctx context.Context, bucket string, objectName string, expiry time.Duration) (string, error)
	// maybe dont use this one
	DownloadFile(ctx context.Context, bucket string, objectName string) (io.ReadCloser, error)
}
