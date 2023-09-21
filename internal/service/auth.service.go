package service

import (
	"context"

	"Uploader/conf"
	"Uploader/internal/jwt-handler"
	"github.com/sirupsen/logrus"
)

type AuthRepository interface {
	CreateUser(ctx context.Context, req SignupRequest) (SignupResponse, error)
	GetUser(ctx context.Context, res LoginRequest) (LoginResponse, error)
	UpdateRefreshToken(ctx context.Context, req UpdateRefreshTokenRequest) error
	GetByEmail(ctx context.Context, email string) (UserEntity, error)
}

type Auth struct {
	cfg            *conf.AppConfig
	logger         *logrus.Entry
	authRepository AuthRepository
	jwtService     jwt_handler.Jwt
}

func NewAuth(cfg *conf.AppConfig, logger *logrus.Entry, authRepository AuthRepository, jwtService jwt_handler.Jwt) *Auth {
	return &Auth{
		cfg:            cfg,
		logger:         logger,
		authRepository: authRepository,
		jwtService:     jwtService,
	}
}

func (a Auth) Login(ctx context.Context, req LoginRequest) (LoginResponse, error) {
	if err := checkEmailAndPassword(req.Email, req.Password); err != nil {
		return LoginResponse{}, err
	}

	userEntity, err := a.authRepository.GetByEmail(ctx, req.Email)
	if err != nil {
		return LoginResponse{}, err
	}

	if err := comparePasswords(userEntity.HashedPassword, req.Password); err != nil {
		return LoginResponse{}, err
	}

	refreshToken, err := a.jwtService.GenerateJWT(userEntity.Id)
	if err != nil {
		return LoginResponse{}, err
	}

	if err := a.authRepository.UpdateRefreshToken(ctx,
		toUpdateRefreshTokenRequest(userEntity.Id, refreshToken)); err != nil {
		return LoginResponse{}, err
	}

	accessToken, err := a.jwtService.GenerateJWT(userEntity.Id)
	if err != nil {
		return LoginResponse{}, err
	}

	return LoginResponse{
		Id: userEntity.Id,
		Tokens: Tokens{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	}, nil
}

func (a Auth) Signup(ctx context.Context, req SignupRequest) (SignupResponse, error) {
	if err := checkEmailAndPassword(req.Email, req.Password); err != nil {
		return SignupResponse{}, err
	}

	hashedPassword, err := hashPassword(req.Password)

	if err != nil {
		return SignupResponse{}, err
	}

	req.Password = hashedPassword

	response, err := a.authRepository.CreateUser(ctx, req)

	if err != nil {
		return SignupResponse{}, err
	}

	refreshToken, err := a.jwtService.GenerateJWT(response.Id)

	if err != nil {
		return SignupResponse{}, err
	}

	if err = a.authRepository.UpdateRefreshToken(ctx, UpdateRefreshTokenRequest{
		UserId:       response.Id,
		RefreshToken: refreshToken,
	}); err != nil {
		return SignupResponse{}, err
	}

	accessToken, err := a.jwtService.GenerateJWT(response.Id)

	if err != nil {
		return SignupResponse{}, err
	}

	response.Tokens.RefreshToken = refreshToken
	response.Tokens.AccessToken = accessToken

	return response, nil
}
