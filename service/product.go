package service

import (
	"WinterDemo/api/types"
	"WinterDemo/dao"
	"WinterDemo/models"
)

// 为了符合API文档的返回格式，我需要将Product转换为ProductResponse
func convertToProductResponse(p models.Product) types.ProductResponse {
	var categories []types.CategoryResponse
	for _,c := range p.Categories {
		categories = append(categories, types.CategoryResponse{
			ID:		c.ID,
			Name:	c.Name,
			Description:	c.Description,
		})
	}
	return types.ProductResponse{
		ProductID:   p.ID,
        Name:        p.Name,
        Description: p.Description,
        Type:        p.Type,
        CommentNum:  p.CommentNum,
        Price:       p.Price,
        IsAddCart:   p.IsAddCart,
        Cover:       p.Cover,
        PublishTime: p.PublishTime,
        Link:        p.Link,
        Categories:  categories,
	}
}

// 获取所有商品的列表
func GetProductList() (types.ProductListResponse, error) {
	products, err := dao.GetAllProducts()
	if err != nil {
		return	types.ProductListResponse{}, err
	}

	var response types.ProductListResponse
	// range遍历products，将每个Product(p)转换为ProductResponse，并添加到response.Products中
	for _, p := range products {
		response.Products = append(response.Products,convertToProductResponse(p))
	}
	return response, nil
}

// 搜索商品(模糊搜索获取)
func SearchProduct(name string) (types.ProductListResponse, error) {
	products, err := dao.SearchProducts(name)
	if err != nil {
		return types.ProductListResponse{}, err
	}

	var response types.ProductListResponse
	for _, p := range products {
		response.Products = append(response.Products, convertToProductResponse(p))
	}
	return response, nil
}

// 获取商品列表(根据分类ID)
func GetProductsByCategory(categoryID int) (types.ProductListResponse, error) {
	products, err := dao.GetProductsByCategory(categoryID)
	if err != nil {
		return types.ProductListResponse{}, err
	}
	
	var response types.ProductListResponse
	for _, p := range products {
		response.Products = append(response.Products, convertToProductResponse(p))
	}
	return response, nil
}


func GetProductDetail(id int) (types.ProductResponse, error) {
	product, err := dao.GetProductByID(id)
	if err != nil {
		return types.ProductResponse{}, err
	}
	
	return convertToProductResponse(product), nil
}
