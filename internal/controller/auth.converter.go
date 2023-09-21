package controller

import (
	"Uploader/internal/service"
)

func toSvcSignupRequest(req SignupRequest) service.SignupRequest {
	return service.SignupRequest{
		Email:    req.Email,
		Password: req.Password,
	}
}

func toViewSignupResponse(res service.SignupResponse) SignupResponse {
	return SignupResponse{
		Id: res.Id,
		Tokens: Tokens{
			AccessToken:  res.AccessToken,
			RefreshToken: res.RefreshToken,
		},
	}
}

func toSvcLoginRequest(req LoginRequest) service.LoginRequest {
	return service.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
	}
}

func toViewLoginResponse(res service.LoginResponse) LoginResponse {
	return LoginResponse{
		Id: res.Id,
		Tokens: Tokens{
			AccessToken:  res.AccessToken,
			RefreshToken: res.RefreshToken,
		},
	}
}
