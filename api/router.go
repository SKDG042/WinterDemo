package api

import (
	"context"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"WinterDemo/configs"
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
	

	h.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(200, utils.H{"message": "pong"})
	})	

	return h
}
