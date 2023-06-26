package transfer

import (
	"context"

	"github.com/adao-henrique/go-challenge/domain/account"
	"github.com/adao-henrique/go-challenge/domain/entities"
)

type TransferUseCases struct {
	repository      Repository
	accountUseCases account.AccountUseCasesInterFace
}

type TransferUseCasesInterFace interface {
	transfer(ctx context.Context, input entities.CreateTransferInput) (entities.Account, error)
}
