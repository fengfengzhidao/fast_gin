package models

// Pagination 分页查询的结构体
type Pagination struct {
	Limit int    `json:"limit" form:"limit" uri:"limit"` // 显示多少条
	Page  int    `json:"page" form:"page" uri:"page"`    // 第几页
	Sort  string `json:"sort" form:"sort" uri:"sort"`    // 排序  desc 降序  asc 升序 例如 id desc
	Key   string `json:"key" form:"key" uri:"key"`       // 模糊匹配字段
}
