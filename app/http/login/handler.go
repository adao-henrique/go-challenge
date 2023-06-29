package login

import "github.com/adao-henrique/go-challenge/domain/login"

type Handler struct {
	loginService login.LoginService
}

func NewHandler(loginService login.LoginService) Handler {
	return Handler{loginService: loginService}
}
