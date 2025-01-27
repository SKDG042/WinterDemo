package handler

import (
	"WinterDemo/api/types"
	"WinterDemo/service"
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

//在api/types/request.go中已经构建了结构体，接下来直接使用

// 注册函数
func Register(_ context.Context,ctx *app.RequestContext) {
	var req types.RegisterRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(consts.StatusOK, types.ErrorResponse(10001, "请求参数有错误"))
		return
	}

	//接下来要对入参进行合法性校验
	if len(req.Username) < 3 || len(req.Password) < 6 {
		ctx.JSON(consts.StatusOK, types.ErrorResponse(10002, "用户名长度至少3位，密码长度至少6位"))
		return
	}

	//入参合法，接下来要进行注册
	//service.Register已经检查了用户名是否存在，所以这里不需要再检查
	if err := service.Register(req.Username, req.Password); err != nil {
		ctx.JSON(consts.StatusOK, types.ErrorResponse(10004, "注册失败"))
		return
	}

	data := map[string]interface{}{
		"message": "注册成功",
	}
	ctx.JSON(consts.StatusOK, types.SuccessResponse(data))
}

// 登录函数
func Login(_ context.Context,ctx *app.RequestContext) {
	var req types.LoginRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(consts.StatusOK, types.ErrorResponse(10001, "请求参数有错误"))
		return
	}

	if token,err := service.Login(req.Username, req.Password); err != nil {
		ctx.JSON(consts.StatusOK, types.ErrorResponse(10005, err.Error()))
		return
	}else{
		data := map[string]interface{}{
			"token": token,
		}
		ctx.JSON(consts.StatusOK, types.SuccessResponse(data))
	}

}

func UpdateUser(_ context.Context,ctx *app.RequestContext) {
	var req types.UpdateUserRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(consts.StatusOK, types.ErrorResponse(10001, "请求参数有错误"))
		return
	}

	if req.NewUsername == "" && req.NewPassword == "" {
		ctx.JSON(consts.StatusOK, types.ErrorResponse(10007, "请求参数有错误"))
		return
	}

	//!=0即代表有修改
	if req.NewUsername != "" && len(req.NewUsername) < 3 {
		ctx.JSON(consts.StatusOK, types.ErrorResponse(10002, "用户名长度至少3位"))
		return
	}

	if req.NewPassword != "" && len(req.NewPassword) < 6 {
		ctx.JSON(consts.StatusOK, types.ErrorResponse(10003, "密码长度至少6位"))
		return
	}

	//从鉴权jwt的中间件ctx.set("username")中获取username
	username := ctx.GetString("username")

	if err := service.UpdateUser(username, req.NewUsername, req.NewPassword); err != nil {
		ctx.JSON(consts.StatusOK, types.ErrorResponse(10006, err.Error()))
		return
	}

	data := map[string]interface{}{
		"message": "更新成功",
	}
	ctx.JSON(consts.StatusOK, types.SuccessResponse(data))
}

