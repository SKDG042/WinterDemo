package dao

import (
	"fmt"
	"WinterDemo/models"
)

func CreateOrder(order *models.Order) error {
	tx := DB.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Create(order).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("创建订单失败: %s", err)
	}

	// 创建订单后，需要清空购物车
	if err := tx.Where("user_id = ?", order.UserID).Delete(&models.Cart{}).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("清空购物车失败: %s", err)
	}

	tx.Commit()
	return nil
}