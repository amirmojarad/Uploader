package service

import (
	"Uploader/conf"
	"Uploader/internal/gateway"
	"context"
	"github.com/sirupsen/logrus"
)

type FileRepository interface {
	Add(ctx context.Context, req AddFileRequest) (AddFileResponse, error)
	GetById(ctx context.Context, req GetFileByIdRequest) (GetFileByIdResponse, error)
	GetAllByUserId(ctx context.Context, req GetAllFilesByIdRequest) (GetAllFilesByIdResponse, error)
	GetByUserId(ctx context.Context, req GetFileByUserIdRequest) (GetFileByUserIdResponse, error)
	CountFiles(ctx context.Context, userId uint) (int64, error)
	ListFiles(ctx context.Context, req ListFilesRequest) (ListFilesResponse, error)
}

type BucketRepository interface {
	Add(ctx context.Context, req AddBucketRequest) (AddBucketResponse, error)
	GetById(ctx context.Context, req GetBucketByIdRequest) (GetBucketByIdResponse, error)
	GetByUserId(ctx context.Context, req GetBucketByUserIdRequest) (GetBucketByUserIdResponse, error)
	Delete(ctx context.Context, req DeleteBucketRequest) error
	IsExists(ctx context.Context, bucketName string, userId uint) (bool, error)
	CreateBucketIfNotExists(ctx context.Context, req CreateBucketIfNotExistsRequest) error
}

type Uploader struct {
	cfg              *conf.AppConfig
	logger           *logrus.Entry
	minio            *gateway.Minio
	bucketRepository BucketRepository
	fileRepository   FileRepository
}

func NewUploader(cfg *conf.AppConfig,
	logger *logrus.Entry,
	minio *gateway.Minio,
	bucketRepository BucketRepository,
	fileRepository FileRepository) *Uploader {
	return &Uploader{
		cfg:              cfg,
		logger:           logger,
		minio:            minio,
		bucketRepository: bucketRepository,
		fileRepository:   fileRepository,
	}
}

func (u Uploader) UploadFile(ctx context.Context, request UploadFileRequest) (UploadFileResponse, error) {
	return UploadFileResponse{}, nil
}

func (u Uploader) GetFileById(ctx context.Context, req GetFileByIdRequest) (GetFileByIdResponse, error) {
	userBucket, err := u.bucketRepository.GetByUserId(ctx, GetBucketByUserIdRequest{UserId: req.UserId})

	if err != nil {
		return GetFileByIdResponse{}, err
	}

	userFile, err := u.fileRepository.GetByUserId(ctx, GetFileByUserIdRequest{UserId: req.UserId})
	if err != nil {
		return GetFileByIdResponse{}, err
	}

	obj, err := u.minio.GetFile(ctx, userBucket.Name, userFile.Name)

	if err != nil {
		return GetFileByIdResponse{}, err
	}

	fileBuff, err := getFileFromMinioObject(obj)

	return GetFileByIdResponse{
		fileBuff,
	}, nil
}

func (u Uploader) ListFiles(ctx context.Context, request ListFilesRequest) (ListFilesResponse, error) {
	count, err := u.fileRepository.CountFiles(ctx, request.UserId)

	if err != nil {
		return ListFilesResponse{}, err
	}

	request.Pagination.Total = int(count)

	return u.fileRepository.ListFiles(ctx, request)
}
