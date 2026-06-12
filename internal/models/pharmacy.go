package models

type PharmacyFilter struct {
	Limit  int
	Offset int
	Search string
}
type Pharmacy struct {
	Id          int
	Name        string
	Address     string
	Hours       int
	Latitude    float64
	Longitude   float64
	AdminUserId int
}
type PharmacyResponse struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Address     string  `json:"address"`
	Hours       int     `json:"hours"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	AdminUserId int     `json:"adminuserid"`
}
type PharmacyCreateRequest struct {
	Name        string  `json:"name"`
	Address     string  `json:"address"`
	Hours       int     `json:"hours"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	AdminUserId int     `json:"adminuserid"`
}
type NearPharmacies struct {
	Name string `json:"name"`
}
