package models

type Product struct {
	Id         int
	Name       string
	PharmacyId int
}

type ProductResponse struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	PharmacyId int    `json:"pharmacyid"`
}
type ProductCreateRequest struct {
	Name       string `json:"name"`
	PharmacyId int    `json:"pharmacyid"`
}
type ProductFilter struct {
	Limit  int
	Offset int
	Search string
}
