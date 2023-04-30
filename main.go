package main

import (
	"gin-first/config"
	"gin-first/constant"
	"gin-first/db"
	_ "gin-first/docs"
	"gin-first/logger"
	"gin-first/middleware"
	"gin-first/redis"
	"gin-first/router"
	"github.com/gin-gonic/gin"
)

// 初始化
func init() {
	config.InitConfig()
	logger.InitLogger()
	redis.InitRedis()
	db.InitDB()
}

// @title           中药 API
// @version         1.0.0
// @description     中药相关接口，包括用户、banner、中药、药方等
// @contact.email  2366850717@qq.com
// @host      localhost:80
// @BasePath  /api
func main() {
	//创建一个默认的路由引擎
	e := gin.New()
	e.Use(logger.GinLogger(), logger.GinRecovery(true))
	// 注册中间件
	middleware.RegistMiddleWare(e)
	// 添加路由
	router.RegistRouter(e)
	e.Run(constant.MH + config.Global.App.HttpPort)
}
