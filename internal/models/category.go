package models

type CategoryRequest struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
type CategoryCreateRequest struct {
	Name string `json:"name"`
}
