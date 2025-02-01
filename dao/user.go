package dao

import (
	"WinterDemo/models"
	"fmt"
)

// 创建用户
func CreateUser(username, password string) error {
	user := models.User{
		Username: username,
		Password: password,
	}

	result := DB.Create(&user) //create()会返回一个*gorm.DB,包含错误信息等
	if result.Error != nil {   //调用*gorm.DB的Error()方法，获取错误信息
		return fmt.Errorf("创建用户%s失败: %v", user.Username, result.Error)
	}

	return nil
}

// 通过查找用户来避免用户名重复
func GetUser(username string) (*models.User, error) {
	var user models.User
	result := DB.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return nil, fmt.Errorf("获取用户%s失败: %v", username, result.Error)
	}
	return &user, nil
}

// 只更新密码
func UpdatePassword(Username string, newPassword string) error {

	if newPassword == "" || len(newPassword) < 6 {
		return fmt.Errorf("新密码不能为空且长度不能小于6")
	}

	user,err := GetUser(Username)
	if err != nil {
		return fmt.Errorf("获取用户信息失败: %v", err)
	}

	if newPassword == user.Password {
		return fmt.Errorf("新密码不能与旧密码相同")
	}

	result := DB.Model(&models.User{}).Where("username = ?", Username).Update("password", newPassword)
	if result.Error != nil {
		return fmt.Errorf("更新用户%s信息失败: %v", Username, result.Error)
	}

	//接着确定修改成功(当修改的用户名不存在时会出现静默失败，即用户不存在时也会返回成功)
	if result.RowsAffected == 0 {
		return fmt.Errorf("用户%s不存在", Username)
	}
	return nil
}

// 更新用户信息
func UpdateUserInfo(username string, updates map[string]interface{}) error {
	result := DB.Model(&models.User{}).Where("username = ?", username).Updates(updates)
	if result.Error != nil {
		return fmt.Errorf("更新用户%s信息失败: %v", username, result.Error)
	}

	//防止静默失败
	if result.RowsAffected == 0 {
		return fmt.Errorf("用户%s不存在", username)
	}

	return nil
}

