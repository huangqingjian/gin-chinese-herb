package model

// 基础
type Base struct {
	CreateBy   int64		`json:"createBy" gorm:"column(create_by)"`						// 创建人
	CreateTime LocalTime	`json:"createTime" gorm:"autoCreateTime;column(create_time)"`   // 创建时间
	UpdateBy   int64		`json:"updateBy" gorm:"column(update_by)"`						// 更新人
	UpdateTime LocalTime	`json:"updateTime" gorm:"autoUpdateTime;column(update_time)"`	// 更新时间
	Deleted    int8			`json:"deleted" gorm:"column(deleted)"`							// 是否删除
}