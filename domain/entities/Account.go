package entities

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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

func NewAccount(name string, cpf string, secret string) *Account {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(secret), 8)

	return &Account{
		ID:        uuid.NewString(),
		Name:      name,
		Cpf:       cpf,
		Secret:    string(hashed),
		Balance:   0,
		CreatedAt: time.Now(),
	}
}
