package repositories

import (
	"context"
	"strconv"

	"github.com/gochkarovabagul-debug/practice/internal/models"
	"github.com/gochkarovabagul-debug/practice/internal/utils"
)

func LenStr(l []any) string {
	return strconv.Itoa(len(l))
}

func UserList(c context.Context, f models.UserFilter, moreArg ...int) ([]models.User, int, error) {
	db := utils.GetDB()
	if f.Limit == 0 {
		f.Limit = 10
	}
	sqlWhere := ` `
	sqlArgs := []any{f.Limit, f.Offset}
	if f.Search != "" {
		sqlArgs = append(sqlArgs, "%"+f.Search+"%")
		sqlWhere += `and name ilike $` + LenStr(sqlArgs)
	}
	if f.Role != "" {
		sqlArgs = append(sqlArgs, f.Role)
		sqlWhere += `and role=$` + LenStr(sqlArgs)
	}
	rows, err := db.Query(c, `select id, name, role, password, email, count(*) over() as total from users where 1=1 `+sqlWhere+` limit $1 offset  $2`, sqlArgs...)
	if err != nil {
		return nil, 0, err
	}
	var total int
	list := []models.User{}
	for rows.Next() {
		item := models.User{}
		rows.Scan(&item.ID, &item.Name, &item.Role, &item.Password, &item.Email, &total)
		list = append(list, item)
	}
	return list, total, nil
}
func Registration(c context.Context, name string, role string, password string, email string) error {
	db := utils.GetDB()
	_, err := db.Exec(c, "insert into users (name, role, password, email) values ($1, $2, $3, $4, $5)", name, role, password, email)
	if err != nil {
		return err
	}
	return nil
}
func CreateUserByAdmin(c context.Context, name string, role string, password string, email string) error {
	db := utils.GetDB()
	_, err := db.Exec(c, "insert into users (name, role, password, email) values ($1, $2, $3, $4, $5)", name, role, password, email)
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
func DeleteToken(c context.Context, token string) error {
	db := utils.GetDB()
	_, err := db.Exec(c, "delete from tokens where token=$1", token)
	if err != nil {
		return err
	}
	return nil
}

func GetUserByToken(c context.Context, token string, hasPass bool) (models.UserResponse, error) {
	db := utils.GetDB()
	var req models.UserResponse
	rows := db.QueryRow(c, "select  u.id, u.name, u.role, u.password, u.email from users u join tokens t on t.userid=u.id where t.token=$1", token)
	err := rows.Scan(&req.Id, &req.Name, &req.Role, &req.Password, &req.Email)
	if !hasPass {
		req.Password = ""
	}
	if err != nil {
		return models.UserResponse{}, err
	}
	return req, nil
}
func GetUserByEmail(c context.Context, email string, hasPass bool) (models.UserResponse, error) {
	db := utils.GetDB()
	var user models.UserResponse

	rows := db.QueryRow(c, "select  id, name, role, password, email from users where email=$1", email)

	err := rows.Scan(&user.Id, &user.Name, &user.Role, &user.Password, &user.Email)
	if !hasPass {
		user.Password = ""
	}
	if err != nil {
		return models.UserResponse{}, err
	}
	return user, nil
}
func GetUserById(c context.Context, id int, hasPass bool) (models.UserResponse, error) {
	db := utils.GetDB()
	var user models.UserResponse

	rows := db.QueryRow(c, "select  id, name, role, password, email from users where email=$1", id)

	err := rows.Scan(&user.Id, &user.Name, &user.Role, &user.Password, &user.Email)
	if !hasPass {
		user.Password = ""
	}
	if err != nil {
		return models.UserResponse{}, err
	}
	return user, nil
}

func UpdateUser(c context.Context, token string, req models.UserUpdateRequest) error {
	db := utils.GetDB()

	_, err := db.Exec(c, "update users u set name=$1 from tokens t where t.userid=u.id and t.token=$2", req.Name, token)
	if err != nil {
		return err
	}
	return nil
}

func UpdateUserById(c context.Context, id int, req models.UserUpdateRequest) error {
	db := utils.GetDB()

	_, err := db.Exec(c, "update users set name=$1, role=$2 where id=$3", req.Name, req.Role, id)
	if err != nil {
		return err
	}
	return nil
}
func UpdatePassword(c context.Context, token string, passchange models.ChangePasswordRequest) error {
	db := utils.GetDB()
	_, err := db.Exec(c, "update users u set password=$1 from tokens t where t.userid=u.id and t.token=$2", passchange.NewPassword, token)
	if err != nil {
		return err
	}
	return nil
}
