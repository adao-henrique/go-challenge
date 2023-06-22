package account

import (
	"context"

	"github.com/adao-henrique/go-challenge/domain/entities"
)

func (accountUseCases AccountUseCases) GetAccounts(ctx context.Context) (*[]entities.Account, error) {

	accouts, err := accountUseCases.repository.GetAccounts(ctx)
	if err != nil {
		return nil, err
	}

	return accouts, nil
}
