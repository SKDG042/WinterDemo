package dao

import (
	"WinterDemo/models"
	"fmt"
)

// 获得所有商品
func GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	// Preload()方法用于预加载Categories这个关联表，否则查询时会出现Categories: nil
	result := DB.Preload("Categories").Find(&products)
	if result.Error != nil {
		return []models.Product{}, fmt.Errorf("获取商品失败: %v", result.Error)
	}
	return products, nil
}

// 根据商品名称模糊搜索商品
func SearchProducts(name string) ([]models.Product, error) {
	var products []models.Product
	// %name%是模糊搜索，搜索包含name的商品,然后通过Find()
	// 使用BINARY关键字使中文字符能供工作
	result := DB.Preload("Categories").Where("name LIKE BINARY ?", "%"+name+"%").Find(&products)
	if result.Error != nil {
		return []models.Product{}, fmt.Errorf("获取商品失败: %v", result.Error)
	}
	return products, nil
}

// 根据分类ID获取商品
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
	if result.Error != nil {
		return []models.Product{}, fmt.Errorf("获取商品失败: %v", result.Error)
	}
	return products, nil
}

// 根据商品ID获取商品
func GetProductByID(id int) (models.Product, error) {
	var product models.Product
	result := DB.Preload("Categories").First(&product, id) // 根据id查询商品的第一个传给product
	if result.Error != nil {
		return models.Product{}, fmt.Errorf("获取商品失败: %v", result.Error)
	}
	return product, nil
}

// 添加商品分类
func AddCategory(category models.Category) error {
	result := DB.Create(&category)
	if result.Error != nil {
		return fmt.Errorf("添加分类失败: %v", result.Error)
	}
	return nil
}

// 根据名称获取分类
func GetCategoryByName(name string) (models.Category, error) {
	var category models.Category
	result := DB.Where("name = ?", name).First(&category)
	if result.Error != nil {
		return models.Category{}, fmt.Errorf("获取分类失败: %v", result.Error)
	}
	return category, nil
}

// 提娜佳商品,因为涉及创建和关联分类两个操作，所以需要使用事务以保证一致性
func AddProduct(product models.Product, categoryIDs []uint) error {
	// 开启事务
	tx := DB.Begin()

	// 如果发生任何panic，就会回滚所有事务
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	result := tx.Create(&product)
	if result.Error != nil {
		tx.Rollback()
		return fmt.Errorf("添加商品失败: %v", result.Error)
	}

	var categories []models.Category
	if err := tx.Where("id IN ?", categoryIDs).Find(&categories).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("获取分类失败: %v", err)
	}

	// 操作product的Categories与categories关联,即用Replace()方法将product的Categories替换为categories
	if err := tx.Model(&product).Association("Categories").Replace(categories); err != nil {
		tx.Rollback()
		return fmt.Errorf("关联分类失败: %v", err)
	}

	// 提交事务
	return tx.Commit().Error
}