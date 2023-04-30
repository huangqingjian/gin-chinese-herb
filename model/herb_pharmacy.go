package model

import (
)

// 药方
type HerbPharmacy struct {
	Id          int64		`json:"id"`								// id
	HerbId      int64		`json:"herbId" orm:"column(herb_id)"`   // 中药Id
	Content     string		`json:"content"`						// 内容
	Base
}

func init() {

}

// 自定义表名
func (HerbPharmacy) TableName() string {
	return "tbl_herb_pharmacy"
}





