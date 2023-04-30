package controllers

import (
	"gin-first/config"
	"gin-first/constant"
	"gin-first/response"
	"gin-first/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// 上传api
type UploadController struct {
}

// @Summary 图片上传
// @Description 图片上传
// @Param file formData	[]byte true "上传的图片"
// @Success 200 {object} response.Response
// @Router /upload/img [post]
func (u *UploadController) UploadImg(c *gin.Context) {
	fh, err := c.FormFile("file")
	if err != nil {
		panic(any(err))
	}
	fileName := fh.Filename
	fileName = util.GetUuid() + fileName[strings.LastIndex(fileName, constant.POINT):]
	c.SaveUploadedFile(fh, config.Global.App.ImagePath + constant.XG + fileName)
	c.JSON(http.StatusOK, response.Success(constant.PICTURE_PREFIX + constant.XG + fileName))
}