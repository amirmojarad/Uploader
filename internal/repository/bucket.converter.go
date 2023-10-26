package repository

import (
	"Uploader/internal/service"
)

func toSvcBucketEntity(bucket BucketEntity) service.BucketEntity {
	return service.BucketEntity{
		Id:     bucket.ID,
		Name:   bucket.Name,
		UserID: bucket.UserID,
	}
}

func toBucketEntity(req service.AddBucketRequest) BucketEntity {
	return BucketEntity{
		Name:   req.BucketName,
		UserID: req.UserId,
	}
}
