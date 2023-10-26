package repository

import (
	"Uploader/internal/service"
)

func toSvcFileEntities(entities []FileEntity) []service.FileEntity {
	var result []service.FileEntity

	for _, item := range entities {
		result = append(result, toSvcFileEntity(item))
	}

	return result
}

func toSvcFileEntity(entity FileEntity) service.FileEntity {
	return service.FileEntity{
		Id:       entity.ID,
		Name:     entity.Name,
		UserID:   entity.UserID,
		BucketID: entity.BucketID,
		Size:     entity.Size,
		Type:     entity.Type,
	}
}

func toFileEntity(req service.AddFileRequest) FileEntity {
	return FileEntity{
		Name:     req.Name,
		UserID:   req.UserID,
		BucketID: req.BucketID,
		Size:     req.Size,
		Type:     req.Type,
	}
}
