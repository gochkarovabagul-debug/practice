package models

type PharmacyMedicine struct {
	Id          int
	Name        string
	Description string
	Price       int
	NewPrice    int
	CategoryId  int
}

type PharmacyMedicinesResponse struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	NewPrice    int    `json:"newprice"`
	CategoryId  int    `json:"categoryid"`
}
type PharmacyMedicinesCreateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	NewPrice    int    `json:"newprice"`
	CategoryId  int    `json:"categoryid"`
}
