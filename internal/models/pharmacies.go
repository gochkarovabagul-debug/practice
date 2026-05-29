package models

type Pharmacy struct {
	Id        int
	Name      string
	Address   string
	Hours     int
	Latitude  float64
	Longitude float64
}
type PharmacyResponse struct {
	Id        int     `json:"id"`
	Name      string  `json:"name"`
	Address   string  `json:"address"`
	Hours     int     `json:"hours"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
type PharmacyCreateRequest struct {
	Name      string  `json:"name"`
	Address   string  `json:"address"`
	Hours     int     `json:"hours"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
type NearPharmacies struct {
	Name string `json:"name"`
}
