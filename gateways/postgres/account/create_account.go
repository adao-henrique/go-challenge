package account

import (
	"context"
	"log"

	"github.com/adao-henrique/go-challenge/domain/entities"
)

func (r Repository) Create(ctx context.Context, account entities.Account) error {

	tx, err := r.db.Begin(ctx)
	if err != nil {
		log.Printf("error to create transaction ", err)
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		} else {
			tx.Commit(ctx)
		}
	}()

	sql := `insert into account(id, name, cpf, secret, balance, created_at) values ($1,$2,$3,$4,$5,$6)`
	_, err = tx.Exec(
		ctx,
		sql,
		account.ID,
		account.Name,
		account.Cpf,
		account.Secret,
		account.Balance,
		account.CreatedAt)

	return err
}
