package transfer

import (
	"context"

	"github.com/adao-henrique/go-challenge/domain/entities"
	"github.com/jackc/pgx/v5"
)

func (r Repository) updateBalanceFromAccount(ctx context.Context, tx pgx.Tx, acc entities.Account) error {
	sql := `update account set balance=$1 where id=$2`

	_, err := tx.Exec(ctx, sql, acc.Balance, acc.ID)

	return err
}
