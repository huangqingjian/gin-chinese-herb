package controllers

import (
	e "gin-first/error"
	"gin-first/model"
	"gin-first/request"
	"gin-first/response"
	"gin-first/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"
)

// 用户API
type UserController struct {
}

// @Summary 新增用户
// @Description 新增用户
// @Param body body model.User true "请求body"
// @Success 200 {object} response.Response
// @Router /user/add [get]
func (u *UserController) AddUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBind(&user); err != nil {
		panic(any(e.NewParamError(user.GetError(err.(validator.ValidationErrors)))))
	}
	id := service.AddUser(&user)
	c.JSON(http.StatusOK, response.Success(id))
}

// @Summary 查找用户
// @Description 通过id查找用户详情
// @param id path int true "用户id"
// @Success 200 {object} response.Response
// @Router /user/get/:id [get]
func (u *UserController) GetUser(c *gin.Context) {
	id := c.Param("id")
	userId, _ := strconv.ParseInt(id, 10, 64)
	user:= service.GetUser(userId)
	c.JSON(http.StatusOK, response.Success(user))
}

// @Summary 查询用户列表
// @Description 通过条件查询用户列表
// @param q query string true "用户名或手机号"
// @param pageSize query int false "分页大小"
// @param pageNo query int false "页码"
// @Success 200 {object} response.Response
// @Router /user/list [get]
func (u *UserController) GetUserList(c *gin.Context) {
	var query request.UserQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		panic(any(e.NewParamError("查询参数非法～")))
	}
	if query.PageNum < 1 {
		query.PageNum = 1
	}
	if query.PageSize <= 0 {
		query.PageSize = 10
	}
	page := service.GetUserList(&query)
	c.JSON(http.StatusOK, response.Success(page))
}


