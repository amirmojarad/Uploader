package controller

import (
	"Uploader/internal/service"
)

func toSvcUploadFileRequest(req UploadFileRequest) service.UploadFileRequest {
	return service.UploadFileRequest{
		File: req.File,
	}
}

func toViewUploadFileResponse(res service.UploadFileResponse) UploadFileResponse {
	return UploadFileResponse{
		File{
			Id:        res.Id,
			Title:     res.Title,
			Volume:    res.Volume,
			CreatedAt: res.CreatedAt,
			Format:    res.Format,
		},
	}
}
