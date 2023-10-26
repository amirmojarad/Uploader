package controller

import (
	"Uploader/conf"
	"Uploader/consts"
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
	ListFiles(ctx context.Context, request service.ListFilesRequest) (service.ListFilesResponse, error)
}

type Uploader struct {
	svc    UploaderService
	cfg    *conf.AppConfig
	logger *logrus.Entry
}

func NewUploader(cfg *conf.AppConfig, logger *logrus.Entry, svc UploaderService) *Uploader {
	return &Uploader{
		svc:    svc,
		cfg:    cfg,
		logger: logger,
	}
}

func (u Uploader) ListFiles(ctx *gin.Context) {
	var req ListFilesRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		WriteBindingErrorResponse(ctx, err)

		return
	}

	req.UserId = uint(ctx.GetInt64(consts.REQUEST_USER_ID))

	response, err := u.svc.ListFiles(ctx, toSvcListFilesRequest(req))

	if err != nil {
		WriteErrorResponse(ctx, err, u.logger)

		return
	}

	ctx.JSON(http.StatusOK, toViewListFiles(response))

	return
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

	if userId, exists := ctx.Get(consts.REQUEST_USER_ID); exists {
		req.UserId = userId.(uint)
	} else {
		WriteErrorResponse(ctx, err, u.logger)
	}

	req.FileName = multipartFile.Filename
	req.File = file

	response, err := u.svc.UploadFile(ctx, toSvcUploadFileRequest(req))
	if err != nil {
		WriteErrorResponse(ctx, err, u.logger)

		return
	}

	ctx.JSON(http.StatusCreated, toViewUploadFileResponse(response))
}
