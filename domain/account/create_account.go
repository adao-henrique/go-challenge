package account

import (
	"context"

	"github.com/adao-henrique/go-challenge/domain/entities"
)

func (accountUseCases AccountUseCases) CreateAccount(ctx context.Context, input entities.CreateAccountInput) (*entities.Account, error) {

	// Check already cpf in database
	account, err := accountUseCases.repository.FindByCPF(ctx, input.CPF)
	if err != nil {
		return nil, err
	}

	// Create Account
	account = entities.NewAccount(input.Name, input.CPF, input.Secret)

	// Save Account
	err = accountUseCases.repository.Create(ctx, account)
	if err != nil {
		return nil, err
	}

	return account, nil
}
