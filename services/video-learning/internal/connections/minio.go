package connections

import (
	"fmt"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/video-learning/config"
)

type ConnectMinioConfig interface {
	GetMinioConfig() *config.MinioConfig
}

func ConnectMinio(cfg ConnectMinioConfig) (*minio.Client, error) {
	minioConfig := cfg.GetMinioConfig()

	// Initialize MinIO client
	client, err := minio.New(minioConfig.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(minioConfig.AccessKey, minioConfig.SecretKey, ""),
		Secure: false,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create MinIO client: %w", err)
	}

	// Check if the connection is successful by listing buckets
	_, err = client.ListBuckets(minioConfig.Ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list MinIO buckets: %w", err)
	}

	return client, nil
}
