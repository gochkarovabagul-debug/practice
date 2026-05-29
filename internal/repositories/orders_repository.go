package repositories

import (
	"context"
	"strconv"

	"github.com/gochkarovabagul-debug/practice/internal/models"
	"github.com/gochkarovabagul-debug/practice/internal/utils"
)

type OrderFilter struct {
	Limit  int
	Offset int
	Search string
}

func LenStrorder(l []any) string {
	return strconv.Itoa(len(l))
}

func OrderList(c context.Context, f OrderFilter, moreArg ...int) ([]models.Order, error) {
	db := utils.GetDB()
	sqlWhere := ` `
	sqlArgs := []any{f.Limit, f.Offset}
	if f.Search != "" {
		sqlArgs = append(sqlArgs, f.Search)
		sqlWhere += `and (name ilike '%$` + LenStrorder(sqlArgs) + `%')`
	}

	rows, err := db.Query(c, `select id,name, price, description from orders where 1=1 `+sqlWhere+` limit $1 offset  $2`, sqlArgs...)
	if err != nil {
		return nil, err
	}
	list := []models.Order{}
	for rows.Next() {
		item := models.Order{}
		rows.Scan(&item.Id, &item.Name, &item.Price, &item.Description)
		list = append(list, item)
	}
	return list, nil
}

func CreateOrder(c context.Context, name string, price int, description string) error {
	db := utils.GetDB()
	_, err := db.Exec(c, "insert into orders  (name,  price, description) values ($1, $2, $3)", name, price, description)
	if err != nil {
		return err
	}
	return nil
}
func DeleteOrder(c context.Context, id int) error {
	db := utils.GetDB()
	_, err := db.Exec(c, "delete from orders where id=$1", id)
	if err != nil {
		return err
	}
	return nil
}
func GetOrder(c context.Context, id int) (models.OrderResponse, error) {
	db := utils.GetDB()
	var req models.OrderResponse
	rows := db.QueryRow(context.Background(), "select  id, name,  price,description from orders where id=$1", id)
	err := rows.Scan(&req.Id, &req.Name, &req.Price, &req.Description)
	if err != nil {
		return models.OrderResponse{}, err
	}
	return req, nil
}
func UpdateOrder(c context.Context, id int, req models.OrderCreateRequest) error {
	db := utils.GetDB()

	_, err := db.Exec(context.Background(), "update orders set name=$1,  price=$2, description=$3  where id=$4", req.Name, req.Price, req.Description, id)
	if err != nil {
		return err
	}
	return nil
}
