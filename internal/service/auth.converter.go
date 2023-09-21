package service

func toUpdateRefreshTokenRequest(userId uint, refreshToken string) UpdateRefreshTokenRequest {
	return UpdateRefreshTokenRequest{
		UserId:       userId,
		RefreshToken: refreshToken,
	}
}
