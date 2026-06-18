package repositories

import (
	"context"
	"strconv"

	"github.com/gochkarovabagul-debug/practice/internal/models"
	"github.com/gochkarovabagul-debug/practice/internal/utils"
)

func LenStrpharmacymedicine(l []any) string {
	return strconv.Itoa(len(l))
}

func PharmacyMedicineList(c context.Context, f models.PharmacyMedicineFilter, moreArg ...int) ([]models.PharmacyMedicine, error) {
	db := utils.GetDB()
	if f.Limit == 0 {
		f.Limit = 10
	}
	sqlWhere := ` `
	sqlArgs := []any{f.Limit, f.Offset}
	if f.Search != "" {
		sqlArgs = append(sqlArgs, "%"+f.Search+"%")
		sqlWhere += `and (name ilike $` + LenStrpharmacymedicine(sqlArgs)
	}

	rows, err := db.Query(c, `select id,name, description, price, new_price, category_id, pharmacy_id from pharmacymedicines  where 1=1 `+sqlWhere+` limit $1 offset  $2`, sqlArgs...)
	if err != nil {
		return nil, err
	}
	list := []models.PharmacyMedicine{}
	for rows.Next() {
		item := models.PharmacyMedicine{}
		rows.Scan(&item.Id, &item.Name, &item.Description, &item.Price, &item.NewPrice, &item.CategoryId, &item.PharmacyId)
		list = append(list, item)
	}
	return list, nil
}

func CreatePharmacyMedicine(c context.Context, name string, description string, price int, newprice int, categoryid int, pharmacyid int) error {
	db := utils.GetDB()
	_, err := db.Exec(c,
		`insert into pharmacymedicines  (name, description, price, new_price, 
		category_id, pharmacy_id) values ($1, $2, $3, $4, $5, $6)`,
		name, description, price, newprice, categoryid, pharmacyid)
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
	rows := db.QueryRow(c, "select  id, name, description, price, new_price, category_id, pharmacy_id from pharmacymedicines where id=$1", id)
	err := rows.Scan(&req.Id, &req.Name, &req.Description, &req.Price, &req.NewPrice, &req.CategoryId, &req.PharmacyId)
	if err != nil {
		return models.PharmacyMedicinesResponse{}, err
	}
	return req, nil
}
func UpdatePharmacyMedicine(c context.Context, id int, req models.PharmacyMedicinesCreateRequest) error {
	db := utils.GetDB()

	_, err := db.Exec(c, "update pharmacymedicines set name=$1, description=$2, price=$3, new_price=$4, category_id=$5, pharmacy_id where id=$6", req.Name, req.Description, req.Price, req.NewPrice, req.CategoryId, req.PharmacyId, id)
	if err != nil {
		return err
	}
	return nil
}
func FindAdminId(c context.Context, id int) (int, error) {
	db := utils.GetDB()
	var adminid int
	rows := db.QueryRow(c, "select p.admin_user_id from pharmacymedicines pm join pharmacies p on pm.pharmacy_id=p.id where pm.pharmacy_id=$1", id)
	err := rows.Scan(&adminid)
	if err != nil {
		return 0, err
	}
	return adminid, nil
}
