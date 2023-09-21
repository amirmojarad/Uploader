package service

import (
	"Uploader/conf"
	"context"
	"github.com/sirupsen/logrus"
)

type Uploader struct {
	cfg    *conf.AppConfig
	logger *logrus.Entry
}

func (u Uploader) UploadFile(ctx context.Context, request UploadFileRequest) (UploadFileResponse, error) {
	return UploadFileResponse{}, nil
}
