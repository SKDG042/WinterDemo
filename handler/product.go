package handler

import (
    "WinterDemo/api/types"
    "WinterDemo/service"
    "context"
    "github.com/cloudwego/hertz/pkg/app"
    "github.com/cloudwego/hertz/pkg/protocol/consts"
	"strconv"
)

// 获取商品列表
func GetProductList(_ context.Context, ctx *app.RequestContext) {
	response, err := service.GetProductList()
	if err != nil {
		ctx.JSON(consts.StatusOK, types.ErrorResponse(10001, "获取商品列表失败"))
		return
	}
	ctx.JSON(consts.StatusOK, types.SuccessResponse(response))
}

// 搜索商品
func SearchProduct(_ context.Context, ctx *app.RequestContext) {
	var req types.SearchProductRequest
	if err := ctx.BindQuery(&req); err != nil {
		ctx.JSON(consts.StatusOK, types.ErrorResponse(10002, "请求参数有错误"))
		return
	}

	response, err := service.SearchProduct(req.ProductName)
	if err != nil {
		ctx.JSON(consts.StatusOK, types.ErrorResponse(10003, "搜索商品失败"))
		return
	}
	ctx.JSON(consts.StatusOK, types.SuccessResponse(response))
}

// 根据分类ID获取商品列表
func GetProductsByCategory(_ context.Context, ctx *app.RequestContext) {
	var req types.GetProductsByCategoryRequest
	if err := ctx.BindQuery(&req); err != nil {
		ctx.JSON(consts.StatusOK, types.ErrorResponse(10004, "请求参数有错误"))
		return
	}

	response, err := service.GetProductsByCategory(req.CategoryID)
	if err != nil {
		ctx.JSON(consts.StatusOK, types.ErrorResponse(10005, "获取商品列表失败"))
		return
	}
	ctx.JSON(consts.StatusOK, types.SuccessResponse(response))
}

// 获取商品详情
func GetProductDetail(_ context.Context, ctx *app.RequestContext) {
	id := ctx.Query("id")
	if id == "" {
		ctx.JSON(consts.StatusOK, types.ErrorResponse(10007, "商品ID不能为空"))
		return
	}

	// Query获取的id是string类型，需要转换为uint类型
	productID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(consts.StatusOK, types.ErrorResponse(10008, "商品ID格式错误"))
		return
	}

	response, err := service.GetProductDetail(productID)
	if err != nil {
		ctx.JSON(consts.StatusOK, types.ErrorResponse(10009, "获取商品详情失败"))
		return
	}

	ctx.JSON(consts.StatusOK, types.SuccessResponse(response))
}

