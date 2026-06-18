package models

type Token struct {
	ID    int
	Token string
}
type TokenCreateRequest struct {
	Password string
	Email    string
}
type TokenResponse struct {
	Token  string
	UserId int
}
type TokenCheck struct {
	Token string
}
