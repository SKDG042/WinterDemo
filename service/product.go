package service

import (
	"WinterDemo/api/types"
	"WinterDemo/dao"
	"WinterDemo/models"
	"fmt"
)

// 为了符合API文档的返回格式，我需要将Product转换为ProductResponse
func convertToProductResponse(product models.Product) types.ProductResponse {
	//首先把product.Categories这个切片转换为types.CategoryResponse的格式
	var categories []types.CategoryResponse
	for _,c := range product.Categories {
		categories = append(categories, types.CategoryResponse{
			ID:		c.ID,
			Name:	c.Name,
			Description:	c.Description,
		})
	}
	return types.ProductResponse{
		ProductID:   product.ID,
        Name:        product.Name,
        Description: product.Description,
        Type:        product.Type,
        CommentNum:  product.CommentNum,
        Price:       product.Price,
        IsAddCart:   product.IsAddCart,
        Cover:       product.Cover,
        PublishTime: product.PublishTime,
        Link:        product.Link,
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

// 获取商品详情
func GetProductDetail(productID uint) (*types.ProductResponse, error) {
	product, err := dao.GetProductByID(productID)
	if err != nil {
		return nil, err
	}
	
	productResponse := convertToProductResponse(product)
	return &productResponse, nil
}

// 添加商品分类
func AddCategory(category types.AddCategoryRequest) error {

	if _,err := dao.GetCategoryByName(category.Name); err == nil {
		return fmt.Errorf("分类%s已存在", category.Name)
	}

	newCategory := models.Category{
		Name: category.Name,
		Description: category.Description,
	}
	
	err := dao.AddCategory(newCategory)
	if err != nil {
		return err
	}
	return nil
}

// 添加商品
func AddProduct(product types.AddProductRequest) error {
	newProduct := models.Product{
		Name: product.Name,
		Description: product.Description,
		Type: product.Type,
		Price: product.Price,
		Cover: product.Cover,
		PublishTime: product.PublishTime,
		Link: product.Link,
	}
	err := dao.AddProduct(newProduct, product.CategoryIDs)
	if err != nil {
		return err
	}
	return nil
}
