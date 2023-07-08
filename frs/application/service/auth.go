package service

type AuthService interface {
	DoAuthenticate(username string, password string) bool
}
