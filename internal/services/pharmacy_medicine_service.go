package services

import (
	"context"
	"errors"

	"github.com/gochkarovabagul-debug/practice/internal/models"
	"github.com/gochkarovabagul-debug/practice/internal/repositories"
)

func PharmacyMedicineListService(c context.Context, filter models.PharmacyMedicineFilter) (any, error) {
	return repositories.PharmacyMedicineList(c, filter)
}
func CreatePharmacyMedicineService(c context.Context, name string, description string, price int, newprice int, categoryid int, pharmacyid int, token string) error {
	err := repositories.CreatePharmacyMedicine(c, name, description, price, newprice, categoryid, pharmacyid)
	if err != nil {
		return err
	}
	userid, err := repositories.GetUserIdByToken(c, token)
	if err != nil {
		return err
	}
	adminid, err := repositories.FindAdminId(c, pharmacyid)
	if err != nil {
		return err
	}
	if userid != adminid {
		return errors.New("forbidden")
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
