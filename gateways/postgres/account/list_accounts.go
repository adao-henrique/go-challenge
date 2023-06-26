package account

import (
	"context"

	"github.com/adao-henrique/go-challenge/domain/entities"
)

func (r Repository) GetAccounts(ctx context.Context) ([]entities.Account, error) {
	var accounts []entities.Account

	rows, err := r.db.Query(ctx, "select id, name, cpf, balance, created_at from public.account")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		account := entities.Account{}
		err = rows.Scan(
			&account.ID,
			&account.Name,
			&account.Cpf,
			&account.Balance,
			&account.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}

	return accounts, nil
}
