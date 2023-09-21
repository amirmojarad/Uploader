package service

type UserEntity struct {
	Id             uint
	Email          string
	HashedPassword string
	RefreshToken   string
}

type Tokens struct {
	AccessToken  string
	RefreshToken string
}

type LoginRequest struct {
	Email    string
	Password string
}

type LoginResponse struct {
	Id uint
	Tokens
}

type SignupRequest struct {
	Email    string
	Password string
}

type SignupResponse struct {
	Id uint
	Tokens
}

type UpdateRefreshTokenRequest struct {
	UserId       uint
	RefreshToken string
}

type UpdateRefreshTokenResponse struct {
}
