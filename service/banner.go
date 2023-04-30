package service

import (
	"gin-first/db"
	"gin-first/model"
)

// 新增banner
func AddBanner(b *model.Banner) int64 {
	err := db.DB.Create(b).Error
	if err != nil {
		panic(any(err))
	}
	return b.Id
}

// 更新banner
func UpdateBanner(b *model.Banner) int64 {
	db := db.DB.Model(&model.Banner{}).Where("id=?", b.Id).Updates(b)
	if err := db.Error; err != nil {
		panic(any(err))
	}
	count := db.RowsAffected
	return count
}

// 删除banner
func DeleteBanner(id int64) int64 {
	db := db.DB.Model(&model.Banner{}).Where("id=?", id).Update("deleted", 1)
	if err := db.Error; err != nil {
		panic(any(err))
	}
	count := db.RowsAffected
	return count
}

// 查找banner
func GetBanner(id int64) *model.Banner {
	banner := model.Banner{Id : id}
	db := db.DB.Find(&banner)
	if err := db.Error; err != nil {
		panic(any(err))
	}
	return &banner
}

// 查询banner
func GetBannerList() []model.Banner {
	var banners []model.Banner
	db := db.DB.Model(&model.Banner{})
	db = db.Where("deleted = 0")
	if err := db.Find(&banners).Error; err != nil {
		panic(any(err))
	}
	return banners
}
