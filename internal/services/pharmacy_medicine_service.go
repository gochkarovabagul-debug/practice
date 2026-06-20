package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/gochkarovabagul-debug/practice/internal/models"
	"github.com/gochkarovabagul-debug/practice/internal/repositories"
)

func PharmacyMedicineListService(c context.Context, filter models.PharmacyMedicineFilter) (any, int, error) {
	list, total, err := repositories.PharmacyMedicineList(c, filter)
	if err != nil {
		return nil, 0, err
	}
	res := []models.PharmacyMedicinesResponse{}
	for _, v := range list {
		item := models.PharmacyMedicinesResponse{}
		item.Id = v.Id
		item.Name = v.Name
		item.Description = v.Description
		item.Price = v.Price
		item.NewPrice = v.NewPrice
		item.CategoryId = v.CategoryId
		item.PharmacyId = v.PharmacyId
		item.Stock = v.Stock
		res = append(res, item)
	}
	return res, total, nil
}
func CreatePharmacyMedicineService(c context.Context, name string, description string, price int, newprice int, categoryid int, pharmacyid int, stock int, token string) error {
	userid, err := repositories.GetUserIdByToken(c, token)
	if err != nil {
		return err
	}
	fmt.Println("userid: ", userid)
	adminid, err := repositories.FindAdminId(c, pharmacyid)
	if err != nil {
		return err
	}
	fmt.Println("adminid: ", adminid)
	if userid != adminid {
		return errors.New("forbidden")
	}
	err = repositories.CreatePharmacyMedicine(c, name, description, price, newprice, categoryid, pharmacyid, stock)
	if err != nil {
		return err
	}
	return nil
}
func DeletePharmacyMedicineService(c context.Context, id int) error {
	return repositories.DeletePharmacyMedicine(c, id)
}
func UpdatePharmacyMedicineService(c context.Context, id int, req models.PharmacyMedicinesCreateRequest) error {
	return repositories.UpdatePharmacyMedicine(c, id, req)
}
func GetPharmacyMedicineServices(c context.Context, id int) (models.OrderResponse, error) {
	return repositories.GetOrder(c, id)
}
