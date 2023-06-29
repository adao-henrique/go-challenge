package account

import (
	"context"

	"github.com/adao-henrique/go-challenge/domain/entities"
)

type Repository interface {
	Create(ctx context.Context, account entities.Account) error
	GetByCPF(ctx context.Context, cpf string) (entities.Account, error)
	GetAccounts(ctx context.Context) ([]entities.Account, error)
	GetByID(ctx context.Context, accountID string) (entities.Account, error)
}
