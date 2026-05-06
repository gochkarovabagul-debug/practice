package repositories

import (
	"context"
	"strconv"

	"github.com/gochkarovabagul/practice/internal/models"
	"github.com/gochkarovabagul/practice/internal/utils"
)

type UserFilter struct {
	Limit  int
	Offset int
	Search string
	Role   string
}

func LenStr(l []any) string {
	return strconv.Itoa(len(l))
}

func UserList(c context.Context, f UserFilter, moreArg ...int) ([]models.User, error) {
	db := utils.GetDB()
	sqlWhere := ` `
	sqlArgs := []any{f.Limit, f.Offset}
	if f.Search != "" {
		sqlArgs = append(sqlArgs, f.Search)
		sqlWhere += `and (first_name ilike '%$` + LenStr(sqlArgs) + `%')`
	}
	if f.Role != "" {
		sqlArgs = append(sqlArgs, f.Search)
		sqlWhere += `and (role=$` + LenStr(sqlArgs) + `)`
	}
	rows, err := db.Query(c, `select id,first_name, role, password, email from users where 1=1 `+sqlWhere+` limit $1 offset  $2`, sqlArgs...)
	if err != nil {
		return nil, err
	}
	list := []models.User{}
	for rows.Next() {
		item := models.User{}
		rows.Scan(&item.ID, &item.FirstName, &item.LastName, &item.Role, &item.Password, &item.Email)
		list = append(list, item)
	}
	return list, nil
}
func CreateUser(c context.Context, firstname string, lastname string, role string, password string, email string) error {
	db := utils.GetDB()
	_, err := db.Exec(c, "insert into users (first_name, last_name, role, password, email) values ($1, $2, $3, $4, $5)", firstname, lastname, role, password, email)
	// _, err = db.Exec(context.Background(), "insert into expenses (Id, Date, Description, Amount, CategoryId )values ('"+idStr+"', '"+req.Date+"', '"+req.Description+"', '"+amountStr+"', '"+categoryIdstr+"');")
	if err != nil {
		return err
	}
	return nil
}
func DeleteUser(c context.Context, id int) error {
	db := utils.GetDB()
	_, err := db.Exec(c, "delete from users where id=$1", id)
	if err != nil {
		return err
	}
	return nil
}
func GetUser(c context.Context, id int) (models.UserResponse, error) {
	db := utils.GetDB()
	var req models.UserResponse
	rows := db.QueryRow(context.Background(), "select  id, first_name, last_name, role, password, email from users where id=$1", id)
	err := rows.Scan(&req.Id, &req.FirstName, &req.LastName, &req.Role, &req.Password, &req.Email)
	if err != nil {
		return models.UserResponse{}, err
	}
	return req, nil
}
func UpdateUser(c context.Context, id int, req models.UserCreateRequest) error {
	db := utils.GetDB()

	_, err := db.Exec(context.Background(), "update users set first_name=$1, last_name=$2, role=$3, password=$4, email=$5  where id=$6", req.FirstName, req.LastName, req.Role, req.Password, req.Email, id)
	if err != nil {
		return err
	}
	return nil
}

// func UserUpdate()
// func UserDelete()
