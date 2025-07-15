package ports

import (
	"context"
	"io"
	"mime/multipart"
	"time"
)

type StorageRepository interface {
	UploadFile(ctx context.Context, bucket string, objectName string, part *multipart.Part, contentType string) error
	DeleteFile(ctx context.Context, bucket string, objectName string) error
	GeneratePresignedURL(ctx context.Context, bucket string, objectName string, expiry time.Duration) (string, error)
	CopyFile(ctx context.Context, srcBucket string, srcObject string, destBucket string, destObject string) error
	// maybe dont use this one
	DownloadFile(ctx context.Context, bucket string, objectName string) (io.ReadCloser, error)
}
