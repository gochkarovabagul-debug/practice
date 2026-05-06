package repositories

import (
	"context"

	"github.com/gochkarovabagul/practice/internal/utils"
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
