package http

import (
	"github.com/adao-henrique/go-challenge/app/http/account"
	"github.com/adao-henrique/go-challenge/app/http/login"
	"github.com/adao-henrique/go-challenge/app/http/tranfer"
	"github.com/adao-henrique/go-challenge/extensions"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth/v5"
)

type API struct {
	accountHandler  account.Handler
	transferHandler tranfer.Handler
	loginHandler    login.Handler
	r               *chi.Mux
}

var tokenAuth *jwtauth.JWTAuth

func (a API) AccountAPI() {

	a.r.Route("/account", func(rs chi.Router) {
		rs.Use(jwtauth.Verifier(tokenAuth))
		rs.Get("/", a.accountHandler.GetAccounts)
		rs.Get("/{account_id}/balance", a.accountHandler.GetBalanceFromAccount)
	})

	a.r.Route("/account", func(rs chi.Router) {
		rs.Post("/", a.accountHandler.CreateAccount)
	})
}

func (a API) TranferAPI() {
	a.r.Route("/transfer", func(rs chi.Router) {
		rs.Use(jwtauth.Verifier(tokenAuth))
		rs.Get("/", a.transferHandler.GetTransferFromUSer)
		rs.Post("/", a.transferHandler.MakeTransfer)
	})
}

func (a API) LoginAPI() {
	a.r.Post("/login", a.loginHandler.Login)
}

func NewApi(
	r *chi.Mux,
	accountHandler account.Handler,
	transferHandler tranfer.Handler,
	loginHandler login.Handler,
) API {

	tokenAuth = extensions.GetJWTAuth()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	api := API{
		accountHandler,
		transferHandler,
		loginHandler,
		r,
	}

	api.AccountAPI()
	api.TranferAPI()
	api.LoginAPI()

	return api
}
