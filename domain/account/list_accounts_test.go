package account

import (
	"context"
	"errors"
	"testing"

	"github.com/adao-henrique/go-challenge/domain/entities"
	"github.com/stretchr/testify/assert"
)

func TestListAccounts(t *testing.T) {
	t.Run("Should list accounts", func(t *testing.T) {
		t.Parallel()
		secretPassword := "secret_password"

		account, err := entities.NewAccount("name test", "208.856.780-10", secretPassword)
		assert.Nil(t, err)

		expectedAcc := []entities.Account{account}

		inputInit := InputInit{
			GetAccountsMock: func(ctx context.Context) ([]entities.Account, error) {
				return expectedAcc, nil
			},
		}

		accountUseCases := Init(inputInit)
		accounts, err := accountUseCases.GetAccounts(context.Background())
		assert.Nil(t, err)

		assert.Equal(t, accounts, expectedAcc)
	})

	t.Run("Erros should be the same", func(t *testing.T) {
		t.Parallel()
		expectErr := errors.New("unable to fetch accounts")
		inputInit := InputInit{
			GetAccountsMock: func(ctx context.Context) ([]entities.Account, error) {
				return nil, expectErr
			},
		}

		accountUseCases := Init(inputInit)
		_, err := accountUseCases.GetAccounts(context.Background())

		assert.ErrorIs(t, err, expectErr, "Erros should be the same")
	})
}
