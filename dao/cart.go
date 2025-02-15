package dao

import (
	"WinterDemo/models"
	"fmt"
)

func AddCart (userID uint, productID uint, quantity int) error {
	var cart models.Cart

	// 首先检查购物车中是否存在该商品
	result := DB.Where("user_id = ? AND product_id = ?", userID, productID).First(&cart)
	if result.Error == nil {
		// 如果购物车中存在刚商品，则数量加1
		cart.Quantity += quantity
		if err := DB.Update("quantity", cart.Quantity).Error; err != nil {
			return fmt.Errorf("更新购物车失败: %s", err)
		}
		return nil
	}

	// 如果购物车中不存在该商品，则加入
	newCart := models.Cart{
		UserID: userID,
		ProductID:productID,
		Quantity: quantity,
	}

	if err := DB.Create(&newCart).Error; err != nil {
		return fmt.Errorf("添加购物车失败: %s", err)
	}

	return nil
}

func GetCartByUserID(userID uint) ([]models.Cart, error) {
	var cart []models.Cart

	if err := DB.Preload("Product").Where("user_id = ?", userID).Find(&cart).Error; err != nil {
		return nil, fmt.Errorf("您还没有添加任何商品到购物车: %s", err)
	}

	return cart, nil
}