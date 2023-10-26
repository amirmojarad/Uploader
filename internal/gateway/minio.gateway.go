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

func (m Minio) Put(ctx context.Context, uid, bucketName string, reader io.Reader, objectSize int) error {
	_, err := m.client.PutObject(ctx,
		bucketName,
		uid,
		reader,
		int64(objectSize),
		minio.PutObjectOptions{},
	)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (m Minio) BucketExists(ctx context.Context, bucketName string) (bool, error) {
	return m.client.BucketExists(ctx, bucketName)
}

func (m Minio) CreateBucket(ctx context.Context, bucketName string) error {
	return m.client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
}

func (m Minio) GetFile(ctx context.Context, bucketName, fileName string) (*minio.Object, error) {
	return m.client.GetObject(ctx, bucketName, fileName, minio.GetObjectOptions{})
}
