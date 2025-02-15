package service

import (
	"fmt"
	"WinterDemo/dao"
	"WinterDemo/models"
	"WinterDemo/api/types"
)

func convertToCartResponse(cart models.Cart) types.CartResponse {
	CartResponse := types.CartResponse{
		ProductID: cart.ProductID,
		Name: cart.Product.Name,
		Type: cart.Product.Type,
		Price: cart.Product.Price,
		Cover: cart.Product.Cover,
		Link: cart.Product.Link,
		Num: cart.Quantity,
	}
	return CartResponse
}

func AddCart(username string, productID uint, quantity int) error {
	user, err := GetUserInfo(username)
	if err != nil {
		return fmt.Errorf("获取用户信息失败: %v", err)
	}

	_ , err = GetProductDetail(productID)
	if err != nil {
		return fmt.Errorf("商品不存在: %v", err)
	}

	if quantity <= 0 {
		return fmt.Errorf("购买数量必须大于0")
	}

	if err := dao.AddCart(user.ID, productID, quantity); err != nil {
		return fmt.Errorf("添加购物车失败: %v", err)
	}

	return nil
}

func GetCartList(username string) (*types.CartListResponse, error) {
	user, err := GetUserInfo(username)
	if err != nil {
		return nil, fmt.Errorf("获取用户信息失败: %v", err)
	}

	carts, err := dao.GetCartByUserID(user.ID)
	if err != nil {
		return nil, fmt.Errorf("获取购物车失败: %v", err)
	}

	var response types.CartListResponse

	for _, cart := range carts {
		response.Cart = append(response.Cart, convertToCartResponse(cart))
		response.Account += float64(cart.Quantity) * cart.Product.Price
	}

	return &response, nil
}
