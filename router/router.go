package router

import (
	"WinterDemo/configs"
	"WinterDemo/middleware"
	"fmt"
	"WinterDemo/handler"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func InitRouter() *server.Hertz {
	h := server.Default(
		server.WithHostPorts(
			fmt.Sprintf("%s:%d",
				configs.GlobalConfig.Server.Host,
				configs.GlobalConfig.Server.Port,
			),
		),
	)

	// h.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
	// 	ctx.JSON(200, utils.H{"message": "pong"})
	// })

	//用户相关路由
	public := h.Group("/")
	{
		user := public.Group("/user")
		{
			user.POST("/register", handler.Register)
			user.POST("/login", handler.Login)
			user.POST("token/refresh", handler.RefreshToken)
			user.GET("info/:username", handler.GetUserInfo)
		}

		product := public.Group("/product")
		{
			product.GET("/list", handler.GetProductsByCategory)
			product.GET("/info/:id", handler.GetProductDetail)
			product.GET("/search", handler.SearchProduct)
			product.GET("/type", handler.GetProductsByCategory)
		}
	}

	// 需要鉴权的路由
	auth := h.Group("/")
	auth.Use(middleware.JWTauth())
	{
		//用户相关路由
		user := auth.Group("/user")
		{
			user.POST("password/update", handler.UpdatePassword)
			user.POST("info", handler.UpdateUserInfo)
		}

		// //商品相关路由
		// product := auth.Group("/product")
		// {
			
		// }
		// //评论相关路由
		// comment := auth.Group("/comment")
		// {

		// }

		// //购物车相关路由
		// cart := auth.Group("/cart")
		// {

		// }

		// //订单相关路由
		// order := auth.Group("/order")
		// {

		// }

	}
	return h
}
