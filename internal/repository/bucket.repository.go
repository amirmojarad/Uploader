package repository

import (
	"Uploader/internal/service"
	"context"
	"gorm.io/gorm"
	"time"
)

type Bucket struct {
	db *gorm.DB
}

func NewBucket(db *gorm.DB) *Bucket {
	return &Bucket{db: db}
}

func (b Bucket) Add(ctx context.Context, req service.AddBucketRequest) (service.AddBucketResponse, error) {
	entity := toBucketEntity(req)

	if err := b.db.WithContext(ctx).Create(&entity).Error; err != nil {
		return service.AddBucketResponse{}, err
	}

	return service.AddBucketResponse{
		BucketEntity: toSvcBucketEntity(entity),
	}, nil
}

func (b Bucket) GetById(ctx context.Context, req service.GetBucketByIdRequest) (service.GetBucketByIdResponse, error) {
	return service.GetBucketByIdResponse{}, nil
}

func (b Bucket) GetByUserId(ctx context.Context, req service.GetBucketByUserIdRequest) (service.GetBucketByUserIdResponse, error) {
	var entity BucketEntity

	if err := b.db.WithContext(ctx).First(&entity, req.UserId).Error; err != nil {
		return service.GetBucketByUserIdResponse{}, err
	}

	return service.GetBucketByUserIdResponse{BucketEntity: toSvcBucketEntity(entity)}, nil
}

func (b Bucket) Delete(ctx context.Context, req service.DeleteBucketRequest) error {
	return b.db.WithContext(ctx).
		Table(BucketEntity{}.TableName()).
		Where("id = ?", req.Id).
		Where("user_id = ?", req.UserId).
		Update("deleted_at = ?", time.Now()).Error
}

func (b Bucket) IsExists(ctx context.Context, bucketName string, userId uint) (bool, error) {
	var bucket BucketEntity

	if err := b.db.WithContext(ctx).
		Where("name = ? and user_id = ?", bucketName, userId).
		Find(&bucket).Error; err != nil {
		return false, err
	}

	return true, nil
}

func (b Bucket) CreateBucketIfNotExists(ctx context.Context, req service.CreateBucketIfNotExistsRequest) error {
	isExists, err := b.IsExists(ctx, req.BucketName, req.UserId)
	if err != nil {
		return err
	}

	if !isExists {
		_, err = b.Add(ctx, service.AddBucketRequest(req))
		if err != nil {
			return nil
		}
		return nil
	}

	return nil
}
