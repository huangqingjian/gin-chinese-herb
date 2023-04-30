package model

import (
	"github.com/go-playground/validator/v10"
)

// 中药
type ChineseHerb struct {
	Id          int64		`json:"id" form:"id" gorm:"column:id"`              // id
	Type        int64		`json:"type" binding:"required"`					// 类型
	Name        string		`json:"name" binding:"required"`					// 名称
	EnName      string		`json:"enName" form:"enName" gorm:"column:en_name"`	// 英文名
	Alias       string		`json:"alias"`										// 别名
	Pic         string		`json:"pic" binding:"required"`						// 图片
	Zwxt        string		`json:"zwxt"`										// 植物形态
	Yybw        string		`json:"yybw"`										// 药用部位
	Cdfb        string		`json:"cdfb"`										// 产地分布
	Csjg        string		`json:"csjg"`										// 采收加工
	Ycxz        string		`json:"ycxz"`										// 药材性状
	Xwgj        string		`json:"xwgj"`										// 性味归经
	Gxzy        string		`json:"gxzy"`										// 功效与作用
	Lcyy        string		`json:"lcyy"`										// 临床应用
	Ylyj        string		`json:"ylyj"`										// 药理研究
	Hxcf        string		`json:"hxcf"`										// 化学成分
	Syjj        string		`json:"syjj"`										// 使用禁忌
	HerbPharmacys []HerbPharmacy `json:"herbPharmacys" gorm:"-"`				// 药方
	Base
}

func init() {

}

// 自定义表名
func (ChineseHerb) TableName() string {
	return "tbl_chinese_herb"
}

// 绑定模型获取验证错误的方法
func (b *ChineseHerb) GetError (err validator.ValidationErrors) string {
	for _, f := range err {
		if f.StructField() == "Type" {
			switch f.Tag() {
			case "required":
				return "请选择中药类型"
			}
		} else if f.StructField() == "Name" {
			switch f.Tag() {
			case "required":
				return "请输入中药名称"
			}
		} else if f.StructField() == "Pic" {
			switch f.Tag() {
			case "required":
				return "请上传中药图片"
			}
		}
	}
	return "参数错误"
}






