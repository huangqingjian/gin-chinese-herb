package service

import (
	"gin-first/db"
	"gin-first/model"
	"gin-first/request"
	"gin-first/response"
)

// 插入中药
func AddChineseHerb(c *model.ChineseHerb) int64{
	err := db.DB.Create(c).Error
	if err != nil {
		panic(any(err))
	}
	return c.Id
}

// 更新中药
func UpdateChineseHerb(c *model.ChineseHerb) int64 {
	db := db.DB.Model(model.ChineseHerb{}).Where("id=?", c.Id).Updates(c)
	if err := db.Error; err != nil {
		panic(any(err))
	}
	count := db.RowsAffected
	return count
}

// 删除中药
func DeleteChineseHerb(id int64) int64 {
	db := db.DB.Model(model.ChineseHerb{}).Where("id=?", id).Update("deleted", 1)
	if err := db.Error; err != nil {
		panic(any(err))
	}
	count := db.RowsAffected
	return count
}

// 查找中药
func GetChineseHerb(id int64) *model.ChineseHerb {
	chineseHerb := model.ChineseHerb{Id : id}
	db := db.DB.Model(model.ChineseHerb{}).Where("deleted = 0 and id=?", id)
	db = db.First(&chineseHerb)
	if err := db.Error; err != nil {
		panic(any(err))
	}
	// 查找药方
	herbPharmacys := GetHerbPharmacyList(&request.HerbPharmacyQuery{HerbId: id})
	chineseHerb.HerbPharmacys = herbPharmacys
	return &chineseHerb
}

// 查询中药
func GetChineseHerbList(query *request.ChineseHerbQuery) *response.Page {
	var chineseHerbs []model.ChineseHerb
	db := db.DB.Model(&model.ChineseHerb{})
	db = db.Where("deleted = 0")
	if query.Q != "" {
		db = db.Where("name like ?", query.Q +"%", query.Q + "%")
	}
	if query.Type != 0 {
		db = db.Where("type=?", query.Type)
	}
	db = db.Limit(int(query.PageSize)).Offset(int((query.PageNum - 1) * query.PageSize))
	var count int64 = 0
	if err := db.Count(&count).Error; err != nil {
		panic(any(err))
	}
	if err := db.Find(&chineseHerbs).Error; err != nil {
		panic(any(err))
	}
	return response.NewPage(query.PageNum, query.PageSize, int32(len(chineseHerbs)), int32(count),  &chineseHerbs)
}