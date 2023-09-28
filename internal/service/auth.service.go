package service

import (
	"context"

	"Uploader/conf"
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
}

func NewAuth(cfg *conf.AppConfig, logger *logrus.Entry, authRepository AuthRepository) *Auth {
	return &Auth{
		cfg:            cfg,
		logger:         logger,
		authRepository: authRepository,
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

	return LoginResponse{
		Id: userEntity.Id,
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

	return response, nil
}

func (a Auth) SetRefreshToken(ctx context.Context, req UpdateRefreshTokenRequest) error {
	return a.authRepository.UpdateRefreshToken(ctx, req)
}
