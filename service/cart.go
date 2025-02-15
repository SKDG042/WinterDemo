package service

import (
	"WinterDemo/api/types"
	"WinterDemo/dao"
	"WinterDemo/models"
	"fmt"
)

// 将购物车转换为响应格式
func convertToCartResponse(cart models.Cart) []types.CartProductResponse {
    var cartResponses []types.CartProductResponse
    for _, product := range cart.Product {
        cartResponses = append(cartResponses, types.CartProductResponse{
            ProductID: product.ID,
            Name:     product.Name,
            Type:     product.Type,
            Price:    product.Price,
            Cover:    product.Cover,
            Link:     product.Link,
            Num:      cart.Quantity,
        })
    }
    return cartResponses
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
		for _, product := range cart.Product {
			response.Account += float64(product.Price) * float64(cart.Quantity)
			// 使用语法糖将convertToCartResponse(cart)转换后的切片展开再追加到response.Cart中
			response.Cart = append(response.Cart, convertToCartResponse(cart)...)
		}
	}

    return &response, nil
}

// 清空购物车
func DeleteCart(username string) error {
	user, err := GetUserInfo(username)
	if err != nil {
		return fmt.Errorf("获取用户信息失败: %v", err)
	}

	if err := dao.DeleteCart(user.ID); err != nil {
		return fmt.Errorf("删除购物车失败: %v", err)
	}

	return nil
}
