package repositories

import (
	"context"
	"strconv"

	"github.com/gochkarovabagul-debug/practice/internal/models"
	"github.com/gochkarovabagul-debug/practice/internal/utils"
)

func LenStrcategory(l []any) string {
	return strconv.Itoa(len(l))
}

func CategoryList(c context.Context, f models.CategoryFilter, moreArg ...int) ([]models.Category, error) {
	db := utils.GetDB()
	if f.Limit == 0 {
		f.Limit = 10
	}
	sqlWhere := ` `
	sqlArgs := []any{f.Limit, f.Offset}
	if f.Search != "" {
		sqlArgs = append(sqlArgs, f.Search)
		sqlWhere += `and (name ilike '%$` + LenStrcategory(sqlArgs) + `%')`
	}
	rows, err := db.Query(c, `select categoryid, name from categories where 1=1 `+sqlWhere+` limit $1 offset  $2`, sqlArgs...)
	if err != nil {
		return nil, err
	}
	list := []models.Category{}
	for rows.Next() {
		item := models.Category{}
		rows.Scan(&item.CategoryId, &item.Name)
		list = append(list, item)
	}
	return list, nil
}

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
	_, err := db.Exec(c, "delete from categories where id=$1", categoryid)
	if err != nil {
		return err
	}
	return nil
}
func GetCategory(c context.Context, categoryid int) (models.CategoryResponse, error) {
	db := utils.GetDB()
	var req models.CategoryResponse
	rows := db.QueryRow(c, "select  id, name from categories where categoryid=$1", categoryid)
	err := rows.Scan(&req.CategoryId, &req.Name)
	if err != nil {
		return models.CategoryResponse{}, err
	}
	return req, nil
}
func UpdateCategory(c context.Context, categoryid int, req models.CategoryCreateRequest) error {
	db := utils.GetDB()

	_, err := db.Exec(c, "update categories set name=$1 where id=$2", req.Name, categoryid)
	if err != nil {
		return err
	}
	return nil
}
