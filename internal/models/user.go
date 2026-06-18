package models

type UserFilter struct {
	Limit  int
	Offset int
	Search string
	Role   string
}
type User struct {
	ID        int
	FirstName string
	LastName  string
	Role      string
	Password  string
	Email     string
}
type UserCreateRequest struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Role      string `json:"role"`
	Password  string `json:"password" binding:"required,min=8"`
	Email     string `json:"email" binding:"required,email"`
}
type UserResponse struct {
	Id        int    `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Role      string `json:"role"`
	Password  string `json:"-"`
	Email     string `json:"email"`
}
type UserUpdateRequest struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}
type ChangePasswordRequest struct {
	OldPassword string `json:"oldpassword" binding:"required,min=8"`
	NewPassword string `json:"newpassword" binding:"required,min=8"`
}
