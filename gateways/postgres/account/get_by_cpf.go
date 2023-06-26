package account

import (
	"context"

	"github.com/adao-henrique/go-challenge/domain/entities"
)

func (r Repository) GetByCPF(ctx context.Context, cpf string) (entities.Account, error) {

	account := entities.Account{}
	err := r.db.QueryRow(ctx, "select id, name, cpf, balance, created_at from public.account where cpf=$1", cpf).Scan(
		&account.ID,
		&account.Name,
		&account.Cpf,
		&account.Balance,
		&account.CreatedAt,
	)

	if err != nil {
		return entities.Account{}, err
	}

	return account, nil
}
