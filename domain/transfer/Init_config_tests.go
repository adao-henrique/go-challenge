package transfer

import (
	"context"

	"github.com/adao-henrique/go-challenge/domain/entities"
)

type InputInit struct {
	TransferMock             func(ctx context.Context, input entities.InputTransfer) (entities.Transfer, error)
	GetByIDMock              func(ctx context.Context, cpf string) (entities.Account, error)
	ListTransfersByAccIDMock func(ctx context.Context, accountOriginID string) ([]entities.Transfer, error)
}

type AccountUseCaseMock struct {
	GetByIDMock func(ctx context.Context, cpf string) (entities.Account, error)
}

func (uc AccountUseCaseMock) GetByCPF(ctx context.Context, cpf string) (entities.Account, error) {
	return entities.Account{}, nil
}

func (uc AccountUseCaseMock) GetByID(ctx context.Context, accountID string) (entities.Account, error) {
	return uc.GetByIDMock(ctx, accountID)
}

func Init(input InputInit) TransferUseCases {
	repositoryMock := RepositoryMock{
		TransferMock:                 input.TransferMock,
		ListTransfersByAccountIDMock: input.ListTransfersByAccIDMock,
	}
	accountUseCaseMock := AccountUseCaseMock{
		GetByIDMock: input.GetByIDMock,
	}
	return TransferUseCases{repositoryMock, accountUseCaseMock}
}
