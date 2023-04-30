package service

import (
	"gin-first/db"
	"gin-first/model"
	"gin-first/request"
	"gin-first/response"
	"gorm.io/gorm"
)

// 新增用户
func AddUser(u *model.User) int64 {
	err := db.DB.Create(u).Error
	if err != nil {
		panic(any(err))
	}
	return u.Id
}

// 带事务式插入用户
func AddUserWithTrans(u *model.User) int64 {
	err := db.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(u).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		panic(any(err))
	}
	return u.Id
}

// 更新用户
func UpdateUser(u *model.User) int64{
	db := db.DB.Model(&model.User{}).Where("id=?", u.Id).Updates(u)
	if err := db.Error; err != nil {
		panic(any(err))
	}
	count := db.RowsAffected
	return count
}

// 带事务式更新用户
func UpdateUserWithTrans(u *model.User) int64 {
	var count int64
	err := db.DB.Transaction(func(tx *gorm.DB) error {
		db := tx.Model(&model.User{}).Where("id=?", u.Id).Updates(u)
		if err := db.Error; err != nil {
			return err
		}
		count = db.RowsAffected
		return nil
	})
	if err != nil {
		panic(any(nil))
	}
	return count
}

// 删除用户
func DeleteUser(id int64) int64 {
	db := db.DB.Model(&model.User{}).Where("id=?", id).Update("deleted", 1)
	if err := db.Error; err != nil {
		panic(any(err))
	}
	count := db.RowsAffected
	return count
}

// 带事务式删除用户
func DeleteUserWithTrans(u *model.User) int64 {
	var count int64
	err := db.DB.Transaction(func(tx *gorm.DB) error {
		db := tx.Model(&model.User{}).Where("id=?", u.Id).Update("deleted", 1)
		if err := db.Error; err != nil {
			return err
		}
		count = db.RowsAffected
		return nil
	})
	if err != nil {
		panic(any(err))
	}
	return count
}

// 查找用户
func GetUser(id int64) *model.User {
	user := model.User{Id : id}
	db := db.DB.Find(&user)
	if err := db.Error; err != nil {
		panic(any(err))
	}
	return &user
}

// 查询用户
func GetUserList(query *request.UserQuery) *response.Page {
	var users []model.User
	//db := component.DB.Table("tbl_user")
	db := db.DB.Model(&model.User{})
	db = db.Where("deleted = 0")
	if query.Q != "" {
		db = db.Where("name like ? or mobile like ?", query.Q +"%", query.Q + "%")
	}
	db = db.Limit(int(query.PageSize)).Offset(int((query.PageNum - 1) * query.PageSize))
	var count int64 = 0
	if err := db.Count(&count).Error; err != nil {
		panic(any(err))
	}
	if err := db.Find(&users).Error; err != nil {
		panic(any(err))
	}
	return response.NewPage(query.PageNum, query.PageSize, int32(len(users)), int32(count),  &users)
}
