package service

import (
	"WinterDemo/configs"
	"WinterDemo/dao"
	myjwt "WinterDemo/pkg/jwt"
	"WinterDemo/api/types"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"strings"
)

func Register(username, password string) error {
	if _,err := dao.GetUser(username); err == nil {
		return fmt.Errorf("用户%s已存在", username)
	}

	return dao.CreateUser(username, password) //return会执行函数
}

func Login(username, password string) (types.TokenResponse, error) {
	user,err := dao.GetUser(username)
	if err != nil {
		return types.TokenResponse{}, fmt.Errorf("用户%s不存在", username)
	}

	if user.Password != password {
		return types.TokenResponse{}, fmt.Errorf("密码错误")
	}

	token,refreshToken, err := myjwt.GenerateToken(user.Username)
	if err != nil {
		return types.TokenResponse{}, fmt.Errorf("生成token失败: %s", err)
	}

	return types.TokenResponse{
		Token: token,
		RefreshToken: refreshToken,
	}, nil
}

// 根据refresh_token刷新access_token
func RefreshToken(refreshToken string) (types.TokenResponse, error) {

	// 同middleware/auth.go中的JWTauth()判断refresh_token是否以Bearer开头
	if !strings.HasPrefix(refreshToken, "Bearer ") {
        return types.TokenResponse{}, fmt.Errorf("无效的token格式")
    }
    refreshToken = refreshToken[7:]

	// 解析refresh_token来获得claims
    claims := &myjwt.Claims{}
    token, err := jwt.ParseWithClaims(refreshToken, claims, func(token *jwt.Token) (interface{}, error) {
        return []byte(configs.GlobalConfig.Server.JWTSecret), nil
    })

	if err != nil || !token.Valid || claims.TokenType != "refresh_token" {
		return types.TokenResponse{}, fmt.Errorf("无效的refresh_token: %s", err)
	}

	accessToken, newRefreshToken, err := myjwt.GenerateToken(claims.Username)
	if err != nil {
		return types.TokenResponse{}, fmt.Errorf("刷新token失败: %s", err)
	}

	return types.TokenResponse{
		Token: accessToken,
		RefreshToken: newRefreshToken,
	}, nil
}

func UpdatePassword(username, newPassword string) error {
	return dao.UpdatePassword(username, newPassword)
}
