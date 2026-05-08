package models

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
