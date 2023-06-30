package account

import (
	"context"

	"github.com/adao-henrique/go-challenge/domain/entities"
)

type AccountUseCases struct {
	repository Repository
}

type AccountUseCasesInterFace interface {
	GetByCPF(ctx context.Context, cpf string) (entities.Account, error)
	GetByID(ctx context.Context, accountID string) (entities.Account, error)
}

func NewUseCase(repository Repository) AccountUseCases {
	return AccountUseCases{
		repository: repository,
	}
}
