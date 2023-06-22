package account

import (
	"context"

	"github.com/adao-henrique/go-challenge/domain/entities"
)

type Repository interface {
	Create(ctx context.Context, account *entities.Account) error
	FindByCPF(ctx context.Context, cpf string) (*entities.Account, error)
	GetAccounts(ctx context.Context) (*[]entities.Account, error)
}
