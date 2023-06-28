package login

import (
	"context"

	"github.com/adao-henrique/go-challenge/domain/entities"
)

type InputInit struct {
	GetByCPFMock func(ctx context.Context, cpf string) (entities.Account, error)
}

type AccountUseCaseMock struct {
	GetByCPFMock func(ctx context.Context, cpf string) (entities.Account, error)
	GetByIDMock  func(ctx context.Context, accountID string) (*entities.Account, error)
}

func (uc AccountUseCaseMock) GetByCPF(ctx context.Context, cpf string) (entities.Account, error) {
	return uc.GetByCPFMock(ctx, cpf)
}

func (uc AccountUseCaseMock) GetByID(ctx context.Context, accountID string) (*entities.Account, error) {
	return uc.GetByIDMock(ctx, accountID)
}

func Init(input InputInit) LoginService {

	accountUseCase := AccountUseCaseMock{
		GetByCPFMock: input.GetByCPFMock,
	}

	return LoginService{accountUseCase}
}
