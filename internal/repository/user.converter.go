package repository

import "Uploader/internal/service"

func toUserEntity(req service.SignupRequest) UserEntity {
	return UserEntity{
		Email:          req.Email,
		HashedPassword: req.Password,
	}
}

func toSvcUserEntity(user UserEntity) service.UserEntity {
	return service.UserEntity{
		Id:             user.ID,
		Email:          user.Email,
		HashedPassword: user.HashedPassword,
		RefreshToken:   user.RefreshToken,
	}
}
