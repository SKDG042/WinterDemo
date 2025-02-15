package router

import (
	"WinterDemo/configs"
	"WinterDemo/handler"
	"WinterDemo/middleware"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app/server"
)

func InitRouter() *server.Hertz {
	 h := server.Default(
        server.WithHostPorts(fmt.Sprintf("%s:%d", 
            configs.GlobalConfig.Server.Host,
            configs.GlobalConfig.Server.Port,
        )),
        server.WithNetwork("tcp4"), // 强制使用 IPv4
    )

	// h.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
	// 	ctx.JSON(200, utils.H{"message": "pong"})
	// })

	//用户相关路由
	public := h.Group("/")
	{
		user := public.Group("/user")
		{
			//用户注册
			user.POST("/register", handler.Register)
			//用户登录
			user.POST("/login", handler.Login)
			//刷新token
			user.POST("token/refresh", handler.RefreshToken)
			//获取用户信息
			user.GET("info/:username", handler.GetUserInfo)
		}

		product := public.Group("/product")
		{	
			//获取商品列表
			product.GET("/list", handler.GetProductList)
			//获取商品详情
			product.GET("/info/:id", handler.GetProductDetail)
			//搜索商品
			product.GET("/search", handler.SearchProduct)
			//按照分类id获取商品
			product.GET("/category/:id", handler.GetProductsByCategory)
			// 先放这，等之后有管理员账号再移动
			// 添加商品分类
			product.POST("/add/category", handler.AddCategory)
			// 添加商品
			product.POST("/add/product", handler.AddProduct)
		}
	}

	// 需要鉴权的路由
	auth := h.Group("/")
	auth.Use(middleware.JWTauth())
	{
		//用户相关路由
		user := auth.Group("/user")
		{
			// 仅修改密码
			user.POST("password/update", handler.UpdatePassword)
			// 修改用户信息
			user.POST("info", handler.UpdateUserInfo)
		}

		// //商品相关路由
		// product := auth.Group("/product")
		// {

		// }
		//评论相关路由
		comment := auth.Group("/comment")
		{	// 添加评论(支持匿名评论)
			comment.POST("/:product_id", handler.AddComment)
			// 删除评论
			comment.DELETE("/:comment_id", handler.DeleteComment)
			// 获取商品评论
			comment.GET("/:product_id", handler.GetCommentsByProductID)
			// 修改评论
			comment.PUT("/:comment_id", handler.UpdateComment)
		}

		//购物车相关路由
		cart := auth.Group("/cart")
		{	// 商品加入购物车
			cart.POST("/addCart", handler.AddCart)
			// 获取购物车列表
			cart.GET("/list", handler.GetCartList)
			// 清空购物车
			cart.DELETE("/clear", handler.DeleteCart)
		}

		//订单相关路由
		order := auth.Group("/operate")
		{
			// 创建订单
			order.POST("/order", handler.CreateOrder)
		}

	}
	return h
}
