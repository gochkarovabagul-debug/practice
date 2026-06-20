package models

type PharmacyMedicineFilter struct {
	Limit  int
	Offset int
	Search string
}

type PharmacyMedicine struct {
	Id          int
	Name        string
	Description string
	Price       int
	NewPrice    int
	CategoryId  int
	PharmacyId  int
	Stock       int
}

type PharmacyMedicinesResponse struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	NewPrice    int    `json:"new_price"`
	CategoryId  int    `json:"category_id"`
	PharmacyId  int    `json:"pharmacy_id"`
	Stock       int    `json:"stock"`
}
type PharmacyMedicinesCreateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	NewPrice    int    `json:"new_price"`
	CategoryId  int    `json:"category_id"`
	PharmacyId  int    `json:"pharmacy_id"`
	Stock       int    `json:"stock"`
}
