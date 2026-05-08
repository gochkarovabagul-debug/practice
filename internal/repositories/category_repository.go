package repositories

import (
	"context"

	"github.com/gochkarovabagul-debug/practice/internal/models"
	"github.com/gochkarovabagul-debug/practice/internal/utils"
)

// func CategoryList()

func CreateCategory(c context.Context, name string) error {
	db := utils.GetDB()
	_, err := db.Exec(c, "insert into categories  (name) values ($1)", name)
	if err != nil {
		return err
	}
	return nil
}
func DeleteCategory(c context.Context, categoryid int) error {
	db := utils.GetDB()
	_, err := db.Exec(c, "delete from categories where categoryid=$1", categoryid)
	if err != nil {
		return err
	}
	return nil
}
func GetCategory(c context.Context, categoryid int) (models.CategoryResponse, error) {
	db := utils.GetDB()
	var req models.CategoryResponse
	rows := db.QueryRow(context.Background(), "select  categoryid, name from categories where categoryid=$1", categoryid)
	err := rows.Scan(&req.CategoryId, &req.Name)
	if err != nil {
		return models.CategoryResponse{}, err
	}
	return req, nil
}
func UpdateCategory(c context.Context, categoryid int, req models.CategoryCreateRequest) error {
	db := utils.GetDB()

	_, err := db.Exec(context.Background(), "update categories set name=$1 where categoryid=$2", req.Name, categoryid)
	if err != nil {
		return err
	}
	return nil
}
