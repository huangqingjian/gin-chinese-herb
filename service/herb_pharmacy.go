package service

import (
	"gin-first/db"
	"gin-first/model"
	"gin-first/request"
)

// 插入药方
func AddHerbPharmacy(h *model.HerbPharmacy) int64 {
	err := db.DB.Create(h).Error
	if err != nil {
		panic(any(err))
	}
	return h.Id
}


// 更新药方
func UpdateHerbPharmacy(h *model.HerbPharmacy) int64 {
	db := db.DB.Model(model.HerbPharmacy{}).Where("id=?", h.Id).Updates(h)
	if err := db.Error; err != nil {
		panic(any(err))
	}
	count := db.RowsAffected
	return count
}

// 删除药方
func DeleteHerbPharmacy(id int64) int64 {
	db := db.DB.Model(model.HerbPharmacy{}).Where("id=?", id).Update("deleted", 1)
	if err := db.Error; err != nil {
		panic(any(err))
	}
	count := db.RowsAffected
	return count
}

// 查找药方
func GetHerbPharmacy(id int64) *model.HerbPharmacy {
	herbPharmacy := model.HerbPharmacy{Id : id}
	db := db.DB.Model(model.HerbPharmacy{}).Where("deleted = 0 and id=?", id)
	db = db.First(&herbPharmacy)
	if err := db.Error; err != nil {
		panic(any(err))
	}
	return &herbPharmacy
}

// 通过药物Id查询药方
func GetHerbPharmacyList(query *request.HerbPharmacyQuery) []model.HerbPharmacy {
	var herbPharmacys []model.HerbPharmacy
	db := db.DB.Model(model.HerbPharmacy{}).Where("deleted = 0")
	if query.HerbId != 0 {
		db = db.Where("herb_id=?", query.HerbId)
	}
	if query.HerbIds != nil && len(query.HerbIds) > 0 {
		db = db.Where("herb_id in ?", query.HerbIds)
	}
	db = db.Find(&herbPharmacys)
	if err := db.Error; err != nil {
		panic(any(err))
	}
	return herbPharmacys
}