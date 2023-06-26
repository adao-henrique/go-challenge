package transfer

import (
	"context"

	"github.com/adao-henrique/go-challenge/domain/entities"
)

func (uc TransferUseCases) List(ctx context.Context, accountOriginID string) ([]entities.Transfer, error) {

	transfers, err := uc.repository.ListTransfersByAccountID(ctx, accountOriginID)
	if err != nil {
		return nil, err
	}

	return transfers, nil
}
