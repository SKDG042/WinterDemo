package handler

import (
	"WinterDemo/api/types"
    "WinterDemo/service"
    "context"
    "github.com/cloudwego/hertz/pkg/app"
    "github.com/cloudwego/hertz/pkg/protocol/consts"
	"strconv"
)

func AddComment(_ context.Context, ctx *app.RequestContext) {
	var req types.AddCommentRequest

	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(consts.StatusOK, types.ErrorResponse(10102, "请求参数有错误"))
		return
	}

	username := ctx.GetString("username")
	productID, err := strconv.Atoi(ctx.Param("product_id")) // 将string转换为int
	if err != nil {
		ctx.JSON(consts.StatusOK, types.ErrorResponse(10104, "商品ID有错误"))
		return
	}

	comment, err := service.AddComment(username, uint(productID), req.Content, req.ParentID)
	if err != nil {
		ctx.JSON(consts.StatusOK, types.ErrorResponse(10105, "添加评论失败"))
		return
	}

	data := map[string]interface{}{
		"comment_id": comment.CommentID,
		"message": "评论成功",
	}

	ctx.JSON(consts.StatusOK, types.SuccessResponse(data))
}

func DeleteComment(_ context.Context, ctx *app.RequestContext) {
	commentID, err := strconv.Atoi(ctx.Param("comment_id"))
	if err != nil {
		ctx.JSON(consts.StatusOK, types.ErrorResponse(10106, "评论ID有错误"))
		return
	}

	username := ctx.GetString("username")

	if err := service.DeleteComment(uint(commentID), username); err != nil {
		if err.Error() == "您没有权限删除该评论" {
			ctx.JSON(consts.StatusOK, types.ErrorResponse(10107, "您没有权限删除该评论"))
		}else{
			ctx.JSON(consts.StatusOK, types.ErrorResponse(10107, "删除评论失败"))
		}
		return
	}

	ctx.JSON(consts.StatusOK, types.SuccessResponse("删除评论成功"))
}

func GetCommentsByProductID(_ context.Context, ctx *app.RequestContext) {
	productID, err := strconv.Atoi(ctx.Param("product_id"))
	if err != nil {
		ctx.JSON(consts.StatusOK, types.ErrorResponse(10106, "商品ID有错误"))
		return
	}

	comments, err := service.GetCommentsByProductID(uint(productID))
	if err != nil {
		ctx.JSON(consts.StatusOK, types.ErrorResponse(10107, "获取评论失败"))
		return
	}

	ctx.JSON(consts.StatusOK, types.SuccessResponse(comments))
}

func UpdateComment(_ context.Context, ctx *app.RequestContext) {
	username := ctx.GetString("username")

	commentID, err := strconv.Atoi(ctx.Param("comment_id"))
	if err != nil {
		ctx.JSON(consts.StatusOK, types.ErrorResponse(10106, "评论ID有错误"))
		return
	}

	var req types.UpdateCommentRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(consts.StatusOK, types.ErrorResponse(10102, "请求参数有错误"))
		return
	}

	if err := service.UpdateComment(uint(commentID), &req.Content, username); err != nil {
		if err.Error() == "您没有权限修改该评论" {
			ctx.JSON(consts.StatusOK, types.ErrorResponse(10107, "您没有权限修改该评论"))
		}else{
			ctx.JSON(consts.StatusOK, types.ErrorResponse(10107, "更新评论失败"))
		}
		return
	}

	ctx.JSON(consts.StatusOK, types.SuccessResponse("更新评论成功"))
}	
