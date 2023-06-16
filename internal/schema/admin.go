package schema

//request

type AdminRequest struct {
	AdminName string `json:"admin_name" validate:"required|string|min_len:4|max_len:10"`
	Password  string `json:"password" validate:"required|min_len:6|max_len:14"`
}

type AdminChangePasswordRequest struct {
	OldPassword     string `json:"old_password" validate:"required|min_len:6|max_len:14"`
	NewPassword     string `json:"new_password" validate:"required|min_len:6|max_len:14"`
	ConfirmPassword string `json:"confirm_password" validate:"required|min_len:6|max_len:14"`
}
