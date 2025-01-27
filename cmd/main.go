package main

import (
	"WinterDemo/configs"
	"WinterDemo/dao"
	"WinterDemo/router"
	"fmt"
	"log"
)

func main() {
	if err := configs.InitConfig(); err != nil {
		log.Fatalf("初始化配置失败: %s", err)
	}

	if err := dao.InitDB(); err != nil {
		log.Fatalf("初始化数据库失败: %s", err)
	}

	defer func() {
		if err := dao.CloseDB(); err != nil {
			log.Fatalf("关闭数据库失败: %s", err)
		}
	}()

	h := router.InitRouter()
	h.Spin()
	fmt.Println("服务器成功关闭")
}
