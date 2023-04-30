package controller

import (
	e "gin-first/error"
	"gin-first/request"
	"gin-first/response"
	"gin-first/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 中药API
type ChineseHerbController struct {
}

// @Summary 查询中药列表
// @Description 通过条件查询中药列表
// @param q query string false "中药名"
// @param type query int64 false "中药类型"
// @param pageSize query int false "分页大小"
// @param pageNo query int false "页码"
// @Success 200 {object} response.Response
// @Router /chineseHerb/list [get]
func (*ChineseHerbController) GetChineseHerbList(c *gin.Context) {
	var query request.ChineseHerbQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		panic(any(e.NewParamError("查询参数非法～")))
	}
	if query.PageNum < 1 {
		query.PageNum = 1
	}
	if query.PageSize <= 0 {
		query.PageSize = 10
	}
	page := service.GetChineseHerbList(&query)
	c.JSON(http.StatusOK, response.Success(page))
}

// @Summary 查找中药
// @Description 通过id查找中药详情
// @param id path int true "中药id"
// @Success 200 {object} response.Response
// @Router /chineseHerb/get/:id [get]
func (*ChineseHerbController) GetChineseHerb(c *gin.Context) {
	id := c.Param("id")
	chineseHerbId, _ := strconv.ParseInt(id, 10, 64)
	// 查找中药详情
	chineseHerb := service.GetChineseHerb(chineseHerbId)
	c.JSON(http.StatusOK, response.Success(chineseHerb))
}
