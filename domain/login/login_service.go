package login

import "github.com/adao-henrique/go-challenge/domain/account"

type LoginService struct {
	accountUseCases account.AccountUseCasesInterFace
}
