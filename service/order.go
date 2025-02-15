package service

import (
	"WinterDemo/api/types"
	"WinterDemo/dao"
	"WinterDemo/models"
	"fmt"
)

func convertToOrderProductResponse(product []models.Product, quantity int) []types.OrderProductResponse {
	var OrderProductResponse []types.OrderProductResponse
	for _, product := range product {
		OrderProductResponse = append(OrderProductResponse, types.OrderProductResponse{
			ProductID:	product.ID,
			Name:		product.Name,
			Price:		product.Price,
			Quantity:	quantity,
		})
	}
	return OrderProductResponse
}

func convertToOrderResponse(product models.Product, quantity int) []types.OrderProductResponse {
    var OrderProductResponse []types.OrderProductResponse
    OrderProductResponse = append(OrderProductResponse, types.OrderProductResponse{
        ProductID: product.ID,
        Name:     product.Name,
        Price:    product.Price,
        Quantity: quantity,
    })
    return OrderProductResponse
}

func CreateOrder(username string) error {
	user, err := GetUserInfo(username)
	if err != nil {
		return fmt.Errorf("获取用户信息失败: %s", err)
	}

	carts, err := dao.GetCartByUserID(user.ID)
	if err != nil {
		return fmt.Errorf("获取购物车信息失败: %s", err)
	}

	if len(carts) == 0 {
		return fmt.Errorf("购物车为空")
	}

	var totalPrice float64
	for _, cart := range carts {
		for _, product := range cart.Product {
			totalPrice += product.Price * float64(cart.Quantity)
		}
	}

	order := models.Order{
		UserID:	user.ID,
		Status:	"待支付",
		TotalPrice:	totalPrice,
	}

	if err := dao.CreateOrder(&order); err != nil {
		return fmt.Errorf("创建订单失败: %s", err)
	}

	return nil
}