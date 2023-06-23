package entities

import (
	"time"

	"github.com/adao-henrique/go-challenge/domain/vo"
	"github.com/adao-henrique/go-challenge/extensions"
)

type Account struct {
	ID        string
	Name      string
	Cpf       string
	Secret    string
	Balance   int
	CreatedAt time.Time
}

type CreateAccountInput struct {
	Name   string
	CPF    string
	Secret string
}

func NewAccount(name string, cpf string, secret string) (Account, error) {
	hashed, err := extensions.HASH(secret)
	if err != nil {
		return Account{}, err
	}
	account := Account{
		ID:        vo.UUID(),
		Name:      name,
		Cpf:       cpf,
		Secret:    *hashed,
		Balance:   0,
		CreatedAt: time.Now(),
	}
	return account, nil
}
