package models

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
	Password  string `json:"password"`
	Email     string `json:"email"`
}
type UserResponse struct {
	Id        int    `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Role      string `json:"role"`
	Password  string `json:"password"`
	Email     string `json:"email"`
}
