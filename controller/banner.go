package controller

import (
	"gin-first/response"
	"gin-first/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 横幅API
type BannerController struct {
}

// @Summary 查询banner列表
// @Description 查询全部banner
// @Success 200 {object} response.Response
// @Router /banner/list [get]
func (b *BannerController) GetBannerList(c *gin.Context) {
	banners := service.GetBannerList()
	c.JSON(http.StatusOK, response.Success(banners))
}
