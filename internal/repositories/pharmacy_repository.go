package repositories

import (
	"context"
	"strconv"

	"github.com/gochkarovabagul-debug/practice/internal/models"
	"github.com/gochkarovabagul-debug/practice/internal/utils"
)

type PharmacyFilter struct {
	Limit  int
	Offset int
	Search string
}

func LenStrpharmacy(l []any) string {
	return strconv.Itoa(len(l))
}

func PharmacyList(c context.Context, f PharmacyFilter, moreArg ...int) ([]models.Pharmacy, error) {
	db := utils.GetDB()
	sqlWhere := ` `
	sqlArgs := []any{f.Limit, f.Offset}
	if f.Search != "" {
		sqlArgs = append(sqlArgs, f.Search)
		sqlWhere += `and (name ilike '%$` + LenStrpharmacy(sqlArgs) + `%')`
	}
	rows, err := db.Query(c, `select id,name, address, hours from pharmacies where 1=1 `+sqlWhere+` limit $1 offset  $2`, sqlArgs...)
	if err != nil {
		return nil, err
	}
	list := []models.Pharmacy{}
	for rows.Next() {
		item := models.Pharmacy{}
		rows.Scan(&item.Id, &item.Name, &item.Address, &item.Hours)
		list = append(list, item)
	}
	return list, nil
}

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
