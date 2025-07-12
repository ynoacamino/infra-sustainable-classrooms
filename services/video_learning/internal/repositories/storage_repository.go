package repositories

import (
	"context"
	"io"
	"net/url"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/video_learning/internal/ports"
)

type storageRepository struct {
	client *minio.Client
}

func NewStorageRepository(client *minio.Client) ports.StorageRepository {
	return &storageRepository{
		client: client,
	}
}

func (s *storageRepository) UploadFile(ctx context.Context, bucket string, objectName string, reader io.Reader, objectSize int64, contentType string) error {
	opts := minio.PutObjectOptions{
		ContentType: contentType,
	}

	_, err := s.client.PutObject(ctx, bucket, objectName, reader, objectSize, opts)
	return err
}

func (s *storageRepository) DeleteFile(ctx context.Context, bucket string, objectName string) error {
	opts := minio.RemoveObjectOptions{}
	return s.client.RemoveObject(ctx, bucket, objectName, opts)
}

func (s *storageRepository) GeneratePresignedURL(ctx context.Context, bucket string, objectName string, expiry time.Duration) (string, error) {
	reqParams := make(url.Values)
	presignedURL, err := s.client.PresignedGetObject(ctx, bucket, objectName, expiry, reqParams)
	if err != nil {
		return "", err
	}
	return presignedURL.String(), nil
}

func (s *storageRepository) DownloadFile(ctx context.Context, bucket string, objectName string) (io.ReadCloser, error) {
	opts := minio.GetObjectOptions{}
	object, err := s.client.GetObject(ctx, bucket, objectName, opts)
	if err != nil {
		return nil, err
	}
	return object, nil
}
