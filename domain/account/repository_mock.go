package account

import (
	"context"

	"github.com/adao-henrique/go-challenge/domain/entities"
)

type RepositoryMock struct {
	CreateMock      func(ctx context.Context, account entities.Account) error
	FindByCPFMock   func(ctx context.Context, cpf string) (entities.Account, error)
	GetAccountsMock func(ctx context.Context) ([]entities.Account, error)
	GetByIDMock     func(ctx context.Context, accountID string) (entities.Account, error)
}

func (r RepositoryMock) Create(ctx context.Context, account entities.Account) error {
	return r.CreateMock(ctx, account)
}
func (r RepositoryMock) GetByCPF(ctx context.Context, cpf string) (entities.Account, error) {
	return r.FindByCPFMock(ctx, cpf)
}
func (r RepositoryMock) GetAccounts(ctx context.Context) ([]entities.Account, error) {
	return r.GetAccountsMock(ctx)
}
func (r RepositoryMock) GetByID(ctx context.Context, accountID string) (entities.Account, error) {
	return r.GetByIDMock(ctx, accountID)
}
