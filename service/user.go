package service

import (
	"WinterDemo/configs"
	"WinterDemo/dao"
	myjwt "WinterDemo/pkg/jwt" // 重命名防止冲突
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
		return types.TokenResponse{}, fmt.Errorf("用户%s不存在", username) //if err !=)
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
func RefreshToken(refreshToken string) (*types.TokenResponse, error) {

	// 同middleware/auth.go中的JWTauth()判断refresh_token是否以Bearer开头
	if !strings.HasPrefix(refreshToken, "Bearer ") {
        return nil, fmt.Errorf("无效的token格式")
    }
    refreshToken = refreshToken[7:]

	// 解析refresh_token来获得claims
    claims := &myjwt.Claims{}
    token, err := jwt.ParseWithClaims(refreshToken, claims, func(token *jwt.Token) (interface{}, error) {
        return []byte(configs.GlobalConfig.Server.JWTSecret), nil
    })

	// 判断token是否有效，claims.TokenType是否为refresh_token
	if err != nil || !token.Valid || claims.TokenType != "refresh_token" {
		return nil, fmt.Errorf("无效的refresh_token: %s", err)
	}

	// 生成新的access_token和refresh_token
	accessToken, newRefreshToken, err := myjwt.GenerateToken(claims.Username)
	if err != nil {
		return nil, fmt.Errorf("刷新token失败: %s", err)
	}

	tokenResponse := types.TokenResponse{
		Token: accessToken,
		RefreshToken: newRefreshToken,
	}
	return &tokenResponse, nil
}

func GetUserInfo(username string) (*types.UserInfoResponse, error) {
	user, err := dao.GetUser(username)
	if err != nil {
		return nil, fmt.Errorf("获取用户信息失败: %v", err)
	}

	userResponse := types.UserInfoResponse{
		ID: user.ID,
		Avatar: user.Avatar,
		Nickname: user.Nickname,
		Introduction: user.Introduction,
		Phone: user.Telephone,
		QQ: user.QQ,
		Gender: user.Gender,
		Email: user.Email,
		Birthday: user.Birthday,
	}
	return &userResponse, nil
}

func UpdatePassword(username, newPassword string) error {
	return dao.UpdatePassword(username, newPassword)
}

func UpdateUserInfo(username string, info types.UpdateUserInfoRequest) error {
	// 创建一个map来存储需要更新的字段
	updates := make(map[string]interface{})

	//如果info的字段值不为空，则将字段和值存入map
	if info.Nickname != "" {
		updates["nickname"] = info.Nickname
	}
	if info.Avatar != "" {
		updates["avatar"] = info.Avatar
	}
	if info.Introduction != "" {
		updates["introduction"] = info.Introduction
	}
	if info.Telephone != "" {
		updates["telephone"] = info.Telephone
	}
	if info.QQ != "" {
		updates["qq"] = info.QQ
	}
	if info.Gender != "" {
		updates["gender"] = info.Gender
	}
	if info.Email != "" {
		updates["email"] = info.Email
	}
	if info.Birthday != "" {
		updates["birthday"] = info.Birthday
	}
	if info.Telephone != "" {
		updates["telephone"] = info.Telephone
	}

	//如果map为空，则没有需要更新的信息
	if len(updates) == 0 {
		return fmt.Errorf("没有需要更新的信息")
	}

	// 调用dao.UpdateUserInfo() 更新用户信息
	return dao.UpdateUserInfo(username, updates)
}
