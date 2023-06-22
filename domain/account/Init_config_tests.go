package account

import (
	"context"

	"github.com/adao-henrique/go-challenge/domain/entities"
)

type InputInit struct {
	CreateMock      func(ctx context.Context, account *entities.Account) error
	FindByCPFMock   func(ctx context.Context, cpf string) (*entities.Account, error)
	GetAccountsMock func(ctx context.Context) (*[]entities.Account, error)
}

func Init(input InputInit) AccountUseCases {

	return AccountUseCases{RepositoryMock{
		CreateMock:      input.CreateMock,
		FindByCPFMock:   input.FindByCPFMock,
		GetAccountsMock: input.GetAccountsMock,
	}}
}
