package repositories

import (
	"context"

	"github.com/gochkarovabagul-debug/practice/internal/models"
	"github.com/gochkarovabagul-debug/practice/internal/utils"
)

func CreatePharmacyMedicine(c context.Context, name string, description string, price int, newprice int, categoryid int) error {
	db := utils.GetDB()
	_, err := db.Exec(c, "insert into pharmacymedicines  (name, description, price, new_price, category_id) values ($1, $2, $3, $4, $5)", name, description, price, newprice, categoryid)
	if err != nil {
		return err
	}
	return nil
}
func DeletePharmacyMedicine(c context.Context, id int) error {
	db := utils.GetDB()
	_, err := db.Exec(c, "delete from pharmacymedicines where id=$1", id)
	if err != nil {
		return err
	}
	return nil
}
func GetPharmacyMedicine(c context.Context, id int) (models.PharmacyMedicinesResponse, error) {
	db := utils.GetDB()
	var req models.PharmacyMedicinesResponse
	rows := db.QueryRow(context.Background(), "select  id, name, description, price, new_price, category_id from pharmacymedicines where id=$1", id)
	err := rows.Scan(&req.Id, &req.Name, &req.Description, &req.Price, &req.NewPrice, &req.CategoryId)
	if err != nil {
		return models.PharmacyMedicinesResponse{}, err
	}
	return req, nil
}
func UpdatePharmacyMedicine(c context.Context, id int, req models.PharmacyMedicinesCreateRequest) error {
	db := utils.GetDB()

	_, err := db.Exec(context.Background(), "update pharmacymedicines set name=$1, description=$2, price=$3, new_price=$4, category_id=$5  where id=$6", req.Name, req.Description, req.Price, req.NewPrice, req.CategoryId, id)
	if err != nil {
		return err
	}
	return nil
}
