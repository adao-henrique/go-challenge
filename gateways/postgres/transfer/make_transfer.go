package transfer

import (
	"context"

	"github.com/adao-henrique/go-challenge/domain/entities"
)

func (r Repository) MakeTransfer(ctx context.Context, input entities.InputTransfer) (entities.Transfer, error) {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return entities.Transfer{}, err
	}

	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		} else {
			tx.Commit(ctx)
		}
	}()

	sql := `insert into transfer(id, account_origin_id, account_destination_id, amount, created_at) Values ($1, $2, $3, $4, $5)`

	_, err = tx.Exec(
		ctx,
		sql,
		input.Transfer.ID,
		input.Transfer.AccOriginId,
		input.Transfer.AccDestId,
		input.Transfer.Amount,
		input.Transfer.CreatedAt,
	)

	if err != nil {
		return entities.Transfer{}, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return entities.Transfer{}, err
	}

	return input.Transfer, nil

}
