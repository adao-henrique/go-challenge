package account

import (
	"context"

	"github.com/adao-henrique/go-challenge/domain/entities"
)

func (accountUseCases AccountUseCases) GetByID(ctx context.Context, accountID string) (*entities.Account, error) {
	account, err := accountUseCases.repository.GetByID(ctx, accountID)
	if err != nil {
		return nil, err
	}

	return account, err
}
