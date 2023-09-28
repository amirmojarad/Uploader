package controller

import (
	"Uploader/conf"
	jwt_handler "Uploader/internal/jwt-handler"
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
	SetRefreshToken(ctx context.Context, req service.UpdateRefreshTokenRequest) error
}

type Auth struct {
	svc        AuthService
	cfg        *conf.AppConfig
	logger     *logrus.Entry
	jwtService *jwt_handler.Jwt
}

func NewAuth(svc AuthService, cfg *conf.AppConfig, logger *logrus.Entry, jwtService *jwt_handler.Jwt) *Auth {
	return &Auth{
		svc:        svc,
		cfg:        cfg,
		logger:     logger,
		jwtService: jwtService,
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

	tokens, err := a.getTokens(ctx, response.Id)
	if err != nil {
		WriteErrorResponse(ctx, err, a.logger)

		return
	}

	response.Tokens = toSvcTokens(tokens)

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

	tokens, err := a.getTokens(ctx, response.Id)
	if err != nil {
		WriteErrorResponse(ctx, err, a.logger)

		return
	}

	response.Tokens = toSvcTokens(tokens)

	ctx.JSON(http.StatusCreated, toViewSignupResponse(response))
}

func (a Auth) getTokens(ctx context.Context, userId uint) (Tokens, error) {
	refreshToken, err := a.jwtService.GenerateJWT(userId)

	if err != nil {
		return Tokens{}, err
	}

	if err := a.svc.SetRefreshToken(ctx, toSvcUpdateRefreshTokenRequest(userId, refreshToken)); err != nil {
		return Tokens{}, err
	}

	accessToken, err := a.jwtService.GenerateJWT(userId)

	if err != nil {
		return Tokens{}, err
	}

	return Tokens{
		accessToken, refreshToken,
	}, nil
}
