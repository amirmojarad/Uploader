package gateway

import (
	"context"
	"io"

	minio "github.com/minio/minio-go/v7"
	"github.com/pkg/errors"
)

type Minio struct {
	client     *minio.Client
	bucketName string
}

func NewMinio(client *minio.Client, bucketName string) *Minio {
	return &Minio{
		client:     client,
		bucketName: bucketName,
	}
}

func (m Minio) Put(ctx context.Context, uid string, reader io.Reader, objectSize int) error {
	_, err := m.client.PutObject(ctx, m.bucketName, uid, reader, int64(objectSize), minio.PutObjectOptions{})
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
