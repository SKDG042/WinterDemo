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

func UpdateUser(oldUsername string, newUsername string, newPassword string) error {
	updates := make(map[string]interface{})

	if newUsername == oldUsername {
		newUsername = ""
	}

	if newPassword != "" {
		updates["password"] = newPassword
	}

	if newUsername != "" {
		//首先我们要检查修改后的用户名是否存在
		if _, err := GetUser(newUsername); err == nil {
			return fmt.Errorf("用户名%s已存在", newUsername)
		}
		updates["username"] = newUsername
	}

	//如果操作后updates的len为0，即代表没有修改，所以直接返回
	if len(updates) == 0 {
		return nil
	}

	result := DB.Model(&models.User{}).Where("username = ?", oldUsername).Updates(updates)
	if result.Error != nil {
		return fmt.Errorf("更新用户%s信息失败: %v", oldUsername, result.Error)
	}

	//接着确定修改成功(当修改的用户名不存在时会出现静默失败，即用户不存在时也会返回成功)
	if result.RowsAffected == 0 {
		return fmt.Errorf("用户%s不存在", oldUsername)
	}
	return nil
}
