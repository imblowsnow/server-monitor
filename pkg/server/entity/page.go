package entity

type Page[T any] struct {
	// 当前页
	Page int `json:"page"`
	// 每页显示数量
	PageSize int `json:"pageSize"`
	// 总页数
	TotalPage int `json:"totalPage"`
	// 总记录数
	TotalCount int64 `json:"totalCount"`
	// 数据
	List []T `json:"list"`
}
