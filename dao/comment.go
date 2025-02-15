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

// 获取商品的评论
// 增加嵌套评论的思路 : 先获取顶级评论， 然后通过递归获取每个顶级评论的子评论
func GetCommentsByProductID(productID uint) ([]models.Comment, error) {
	var comments []models.Comment

	// 获取顶级评论
	if err := DB.Where("product_id = ? AND parent_id = 0", productID).Find(&comments).Error; err != nil {
		return nil, fmt.Errorf("获取评论失败: %s", err)
	}

	// 获取每个顶级评论的子评论
	for i := range comments {
		if err := getChildrenComments(&comments[i]); err != nil {
			return nil, fmt.Errorf("获取%d评论失败: %s", comments[i].ID, err)
		}
	}
	
	return comments, nil
}

// 递归获得子评论
func getChildrenComments(comment *models.Comment) error {
	var children []models.Comment

	// 寻找prent_id为comment.ID的评论
	if err := DB.Where("parent_id = ?", comment.ID).Find(&children).Error; err != nil {
		return fmt.Errorf("获取子评论失败: %s", err)
	}

	// 然后递归继续获得子子评论，子子子评论...
	for i := range children {
		if err := getChildrenComments(&children[i]); err != nil {
			return fmt.Errorf("获取%d评论失败: %s", children[i].ID, err)
		}
	}

	comment.Children = children

	return nil
}


func UpdateComment(commentID uint, content *string) error {
	err := DB.Model(&models.Comment{}).Where("id = ?", commentID).Update("content", content).Error
	if err != nil {
		return fmt.Errorf("更新评论失败: %s", err)
	}
	return nil
}
