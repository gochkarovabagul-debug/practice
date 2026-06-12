package models

type CategoryFilter struct {
	Limit  int
	Offset int
	Search string
}
type Category struct {
	CategoryId int    `json:"id"`
	Name       string `json:"name"`
}
type CategoryResponse struct {
	CategoryId int    `json:"id"`
	Name       string `json:"name"`
}
type CategoryCreateRequest struct {
	Name string `json:"name" binding:"required"`
}
