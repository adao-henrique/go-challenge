package transfer

import (
	"context"

	"github.com/adao-henrique/go-challenge/domain/entities"
)

type RepositoryMock struct {
	TransferMock                 func(ctx context.Context, input entities.InputTransfer) (entities.Transfer, error)
	ListTransfersByAccountIDMock func(ctx context.Context, accountOriginID string) ([]entities.Transfer, error)
}

func (r RepositoryMock) MakeTransfer(ctx context.Context, input entities.InputTransfer) (entities.Transfer, error) {
	return r.TransferMock(ctx, input)
}

func (r RepositoryMock) ListTransfersByAccountID(ctx context.Context, accountOriginID string) ([]entities.Transfer, error) {
	return r.ListTransfersByAccountIDMock(ctx, accountOriginID)
}
