package entity

type Page[T any] struct {
	// 当前页
	Page int `json:"page"`
	// 每页显示数量
	PageSize int `json:"pageSize"`
	// 总记录数
	Total int64 `json:"total"`
	// 数据
	List []T `json:"list"`
	// 排序
	Order string `json:"order"`
}
