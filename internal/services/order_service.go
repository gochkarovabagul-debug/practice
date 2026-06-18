package services

import (
	"context"

	"github.com/gochkarovabagul-debug/practice/internal/models"
	"github.com/gochkarovabagul-debug/practice/internal/repositories"
)

func OrderListService(c context.Context, filter models.OrderFilter) (any, int, error) {
	list, total, err := repositories.OrderList(c, filter)
	if err != nil {
		return nil, 0, err
	}
	res := []models.OrderResponse{}
	for _, v := range list {
		item := models.OrderResponse{}
		item.Id = v.Id
		item.Name = v.Name
		item.Price = v.Price
		item.Description = v.Description
		res = append(res, item)
	}
	return res, total, nil
}
func CreateOrderService(c context.Context, name string, price int, description string) error {
	return repositories.CreateOrder(c, name, price, description)
}
func DeleteOrderService(c context.Context, id int) error {
	return repositories.DeleteCategory(c, id)
}
func UpdateOrderService(c context.Context, id int, req models.OrderCreateRequest) error {
	return repositories.UpdateOrder(c, id, req)
}
func GetOrderServices(c context.Context, id int) (models.OrderResponse, error) {
	return repositories.GetOrder(c, id)
}
