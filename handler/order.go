package handler

import (
	"WinterDemo/api/types"
	"WinterDemo/service"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func CreateOrder(_ context.Context, ctx *app.RequestContext) {
	username := ctx.GetString("username")

	if err := service.CreateOrder(username); err != nil {
		if err.Error() == "购物车为空" {
			ctx.JSON(consts.StatusOK, types.ErrorResponse(10101, "请先添加商品到购物车"))
		} else {
			ctx.JSON(consts.StatusOK, types.ErrorResponse(10101, "创建订单失败"))
		}
		return
	}
	ctx.JSON(consts.StatusOK, types.SuccessResponse("创建订单成功"))
}