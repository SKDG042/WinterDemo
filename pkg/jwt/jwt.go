package jwt

import (
	"WinterDemo/configs"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Username string `json:"username"`
	UserID	uint	`json:"user_id"`
	//这里的种类主要是为了在刷新token时避免使用access_token执行service.RefreshToken刷新access_token
	TokenType string `json:"token_type"` 
	jwt.RegisteredClaims
}

// 生成token
func GenerateToken(username string) (string, string, error) {

	//生成access_token
	claims := Claims{
		Username: username,
		TokenType: "access_token",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(configs.GlobalConfig.Server.AccessTokenExpire) * time.Minute )), // 有效时间1分钟
			IssuedAt:  jwt.NewNumericDate(time.Now()),  // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),  // 生效时间
			Issuer:    "042",                          // 签发人
		},
	}

	// 依据claims创建一个未签名的access_token
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)	
	// 使用configs中的密钥对access_token进行签名
	accessTokenString, err := accessToken.SignedString([]byte(configs.GlobalConfig.Server.JWTSecret))
	if err != nil {
		return "","",fmt.Errorf("生成access_token失败: %s", err)
	}

	//生成refresh_token 同上
	refreshClaims := Claims{
		Username: username,
		TokenType: "refresh_token",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(configs.GlobalConfig.Server.RefreshTokenExpire) * time.Minute)), // 有效时间1分钟
			IssuedAt:  jwt.NewNumericDate(time.Now()),         // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),         // 生效时间
			Issuer:    "042",                          // 签发人
		},
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshTokenString, err := refreshToken.SignedString([]byte(configs.GlobalConfig.Server.JWTSecret))
	if err != nil {
		return "","",fmt.Errorf("生成refresh_token失败: %s", err)
	}

	return accessTokenString, refreshTokenString, nil
}

// 解析token
func ParseToken(tokenString string) (string, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(configs.GlobalConfig.Server.JWTSecret), nil
	})

	if err != nil {
		return "", fmt.Errorf("token解析错误: %s", err)	
	}

	//检验token是否有效(例如过期,签名错误,格式错误等)
	if !token.Valid {
		return "", fmt.Errorf("token无效: %s", err)
	}

	return claims.Username, nil
}
