package dao

import (
	"WinterDemo/models"
	"fmt"
)

func AddCart (userID uint, productID uint, quantity int) error {
	var cart models.Cart
	var product models.Product
	

	// 获取加入购物车的商品
	if err := DB.First(&product, productID).Error; err != nil {
		return fmt.Errorf("该商品不存在: %s", err)
	}

	// 首先检查购物车中是否存在该商品
	result := DB.Where("user_id = ? AND product_id = ?", userID, productID).First(&cart)
	if result.Error == nil {
		// 如果购物车中存在刚商品，则数量加1
		cart.Quantity += quantity
		if err := DB.Model(&cart).Where("id = ?", cart.ID).Update("quantity", cart.Quantity).Error; err != nil {
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

	// 使用事务以保证添加商品和购物车关联的一致性
	tx := DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Create(&newCart).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("添加购物车失败: %s", err)
	}

	// 将商品添加到购物车中时，需要将商品和购物车关联
	if err := tx.Model(&newCart).Association("Product").Append(&product); err != nil {
		tx.Rollback()
		return fmt.Errorf("关联购物车失败: %s", err)
	}

	tx.Commit()
	return nil
}

func GetCartByUserID(userID uint) ([]models.Cart, error) {
	var carts []models.Cart

	if err := DB.Preload("Product").Where("user_id = ?", userID).Find(&carts).Error; err != nil {
		return carts, fmt.Errorf("您还没有添加任何商品到购物车: %s", err)
	}

	return carts, nil
}

func DeleteCart(userID uint) error {
	if err := DB.Where("user_id = ?", userID).Delete(&models.Cart{}).Error; err != nil {
		return fmt.Errorf("删除购物车失败: %s", err)
	}
	return nil
}