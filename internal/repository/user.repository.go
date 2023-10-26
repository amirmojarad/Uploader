package repository

import (
	"Uploader/internal/service"
	"context"
	"gorm.io/gorm"
)

type User struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) *User {
	return &User{db: db}
}

func (u User) CreateUser(ctx context.Context, req service.SignupRequest) (service.SignupResponse, error) {
	userEntity := toUserEntity(req)

	if err := u.db.WithContext(ctx).Create(&userEntity).Error; err != nil {
		return service.SignupResponse{}, err
	}

	return service.SignupResponse{
		Id: userEntity.ID,
	}, nil
}

func (u User) GetUser(ctx context.Context, res service.LoginRequest) (service.LoginResponse, error) {
	return service.LoginResponse{}, nil
}

func (u User) GetByEmail(ctx context.Context, email string) (service.UserEntity, error) {
	var userEntity UserEntity

	if err := u.db.WithContext(ctx).Where("email = ?", email).Find(&userEntity).Error; err != nil {
		return service.UserEntity{}, err
	}

	return toSvcUserEntity(userEntity), nil
}

func (u User) GetUserById(ctx context.Context, id uint) (UserEntity, error) {
	var userEntity UserEntity

	if err := u.db.WithContext(ctx).Where("id = ?", id).Find(&userEntity).Error; err != nil {
		return UserEntity{}, err
	}

	return userEntity, nil
}

func (u User) UpdateRefreshToken(ctx context.Context, req service.UpdateRefreshTokenRequest) error {
	userEntity, err := u.GetUserById(ctx, req.UserId)

	if err != nil {
		return err
	}

	userEntity.RefreshToken = req.RefreshToken

	if err := u.db.WithContext(ctx).Save(&userEntity).Error; err != nil {
		return err
	}

	return err
}
