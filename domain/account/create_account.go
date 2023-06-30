package account

import (
	"context"
	"log"

	"github.com/adao-henrique/go-challenge/domain/entities"
)

func (uc AccountUseCases) CreateAccount(ctx context.Context, input entities.CreateAccountInput) (entities.Account, error) {

	// Check already cpf in database
	account, err := uc.repository.GetByCPF(ctx, input.CPF)
	if account.ID != "" {
		log.Printf("error to get by cpf ", account.ID == "")
		return entities.Account{}, err
	}

	// Create Account
	account, err = entities.NewAccount(input.Name, input.CPF, input.Secret)
	if err != nil {
		log.Printf("error to create account ", err)
		return entities.Account{}, err
	}

	// Save Account
	err = uc.repository.Create(ctx, account)
	if err != nil {
		log.Printf("error to save account %s, %s", err, account)
		return entities.Account{}, err
	}

	return account, nil
}
