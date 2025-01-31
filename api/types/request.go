package types

// 首先我要构建结构体，方便后续使用BindJSON()来解析请求中的数据
type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UpdatePasswordRequest struct{
	NewPassword string `json:"newPassword"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}