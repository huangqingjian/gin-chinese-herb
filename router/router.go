package router

import (
	"gin-first/config"
	"gin-first/constant"
	"gin-first/controllers"
	"gin-first/response"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"net/http"
)

// 注册路由
func RegistRouter(e *gin.Engine) {
	router(e)
	e.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, response.Fail(http.StatusNotFound, "路径找不到～"))
	})
	e.Static(constant.PICTURE_PREFIX, config.Global.ImagePath)
}

// 路由
func router(e *gin.Engine) {
	apiRouter := e.Group("/api")
	{
		userRouter := apiRouter.Group("user")
		{
			user := &controllers.UserController{}
			userRouter.GET("list", user.GetUserList)
			userRouter.GET("get/:id", user.GetUser)
			userRouter.POST("add", user.AddUser)
		}
		bannerRouter := apiRouter.Group("banner")
		{
			banner := &controllers.BannerController{}
			bannerRouter.GET("list", banner.GetBannerList)
		}
		chineseHerbRouter := apiRouter.Group("chineseHerb")
		{
			chineseHerb := &controllers.ChineseHerbController{}
			chineseHerbRouter.GET("list", chineseHerb.GetChineseHerbList)
			chineseHerbRouter.GET("get/:id", chineseHerb.GetChineseHerb)
		}
		uploadRouter := apiRouter.Group("upload")
		{
			upload := &controllers.UploadController{}
			uploadRouter.GET("img", upload.UploadImg)
		}
	}

	// dev环境开启swagger
	if config.Global.RunMode == "dev" {
		e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}
