package dto

type UpdateRoleRequest struct {
	Role string `json:"role" validate:"required,oneof=user admin"`
}

type AdminStatsResponse struct {
	ActiveUsers   int64 `json:"active_users"`
	DeletedUsers  int64 `json:"deleted_users"`
	TotalFiles    int64 `json:"total_files"`
	TotalFileSize int64 `json:"total_file_size"`
}

type AdminUserQuery struct {
	PaginationQuery
}
