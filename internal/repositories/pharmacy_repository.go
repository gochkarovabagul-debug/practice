package repositories

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gochkarovabagul-debug/practice/internal/models"
	"github.com/gochkarovabagul-debug/practice/internal/utils"
)

func LenStrpharmacy(l []any) string {
	return strconv.Itoa(len(l))
}

func PharmacyList(c context.Context, f models.PharmacyFilter, moreArg ...int) ([]models.Pharmacy, error) {
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

func CreatePharmacy(c context.Context, name string, address string, hours int, adminuserid int) error {
	db := utils.GetDB()
	_, err := db.Exec(c, "insert into pharmacies (name, address, hours, admin_user_id) values ($1, $2, $3, $4)", name, address, hours, adminuserid)
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
	rows := db.QueryRow(c, "select  id, name, address, hours, admin_user_id from pharmacies where id=$1", id)
	err := rows.Scan(&req.Id, &req.Name, &req.Address, &req.Hours, &req.AdminUserId)
	if err != nil {
		return models.PharmacyResponse{}, err
	}
	return req, nil
}
func UpdatePharmacy(c context.Context, id int, req models.PharmacyCreateRequest) error {
	db := utils.GetDB()
	_, err := db.Exec(c, "update pharmacies set name=$1, address=$2, hours=$3, admin_user_id=$4 where id=$4", req.Name, req.Address, req.Hours, req.AdminUserId, id)
	if err != nil {
		return err
	}
	return nil
}
func FindNearbyPharmacies(c context.Context, lon float64, lat float64) (models.NearPharmacies, error) {
	db := utils.GetDB()
	fmt.Println("LON:", lon)
	fmt.Println("LAT:", lat)
	var result models.NearPharmacies

	rows := db.QueryRow(c, "select name from pharmacies order by ST_Distance(ST_MakePoint(longitude, latitude)::geography, ST_MakePoint($1, $2)::geography) ASC LIMIT 1", lon, lat)
	err := rows.Scan(&result.Name)
	if err != nil {
		return models.NearPharmacies{}, err
	}
	return result, nil
}
