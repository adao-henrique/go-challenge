package account

import "github.com/adao-henrique/go-challenge/domain/account"

type Handler struct {
	accountUseCase account.AccountUseCases
}

func NewHandler(accountUseCase account.AccountUseCases) Handler {
	return Handler{accountUseCase: accountUseCase}
}
