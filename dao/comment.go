package dao

import (
	"WinterDemo/models"
	"fmt"
)

func AddComment(comment *models.Comment) error {
	// Create()方法需要一个指针
	if err := DB.Create(comment).Error; err != nil {
		return fmt.Errorf("添加评论失败: %s", err)
	}
	return nil
}

func GetComment(commentID uint) (*models.Comment, error) {
	var comment models.Comment
	if err := DB.First(&comment, commentID).Error; err != nil {
		return nil, fmt.Errorf("获取评论失败: %s", err)
	}
	return &comment, nil
}

func DeleteComment(commentID uint) error {
	if err := DB.Delete(&models.Comment{}, commentID).Error; err != nil {
		return fmt.Errorf("删除评论失败: %s", err)
	}
	return nil
}

func GetCommentsByProductID(productID uint) ([]models.Comment, error) {
	var comments []models.Comment

	if err := DB.Where("product_id = ?", productID).Find(&comments).Error; err != nil {
		return nil, fmt.Errorf("获取评论失败: %s", err)
	}
	return comments, nil
}

func UpdateComment(commentID uint, content *string) error {
	err := DB.Model(&models.Comment{}).Where("id = ?", commentID).Update("content", content).Error
	if err != nil {
		return fmt.Errorf("更新评论失败: %s", err)
	}
	return nil
}
