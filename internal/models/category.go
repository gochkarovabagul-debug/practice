package models

type CategoryResponse struct {
	CategoryId int    `json:"id"`
	Name       string `json:"name"`
}
type CategoryCreateRequest struct {
	Name string `json:"name"`
}
