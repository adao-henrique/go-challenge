package entities

import (
	"errors"
	"time"

	"github.com/adao-henrique/go-challenge/domain/vo"
	"github.com/adao-henrique/go-challenge/extensions"
)

type Account struct {
	ID        string
	Name      string
	Cpf       string
	Secret    string
	Balance   float64
	CreatedAt time.Time
}

func (a *Account) WithdrawBalance(amount float64) error {
	if amount <= 0 {
		return errors.New("invalid amount value")
	}

	if amount > a.Balance {
		return errors.New("the account has insufficient funds")
	}

	a.Balance -= amount

	return nil
}

func (a *Account) DepositBalance(amount float64) error {
	if amount <= 0 {
		return errors.New("invalid amount value")
	}

	a.Balance += amount

	return nil
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
		Balance:   1000.0,
		CreatedAt: time.Now(),
	}
	return account, nil
}
