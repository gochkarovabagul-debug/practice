package models

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
