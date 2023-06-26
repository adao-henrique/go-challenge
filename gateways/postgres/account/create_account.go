package account

import (
	"context"

	"github.com/adao-henrique/go-challenge/domain/entities"
)

func (r Repository) Create(ctx context.Context, account entities.Account) error {

	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		} else {
			tx.Commit(ctx)
		}
	}()

	sql := `insert into public.account(id, name, cpf, balance, created_at) values ($1,$2,$3,$4,$5)`
	_, err = tx.Exec(
		ctx,
		sql,
		account.ID,
		account.Name,
		account.Cpf,
		account.Balance,
		account.CreatedAt)

	return err
}
