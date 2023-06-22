package account

import (
	"context"

	"github.com/adao-henrique/go-challenge/domain/entities"
)

func (uc AccountUseCases) CreateAccount(ctx context.Context, input entities.CreateAccountInput) (*entities.Account, error) {

	// Check already cpf in database
	_, err := uc.repository.FindByCPF(ctx, input.CPF)
	if err != nil {
		return nil, err
	}

	// Create Account
	account, err := entities.NewAccount(input.Name, input.CPF, input.Secret)

	// Save Account
	err = uc.repository.Create(ctx, &account)
	if err != nil {
		return nil, err
	}

	return &account, nil
}
