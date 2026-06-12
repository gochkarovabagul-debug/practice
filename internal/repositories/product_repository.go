package repositories

import (
	"context"
	"strconv"

	"github.com/gochkarovabagul-debug/practice/internal/models"
	"github.com/gochkarovabagul-debug/practice/internal/utils"
)

func LenStrProduct(l []any) string {
	return strconv.Itoa(len(l))
}

func ProductList(c context.Context, f models.ProductFilter, moreArg ...int) ([]models.Product, error) {
	db := utils.GetDB()
	sqlWhere := ` `
	sqlArgs := []any{f.Limit, f.Offset}
	if f.Search != " " {
		sqlArgs = append(sqlArgs, f.Search)
		sqlWhere += `and (name ilike '%$` + LenStrProduct(sqlArgs) + `%')`
	}
	rows, err := db.Query(c, `select id, name, pharmacy_id from products where 1=1 `+sqlWhere+` limit $1 offset $2`, sqlArgs...)
	if err != nil {
		return nil, err
	}
	list := []models.Product{}
	for rows.Next() {
		item := models.Product{}
		rows.Scan(&item.Id, &item.Name, &item.PharmacyId)
		list = append(list, item)
	}
	return list, nil
}

func CreateProduct(c context.Context, name string, pharmacyid int) error {
	db := utils.GetDB()
	_, err := db.Exec(c, "insert into products (name, pharmacy_id) values ($1, $2)", name, pharmacyid)
	if err != nil {
		return err
	}
	return nil
}
func DeleteProduct(c context.Context, id int) error {
	db := utils.GetDB()
	_, err := db.Exec(c, "delete from products where id=$1", id)
	if err != nil {
		return err
	}
	return nil
}
func UpdateProduct(c context.Context, id int, product models.ProductCreateRequest) error {
	db := utils.GetDB()
	_, err := db.Exec(c, "update products set name=$1, pharmacy_id=$2 where id=$3", product.Name, product.PharmacyId, id)
	if err != nil {
		return err
	}
	return nil
}
func GetProduct(c context.Context, id int) (models.ProductResponse, error) {
	db := utils.GetDB()
	var product models.ProductResponse
	rows := db.QueryRow(c, "select name, pharmacy_id from products where id=$1", id)
	err := rows.Scan(&product.Name, &product.PharmacyId)
	if err != nil {
		return models.ProductResponse{}, err
	}
	return product, nil
}
