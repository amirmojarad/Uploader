package controller

import (
	"Uploader/conf"
	"Uploader/internal/service"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/sirupsen/logrus"
	"net/http"
)

type AuthService interface {
	Login(ctx context.Context, req service.LoginRequest) (service.LoginResponse, error)
	Signup(ctx context.Context, req service.SignupRequest) (service.SignupResponse, error)
}

type Auth struct {
	svc    AuthService
	cfg    *conf.AppConfig
	logger *logrus.Entry
}

func NewAuth(svc AuthService, cfg *conf.AppConfig, logger *logrus.Entry) *Auth {
	return &Auth{
		svc:    svc,
		cfg:    cfg,
		logger: logger,
	}
}

func (a Auth) Login(ctx *gin.Context) {
	var req LoginRequest

	if err := ctx.MustBindWith(&req, binding.JSON); err != nil {
		WriteBindingErrorResponse(ctx, err)

		return
	}

	response, err := a.svc.Login(ctx.Request.Context(), toSvcLoginRequest(req))

	if err != nil {
		WriteErrorResponse(ctx, err, a.logger)

		return
	}

	ctx.JSON(http.StatusOK, toViewLoginResponse(response))
}

func (a Auth) Signup(ctx *gin.Context) {
	var req SignupRequest

	if err := ctx.MustBindWith(&req, binding.JSON); err != nil {
		WriteBindingErrorResponse(ctx, err)

		return
	}

	response, err := a.svc.Signup(ctx, toSvcSignupRequest(req))

	if err != nil {
		WriteErrorResponse(ctx, err, a.logger)

		return
	}

	ctx.JSON(http.StatusCreated, toViewSignupResponse(response))
}
