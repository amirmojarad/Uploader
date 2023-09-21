package controller

import (
	"Uploader/conf"
	"Uploader/internal/service"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

const (
	FORM_FILE_NAME = "file"
)

type UploaderService interface {
	UploadFile(ctx context.Context, request service.UploadFileRequest) (service.UploadFileResponse, error)
}

type Uploader struct {
	svc    UploaderService
	cfg    *conf.AppConfig
	logger *logrus.Entry
}

func (u Uploader) UploadFile(ctx *gin.Context) {
	var req UploadFileRequest

	multipartFile, err := ctx.FormFile(FORM_FILE_NAME)
	if err != nil {
		WriteBindingErrorResponse(ctx, err)

		return
	}

	file, err := multipartFile.Open()
	if err != nil {
		WriteErrorResponse(ctx, err, u.logger)

		return
	}

	req.File = &file

	response, err := u.svc.UploadFile(ctx, toSvcUploadFileRequest(req))
	if err != nil {
		WriteErrorResponse(ctx, err, u.logger)

		return
	}

	ctx.JSON(http.StatusCreated, toViewUploadFileResponse(response))
}
