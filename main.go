package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mindguard/dao"
	"mindguard/router"
	"mindguard/utils"
)

func main() {
	// 初始化配置
	cfg, err := utils.ParseConfig("./config/app.json")
	if err != nil {
		fmt.Println(err)
	}

	// 初始化数据库
	err = dao.InitMysql(cfg.Database)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 关闭数据库
	defer dao.Close()

	// 初始化redis
	utils.InitRedisStore(cfg.Redis)

	// 路由
	r := gin.Default()
	r.Use(router.Cors())
	router.RegisterRouter(r)
	fmt.Println("运行ok")

	err = r.Run(cfg.AppHost + ":" + cfg.AppPort)
	if err != nil {
		panic(err.Error())
	}

}
