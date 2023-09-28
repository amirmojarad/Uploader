package service

import (
	"Uploader/conf"
	"Uploader/internal/gateway"
	"context"
	"github.com/sirupsen/logrus"
)

type Uploader struct {
	cfg    *conf.AppConfig
	logger *logrus.Entry
	minio  *gateway.Minio
}

func NewUploader(cfg *conf.AppConfig, logger *logrus.Entry, minio *gateway.Minio) *Uploader {
	return &Uploader{
		cfg:    cfg,
		logger: logger,
		minio:  minio,
	}
}

func (u Uploader) UploadFile(ctx context.Context, request UploadFileRequest) (UploadFileResponse, error) {
	return UploadFileResponse{}, nil
}
