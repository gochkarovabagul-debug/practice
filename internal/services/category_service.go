package services

import (
	"context"

	"github.com/gochkarovabagul-debug/practice/internal/models"
	"github.com/gochkarovabagul-debug/practice/internal/repositories"
)

func CategoryListService(c context.Context, filter models.CategoryFilter) (any, int, error) {
	list, total, err := repositories.CategoryList(c, filter)
	if err != nil {
		return nil, 0, err
	}
	res := []models.CategoryResponse{}
	for _, v := range list {
		item := models.CategoryResponse{}
		item.CategoryId = v.CategoryId
		item.Name = v.Name
		res = append(res, item)
	}
	return res, total, nil
}
func CreateCategoryService(c context.Context, name string) error {
	return repositories.CreateCategory(c, name)
}
func DeleteCategoryService(c context.Context, categoryid int) error {
	return repositories.DeleteCategory(c, categoryid)
}
func UpdateCategoryService(c context.Context, categoryid int, req models.CategoryCreateRequest) error {
	return repositories.UpdateCategory(c, categoryid, req)
}
func GetCategoryService(c context.Context, categoryid int) (models.CategoryResponse, error) {
	return repositories.GetCategory(c, categoryid)
}
