package models

type Order struct {
	Id          int
	Name        string
	Price       int
	Description string
}
type OrderResponse struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
}
type OrderCreateRequest struct {
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
}
type OrderFilter struct {
	Limit  int
	Offset int
	Search string
}
