package mocks

import (
	"context"
	"io"
	"time"

	"github.com/stretchr/testify/mock"
)

// MockStorageRepository is a mock implementation of StorageRepository
type MockStorageRepository struct {
	mock.Mock
}

func (m *MockStorageRepository) UploadFile(ctx context.Context, bucketName, objectName string, reader io.Reader, objectSize int64, contentType string) error {
	args := m.Called(ctx, bucketName, objectName, reader, objectSize, contentType)
	return args.Error(0)
}

func (m *MockStorageRepository) CopyFile(ctx context.Context, srcBucket, srcObject, destBucket, destObject string) error {
	args := m.Called(ctx, srcBucket, srcObject, destBucket, destObject)
	return args.Error(0)
}

func (m *MockStorageRepository) DeleteFile(ctx context.Context, bucketName, objectName string) error {
	args := m.Called(ctx, bucketName, objectName)
	return args.Error(0)
}

func (m *MockStorageRepository) GeneratePresignedURL(ctx context.Context, bucketName, objectName string, expires time.Duration) (string, error) {
	args := m.Called(ctx, bucketName, objectName, expires)
	return args.String(0), args.Error(1)
}

func (m *MockStorageRepository) DownloadFile(ctx context.Context, bucket, objectName string) (io.ReadCloser, error) {
	args := m.Called(ctx, bucket, objectName)
	return args.Get(0).(io.ReadCloser), args.Error(1)
}
