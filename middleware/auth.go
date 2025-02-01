package middleware

import (
	"WinterDemo/api/types"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"strings"
	"WinterDemo/pkg/jwt"
)


func JWTauth() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		authHeader := string(ctx.GetHeader("Authorization"))
		if authHeader == "" {
			ctx.JSON(consts.StatusUnauthorized, types.ErrorResponse(10006,"请先登录"))
			ctx.Abort() // Abort()为中间件特有的方法，用于停止请求处理
			return
		}

		//JWT的token格式为Bearer token
		//通过检测authHeader是否以Bearer开头判断是否为Bearer token
		if !strings.HasPrefix(authHeader, "Bearer ") {
			ctx.JSON(consts.StatusUnauthorized, types.ErrorResponse(10006,"无效的token格式"))
			ctx.Abort()
			return 
		}

		token := authHeader[7:] //Bearer带空格一共7个字符

		username, err := jwt.ParseToken(token)
		if err != nil {
			ctx.JSON(consts.StatusUnauthorized, types.ErrorResponse(10006, "token解析错误"))
			ctx.Abort()
			return
		}

		ctx.Set("username", username) //将username设置到*app.RequestContext中
		ctx.Next(c) //调用下一个处理函数，这个请求周期没有结束
	}
}