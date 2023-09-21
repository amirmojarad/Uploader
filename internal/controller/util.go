package controller

import (
	"Uploader/internal/errorext"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func WriteErrorResponse(ctx *gin.Context, err error, logger *logrus.Entry) {
	apiErr := errorext.ToApiError(err)
	logger.WithField("path", ctx.Request.URL.Path).
		WithField("status", apiErr.StatusCode).
		WithField("stackTrace", apiErr.StackTrace()).
		Error(apiErr.Error())

	ctx.JSON(apiErr.StatusCode, apiErr)
}

func WriteBindingErrorResponse(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
}

func WriteNotFoundErrorResponse(ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, gin.H{})
}
