package repositories

import (
	"context"
	"time"

	"github.com/gochkarovabagul-debug/practice/internal/models"
	"github.com/gochkarovabagul-debug/practice/internal/utils"
)

func InsertToken(c context.Context, Token string, userid int) error {
	db := utils.GetDB()
	_, err := db.Exec(c, "insert into tokens (token, userid) values ($1, $2)", Token, userid)
	if err != nil {
		return err
	}
	return nil
}
func GetUserIdByToken(c context.Context, token string) (int, error) {
	db := utils.GetDB()
	var req models.TokenResponse
	rows := db.QueryRow(c, "select token, userid from tokens where token=$1", token)
	err := rows.Scan(&req.Token, &req.UserId)
	if err != nil {
		return 0, err
	}
	return req.UserId, nil
}
func GetRoleByToken(c context.Context, token string) (string, error) {
	db := utils.GetDB()
	var role string
	rows := db.QueryRow(c, "select u.role from users u join tokens t on t.userid=u.id where t.token=$1", token)
	err := rows.Scan(&role)
	if err != nil {
		return "", err
	}
	return role, nil
}

func CheckIsTokenReal(c context.Context, token string) bool {
	db := utils.GetDB()
	var Token models.TokenCheck
	rows, err := db.Query(c, "select token from tokens where token=$1", token)
	if err != nil {
		return false
	}
	err = rows.Scan(&Token.Token)
	if err != nil {
		return true
	}
	return false
}
func GetExpiresAtByToken(c context.Context, token string) (time.Time, error) {
	db := utils.GetDB()
	var ExpiresAt time.Time
	rows := db.QueryRow(c, "select expires_at from tokens where token=$1", token)
	err := rows.Scan(&ExpiresAt)
	if err != nil {
		return time.Time{}, err
	}
	return ExpiresAt, nil
}
