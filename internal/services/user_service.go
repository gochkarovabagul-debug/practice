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
		item.Name = v.Name
		item.Role = v.Role
		item.Password = v.Password
		item.Email = v.Email
		res = append(res, item)
	}
	return res, total, nil
}
func RegistrationService(c context.Context, name string, role string, password string, email string) (string, error) {
	err := repositories.Registration(c, name, "customer", password, email)
	if err != nil {
		return "", err
	}
	user, err := repositories.GetUserByEmail(c, email, true)
	if err != nil {
		return "", err
	}
	token, err := NewToken(c, 32, user.Id)
	if err != nil {
		return "", err
	}
	return token, nil
}
func CreateUserByAdminService(c context.Context, name string, role string, password string, email string) (models.UserResponse, error) {
	err := repositories.CreateUserByAdmin(c, name, role, password, email)
	if err != nil {
		return models.UserResponse{}, err
	}
	user, err := repositories.GetUserByEmail(c, email, false)
	return user, nil
}
func DeleteUserService(c context.Context, id int) error {
	return repositories.DeleteUser(c, id)
}
func GetUserByTokenService(c context.Context, token string, word bool) (models.UserResponse, error) {
	return repositories.GetUserByToken(c, token, word)
}
func GetUserByIdService(c context.Context, id int, word bool) (models.UserResponse, error) {
	return repositories.GetUserById(c, id, word)
}
func UpdateUserService(c context.Context, token string, req models.UserUpdateRequest) error {
	return repositories.UpdateUser(c, token, req)
}
func UpdateUserByIdService(c context.Context, id int, req models.UserUpdateRequest) error {
	return repositories.UpdateUserById(c, id, req)
}
func GenerateSecureToken(length int) string {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}
func NewToken(c context.Context, length int, id int) (string, error) {
	token := GenerateSecureToken(32)
	err := repositories.InsertToken(c, token, id)
	if err != nil {
		return "", err
	}
	return token, nil
}
func LoginService(c context.Context, email string, password string) (string, error) {
	user, err := repositories.GetUserByEmail(c, email, true)
	if err != nil {
		return "", err
	}
	if user.Password != password {
		return "", errors.New("wrong password")
	}
	token, err := NewToken(c, 32, user.Id)
	if err != nil {
		return "", err
	}
	return token, nil
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
	req, err := repositories.GetUserByToken(c, token, word)
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
