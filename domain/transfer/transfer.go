package transfer

import (
	"context"

	"github.com/adao-henrique/go-challenge/domain/entities"
)

func (uc TransferUseCases) Transfer(ctx context.Context, input entities.CreateTransferInput) (entities.Transfer, error) {

	accDest, err := uc.accountUseCases.GetByID(ctx, input.AccDestId)
	if err != nil {
		return entities.Transfer{}, err
	}

	accOrig, err := uc.accountUseCases.GetByID(ctx, input.AccOriginId)
	if err != nil {
		return entities.Transfer{}, err
	}

	err = accDest.DepositBalance(input.Amount)
	if err != nil {
		return entities.Transfer{}, err
	}

	err = accOrig.WithdrawBalance(input.Amount)
	if err != nil {
		return entities.Transfer{}, err
	}

	transfer, err := entities.NewTransfer(input.AccOriginId, input.AccDestId, input.Amount)
	if err != nil {
		return entities.Transfer{}, err
	}

	transfer, err = uc.repository.MakeTransfer(ctx, entities.InputTransfer{
		AccOrigin: accOrig,
		AccDest:   accDest,
		Transfer:  transfer,
	})

	if err != nil {
		return entities.Transfer{}, err
	}

	return transfer, nil
}
