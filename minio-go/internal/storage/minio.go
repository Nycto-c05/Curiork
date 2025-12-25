package object

import (
	"context"

	minio "github.com/minio/minio-go/v7"
	// credentials "github.com/minio/minio-go/v7/pkg/credentials"
)

type MinioStore struct {
	client *minio.Client
	bucket string
}

func NewMinioStore(client *minio.Client, bucket string) *MinioStore {
	return &MinioStore{client: client, bucket: bucket}
}

func (s *MinioStore) Get(ctx context.Context, key string) ([]byte, error) {
	return nil, nil
}

func (s *MinioStore) Put(ctx context.Context, key string, data []byte) error {
	return nil
}
