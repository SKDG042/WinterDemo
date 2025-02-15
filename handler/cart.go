package handler

import (
	"WinterDemo/api/types"
	"WinterDemo/service"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func AddCart(_ context.Context, ctx *app.RequestContext) {
	var req types.AddCartRequest

	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(consts.StatusOK, types.ErrorResponse(10102, "请求参数有错误"))
		return
	}

	username := ctx.GetString("username")

	if err := service.AddCart(username, req.ProductID, req.Quantity); err != nil {
		ctx.JSON(consts.StatusOK, types.ErrorResponse(10105, "添加购物车失败"))
		return
	}

	ctx.JSON(consts.StatusOK, types.SuccessResponse("添加购物车成功"))
}

func GetCartList(_ context.Context, ctx *app.RequestContext) {
	username := ctx.GetString("username")

	data , err := service.GetCartList(username)
	if err != nil {
		ctx.JSON(consts.StatusOK, types.ErrorResponse(10105, "获取购物车失败"))
		return
	}

	ctx.JSON(consts.StatusOK, types.SuccessResponse(data))
}