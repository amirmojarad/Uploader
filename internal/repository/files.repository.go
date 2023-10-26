package repository

import (
	"Uploader/internal/service"
	"context"
	"gorm.io/gorm"
)

type File struct {
	db *gorm.DB
}

func NewFile(db *gorm.DB) *File {
	return &File{db: db}
}

func (f File) Add(ctx context.Context, req service.AddFileRequest) (service.AddFileResponse, error) {

	entity := toFileEntity(req)

	if err := f.db.WithContext(ctx).Create(&entity).Error; err != nil {
		return service.AddFileResponse{}, err
	}
	return service.AddFileResponse{
		toSvcFileEntity(entity),
	}, nil
}

func (f File) GetById(ctx context.Context, req service.GetFileByIdRequest) (service.GetFileByIdResponse, error) {
	return service.GetFileByIdResponse{}, nil
}

func (f File) GetAllByUserId(ctx context.Context, req service.GetAllFilesByIdRequest) (service.GetAllFilesByIdResponse, error) {
	return service.GetAllFilesByIdResponse{}, nil
}

func (f File) GetByUserId(ctx context.Context, req service.GetFileByUserIdRequest) (service.GetFileByUserIdResponse, error) {
	return service.GetFileByUserIdResponse{}, nil
}

func (f File) CountFiles(ctx context.Context, userId uint) (int64, error) {
	var count int64

	if err := f.db.WithContext(ctx).Where("user_id = ?", userId).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (f File) ListFiles(ctx context.Context, req service.ListFilesRequest) (service.ListFilesResponse, error) {
	query := f.db.WithContext(ctx).Table("files").Where("user_id = ?", req.UserId)

	if !req.From.IsZero() {
		query = query.Where("created_at > ?", req.From)
	}

	if !req.To.IsZero() {
		query = query.Where("created_at < ?", req.To)
	}

	query = query.Limit(req.Pagination.Size).Limit(req.Pagination.Index)

	var files []FileEntity

	if err := query.Find(&files).Error; err != nil {
		return service.ListFilesResponse{}, err
	}

	return service.ListFilesResponse{
		Result:     toSvcFileEntities(files),
		Pagination: req.Pagination,
	}, nil
}
