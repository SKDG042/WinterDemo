package dao

import (
	"WinterDemo/models"
)

// 获得所有商品
func GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	// Preload()方法用于预加载Categories这个关联表，否则查询时会出现Categories: nil
	result := DB.Preload("Categories").Find(&products)
	return products, result.Error
}

func SearchProducts(name string) ([]models.Product, error) {
	var products []models.Product
	// %name%是模糊搜索，搜索包含name的商品,然后通过Find()方法传给products
	result := DB.Preload("Categories").Where("name LIKE ?", "%"+name+"%").Find(&products)
	return products, result.Error
}

func GetProductsByCategory(categoryID int) ([]models.Product, error) {
	var products []models.Product
	// 隐式写法省略了DB.Model(&models.Product{}).,由GORM从Find()方法中推断出要查询的表
	result := DB.Preload("Categories"). 
		// Joins用于内连接，返回所连接两个表中都匹配的记录
		// 从 product_categories 表这个中间表中找到 category_id = 1 的记录
		// 通过 product_id 找到 products 表中对应的商品,输出到products
		Joins("JOIN product_categories ON products.id = product_categories.product_id").
		Where("product_categories.category_id = ?",categoryID).
		Find(&products)
	return products, result.Error	
}

func GetProductByID(id int) (models.Product, error) {
	var product models.Product
	result := DB.Preload("Categories").First(&product, id) // 根据id查询商品的第一个传给product
	return product, result.Error
}