package service

import (
	"WinterDemo/dao"
	"WinterDemo/middleware"
	"fmt"
)

func Register(username, password string) error {
	if _,err := dao.GetUser(username); err == nil {
		return fmt.Errorf("用户%s已存在", username)
	}

	return dao.CreateUser(username, password) //return会执行函数
}

func Login(username, password string) (string, error) {
	user,err := dao.GetUser(username)
	if err != nil {
		return "", fmt.Errorf("用户%s不存在", username)
	}

	if user.Password != password {
		return "", fmt.Errorf("密码错误")
	}

	token, err := middleware.GenerateToken(user.Username)
	if err != nil {
		return "", fmt.Errorf("生成token失败")
	}

	return token, nil
}

func UpdateUser(username, newUsername, newPassword string) error {
	return dao.UpdateUser(username, newUsername, newPassword)
}
