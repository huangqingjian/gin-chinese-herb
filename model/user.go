package model

import (
	"github.com/go-playground/validator/v10"
)

// 用户
type User struct {
	Id         int64		`json:"id" form:"id" gorm:"column:id"`									// id
	Name       string		`json:"name" form:"name" binding:"required" gorm:"column:name"`			// 姓名
	Mobile     string		`json:"mobile" form:"mobile" binding:"required" gorm:"column:mobile"`	// 手机号
	Email      string		`json:"email" form:"email" binding:"required,email" gorm:"column:email"`// 邮箱
	Password   string		`json:"password" form:"password" gorm:"column:password"`				// 密码
	Face       string		`json:"face" form:"face" gorm:"column:face"`							// 头像
	Sex        int8			`json:"sex" form:"sex" binding:"gte=1,lte=2" gorm:"column:sex"`			// 性别
	Desc       string		`json:"desc" form:"desc" gorm:"column:desc"`							// 描述
	Base
}

func init() {
}

// 自定义表名
func (User) TableName() string {
	return "tbl_user"
}

// 绑定模型获取验证错误的方法
func (r *User) GetError (err validator.ValidationErrors) string {
	for _, f := range err {
		if f.StructField() == "Name" {
			switch f.Tag() {
			case "required":
				return "请输入用户名"
			}
		} else if f.StructField() == "Mobile" {
			switch f.Tag() {
			case "required":
				return "请输入手机号"
			}
		} else if f.StructField() == "Email" {
			switch f.Tag() {
			case "required":
				return "请输入邮箱"
			case "email":
				return "输入的邮箱格式非法"
			}
		} else if f.StructField() == "Sex" {
			switch f.Tag() {
			case "gte":
				return "请选择正确的性别"
			case "lte":
				return "请选择正确的性别"
			}
		}
	}
	return "参数错误"
}





