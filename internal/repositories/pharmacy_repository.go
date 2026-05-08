package repositories

import (
	"context"

	"github.com/gochkarovabagul-debug/practice/internal/models"
	"github.com/gochkarovabagul-debug/practice/internal/utils"
)

func CreatePharmacy(c context.Context, name string, address string, hours int) error {
	db := utils.GetDB()
	_, err := db.Exec(c, "insert into pharmacies (name, address, hours) values ($1, $2, $3)", name, address, hours)
	if err != nil {
		return err
	}
	return nil
}
func DeletePharmacy(c context.Context, id int) error {
	db := utils.GetDB()
	_, err := db.Exec(c, "delete from pharmacies where id=$1", id)
	if err != nil {
		return err
	}
	return nil
}
func GetPharmacy(c context.Context, id int) (models.PharmacyResponse, error) {
	db := utils.GetDB()
	var req models.PharmacyResponse
	rows := db.QueryRow(context.Background(), "select  id, name, address, hours from pharmacies where id=$1", id)
	err := rows.Scan(&req.Id, &req.Name, &req.Address, &req.Hours)
	if err != nil {
		return models.PharmacyResponse{}, err
	}
	return req, nil
}
func UpdatePharmacy(c context.Context, id int, req models.PharmacyCreateRequest) error {
	db := utils.GetDB()

	_, err := db.Exec(context.Background(), "update pharmacies set name=$1, address=$2, hours=$3  where id=$4", req.Name, req.Address, req.Hours, id)
	if err != nil {
		return err
	}
	return nil
}
