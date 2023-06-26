package transfer

import (
	"context"

	"github.com/adao-henrique/go-challenge/domain/entities"
)

func (r Repository) ListTransfersByAccountID(ctx context.Context, accountOriginID string) ([]entities.Transfer, error) {
	var transfers []entities.Transfer

	rows, err := r.db.Query(ctx, "select id, account_origin_id, account_destination_id, amount, created_at from transfer")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		transfer := entities.Transfer{}
		err = rows.Scan(
			&transfer.ID,
			&transfer.AccOriginId,
			&transfer.AccDestId,
			&transfer.Amount,
			&transfer.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		transfers = append(transfers, transfer)
	}

	return transfers, nil
}
