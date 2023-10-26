package controller

import (
	"Uploader/internal/service"
)

func toSvcUploadFileRequest(req UploadFileRequest) service.UploadFileRequest {
	return service.UploadFileRequest{
		File:     req.File,
		FileName: req.FileName,
		UserId:   req.UserId,
	}
}

func toViewUploadFileResponse(res service.UploadFileResponse) UploadFileResponse {
	return UploadFileResponse{
		File{
			Id:       res.Id,
			Name:     res.Name,
			UserID:   res.UserID,
			BucketID: res.BucketID,
			Size:     res.Size,
			Type:     res.Type,
		},
	}
}

func toSvcListFilesRequest(req ListFilesRequest) service.ListFilesRequest {
	return service.ListFilesRequest{
		UserId:     req.UserId,
		From:       req.From,
		To:         req.To,
		Pagination: toSvcPagination(req.Pagination),
	}
}

func toViewPagination(pagination service.Pagination) Pagination {
	return Pagination{
		Index: pagination.Index,
		Size:  pagination.Size,
		Total: pagination.Total,
	}
}

func toSvcPagination(pagination Pagination) service.Pagination {
	return service.Pagination{
		Index: pagination.Index,
		Size:  pagination.Size,
	}
}

func toViewListFiles(res service.ListFilesResponse) ListFilesResponse {
	return ListFilesResponse{
		Result:     toViewFiles(res.Result),
		Pagination: toViewPagination(res.Pagination),
	}
}

func toViewFiles(items []service.FileEntity) []File {
	var result []File

	for _, item := range items {
		result = append(result, toViewFile(item))
	}

	return result
}

func toViewFile(item service.FileEntity) File {
	return File{
		Id:       item.Id,
		Name:     item.Name,
		UserID:   item.UserID,
		BucketID: item.BucketID,
		Size:     item.Size,
		Type:     item.Type,
	}
}
