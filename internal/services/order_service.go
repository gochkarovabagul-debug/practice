package services

import (
	"context"

	"github.com/gochkarovabagul-debug/practice/internal/models"
	"github.com/gochkarovabagul-debug/practice/internal/repositories"
)

func OrderListService(c context.Context, filter models.OrderFilter) (any, error) {
	return repositories.OrderList(c, filter)
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
