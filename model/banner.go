package model

import (
	"github.com/go-playground/validator/v10"
)

// banner
type Banner struct {
	Id          int64		`json:"id" form:"id" gorm:"column:id"`								// id
	Title       string		`json:"title" form:"title" binding:"required" gorm:"column:title"`  // 标题
	Url         string		`json:"url" form:"url" binding:"required" gorm:"column:url"`		// 图片url
	Link        string		`json:"link" form:"link" gorm:"column:link"`						// 跳转链接
	Background  string		`json:"background" form:"background" gorm:"column:background"`		// 背景
	Sort        int32		`json:"sort" form:"sort" gorm:"column:sort"`						// 排序
	Desc        string		`json:"desc" form:"desc" gorm:"column:desc"`						// 描述
	Base
}

func init() {
}

// 自定义表名
func (Banner) TableName() string {
	return "tbl_banner"
}

// 绑定模型获取验证错误的方法
func (b *Banner) GetError (err validator.ValidationErrors) string {
	for _, f := range err {
		if f.StructField() == "Title" {
			switch f.Tag() {
			case "required":
				return "请输入banner标题"
			}
		} else if f.StructField() == "Url" {
			switch f.Tag() {
			case "required":
				return "请输入banner url"
			}
		}
	}
	return "参数错误"
}





