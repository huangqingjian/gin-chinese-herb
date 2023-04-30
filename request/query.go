package request

// 分页查询
type PageQuery struct {
	PageNum   int32		`form:"page"`// 当前页
	PageSize  int32		`form:"limit"`// 每页的数量
}

// 用户查询条件
type UserQuery struct {
	Q          string  `form:"q"`  // 查询条件
	PageQuery
}

// 中药查询查询
type ChineseHerbQuery struct {
	Q           string `form:"q"`    // 查询条件
	Type        int64  `form:"type"` // 类型
	PageQuery
}

// 药方查询
type HerbPharmacyQuery struct {
	HerbId      int64		// 中药Id
	HerbIds     []int64		// 中药Id（多个）
}
