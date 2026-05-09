package models

type Category struct {
	CategoryId int
	Name       string
}
type CategoryResponse struct {
	CategoryId int    `json:"id"`
	Name       string `json:"name"`
}
type CategoryCreateRequest struct {
	Name string `json:"name"`
}
