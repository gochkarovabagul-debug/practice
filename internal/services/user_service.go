package services

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"

	"github.com/gochkarovabagul-debug/practice/internal/models"
	"github.com/gochkarovabagul-debug/practice/internal/repositories"
)

func UserListService(c context.Context, filter models.UserFilter) ([]models.UserResponse, int, error) {

	list, total, err := repositories.UserList(c, filter)
	if err != nil {
		return nil, 0, err
	}
	res := []models.UserResponse{}
	for _, v := range list {
		item := models.UserResponse{}
		item.Id = v.ID
		item.FirstName = v.FirstName
		item.LastName = v.LastName
		item.Role = v.Role
		item.Password = v.Password
		item.Email = v.Email
		res = append(res, item)
	}
	return res, total, nil
}
func RegistrationService(c context.Context, firstname string, lastname string, role string, password string, email string) error {
	return repositories.Registration(c, firstname, lastname, "customer", password, email)
}
func DeleteUserService(c context.Context, id int) error {
	return repositories.DeleteUser(c, id)
}
func GetUserService(c context.Context, token string, word bool) (models.UserResponse, error) {
	return repositories.GetUser(c, token, word)
}
func UpdateUserService(c context.Context, token string, req models.UserUpdateRequest) error {
	return repositories.UpdateUser(c, token, req)
}
func GenerateSecureToken(length int) string {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}
func LoginService(c context.Context, email string, password string) (string, error) {
	user, err := repositories.GetUserEmail(c, email)
	if err != nil {
		return "", err
	}
	if user.Password != password {
		return "", errors.New("wrong password")
	}
	Token := GenerateSecureToken(32)
	err = repositories.InsertToken(c, Token, user.Id)
	if err != nil {
		return "", err
	}
	return Token, nil
}
func LogoutService(c context.Context, token string) (string, error) {
	if repositories.CheckIsTokenReal(c, token) {
		err := repositories.DeleteToken(c, token)
		if err != nil {
			return "", err
		}
		return "logout", nil
	}
	return "not token", nil
}
func ChangePasswordService(c context.Context, token string, word bool, passchange models.ChangePasswordRequest, req models.UserResponse) error {
	req, err := repositories.GetUser(c, token, word)
	if err != nil {
		return err
	}
	if req.Password != passchange.OldPassword {
		return errors.New("wrong pass")
	}
	err = repositories.UpdatePassword(c, token, passchange)
	if err != nil {
		return err
	}
	return nil
}
