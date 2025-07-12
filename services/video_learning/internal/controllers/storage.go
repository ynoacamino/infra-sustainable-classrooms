package controllers

import (
	"context"
	"io"
	"net/url"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/video_learning/internal/repositories"
)

type storageRepository struct {
	client     *minio.Client
	bucketName string
}

func NewStorageController(client *minio.Client, bucketName string) repositories.StorageRepository {
	return &storageRepository{
		client:     client,
		bucketName: bucketName,
	}
}

func (s *storageRepository) UploadFile(ctx context.Context, objectName string, reader io.Reader, objectSize int64, contentType string) error {
	opts := minio.PutObjectOptions{
		ContentType: contentType,
	}

	_, err := s.client.PutObject(ctx, s.bucketName, objectName, reader, objectSize, opts)
	return err
}

func (s *storageRepository) DeleteFile(ctx context.Context, objectName string) error {
	opts := minio.RemoveObjectOptions{}
	return s.client.RemoveObject(ctx, s.bucketName, objectName, opts)
}

func (s *storageRepository) GeneratePresignedURL(ctx context.Context, objectName string, expiry time.Duration) (string, error) {
	reqParams := make(url.Values)
	presignedURL, err := s.client.PresignedGetObject(ctx, s.bucketName, objectName, expiry, reqParams)
	if err != nil {
		return "", err
	}
	return presignedURL.String(), nil
}

func (s *storageRepository) DownloadFile(ctx context.Context, objectName string) (io.ReadCloser, error) {
	opts := minio.GetObjectOptions{}
	object, err := s.client.GetObject(ctx, s.bucketName, objectName, opts)
	if err != nil {
		return nil, err
	}
	return object, nil
}
