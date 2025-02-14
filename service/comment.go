package service

import (
	"WinterDemo/dao"
	"WinterDemo/models"
	"fmt"
)

func AddComment(username string, productID uint, content string,ParentID *uint) (*models.Comment,error) {
	if len(content) == 0 {
		return nil, fmt.Errorf("评论内容不能为空")
	}
	if productID == 0 {
		return nil, fmt.Errorf("商品ID不能为空")
	}

	user, err := GetUserInfo(username)
	if err != nil {
		return nil, fmt.Errorf("获取用户信息失败: %s", err)
	}

	// 构建comment
	comment := models.Comment{
		Content:   content,
        UserID:    user.ID,
        Avatar:    user.Avatar,
        Nickname:  user.Nickname,
        ProductID: productID,
        ParentID:  0,
	}

	if  ParentID != nil {
		comment.ParentID = *ParentID
	}

	if err := dao.AddComment(&comment); err != nil {
		return nil, fmt.Errorf("添加评论失败: %s", err)
	}

	return &comment, nil
}

func DeleteComment(commentID uint, username string) error {

	comment, err := dao.GetComment(commentID)
	if err != nil {
		return fmt.Errorf("获取评论失败: %s", err)
	}

	user, err := GetUserInfo(username)
	if err != nil {
		return fmt.Errorf("获取用户信息失败: %s", err)
	}

	// 判断是否是用户本人的评论
	if user.ID != comment.UserID {
		return fmt.Errorf("您没有权限删除该评论")
	}

	if err := dao.DeleteComment(commentID); err != nil {
		return fmt.Errorf("删除评论失败: %s", err)
	}

	return nil
}

func GetCommentsByProductID(productID uint) ([]models.Comment, error) {
	comments, err := dao.GetCommentsByProductID(productID)
	if err != nil {
		return nil, fmt.Errorf("获取评论失败: %s", err)
	}
	return comments, nil
}

func UpdateComment(commentID uint, content *string, username string) error {
	comment, err := dao.GetComment(commentID)
	if err != nil {
		return fmt.Errorf("评论不存在: %s", err)
	}

	user, err := GetUserInfo(username)
	if err != nil {
		return fmt.Errorf("获取用户信息失败: %s", err)
	}

	if comment.UserID != user.ID {
		return fmt.Errorf("您没有权限修改该评论")
	}

	if err := dao.UpdateComment(commentID, content); err != nil {
		return fmt.Errorf("更新评论失败: %s", err)
	}
	return nil
}

