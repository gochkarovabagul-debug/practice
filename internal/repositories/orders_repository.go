package repositories

import (
	"context"

	"github.com/gochkarovabagul-debug/practice/internal/models"
	"github.com/gochkarovabagul-debug/practice/internal/utils"
)

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
