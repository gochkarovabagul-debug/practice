package models

type Pharmacy struct {
	Id      int
	Name    string
	Address string
	Hours   int
}
type PharmacyResponse struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Hours   int    `json:"hours"`
}
type PharmacyCreateRequest struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Hours   int    `json:"hours"`
}
