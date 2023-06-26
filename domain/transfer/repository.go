package transfer

import (
	"context"

	"github.com/adao-henrique/go-challenge/domain/entities"
)

type Repository interface {
	MakeTransfer(ctx context.Context, input entities.InputTransfer) (entities.Transfer, error)
}
