package account

import (
	"context"

	"github.com/adao-henrique/go-challenge/domain/entities"
)

func (accountUseCases AccountUseCases) GetByCPF(ctx context.Context, cpf string) (entities.Account, error) {
	account, err := accountUseCases.repository.GetByCPF(ctx, cpf)
	if err != nil {
		return entities.Account{}, err
	}

	return account, err
}
