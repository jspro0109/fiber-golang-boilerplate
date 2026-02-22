package dto

type PaginationQuery struct {
	Page    int `query:"page"`
	PerPage int `query:"per_page"`
}
