package account

import (
	"context"
	"errors"
	"testing"

	"github.com/adao-henrique/go-challenge/domain/entities"
	"github.com/adao-henrique/go-challenge/extensions"
	"github.com/stretchr/testify/assert"
)

func TestCreateAccouunt(t *testing.T) {
	t.Run("Should create an account", func(t *testing.T) {
		t.Parallel()
		inputInit := InputInit{
			CreateMock: func(ctx context.Context, account entities.Account) error {
				return nil
			},

			FindByCPFMock: func(ctx context.Context, cpf string) (entities.Account, error) {
				return entities.Account{}, nil
			},
		}
		accountUseCases := Init(inputInit)

		secret := "secret-password"
		input := &entities.CreateAccountInput{
			Name:   "name test",
			CPF:    "451.856.400-65",
			Secret: secret,
		}

		hashed, _ := extensions.HASH(secret)
		expect := &entities.Account{
			Name:    "name test",
			Cpf:     "451.856.400-65",
			Secret:  *hashed,
			Balance: 0,
		}

		account, err := accountUseCases.CreateAccount(context.Background(), *input)
		assert.Nil(t, err)

		assert.Equal(t, expect.Name, account.Name, "Names shold be equals")
		assert.Equal(t, expect.Cpf, account.Cpf, "CPFs shold be equals")
		assert.Equal(t, expect.Balance, account.Balance, "Balances shold be equals")

		err = extensions.CompareHashAndValue(expect.Secret, secret)
		if err != nil {
			assert.Fail(t, "Erro to compare secret")
		}
		assert.Nil(t, err, "Secrets shold be equals")

	})

	t.Run("Erro: Already is used this CPF", func(t *testing.T) {
		t.Parallel()
		ErrorCPFAlreadyUsed := errors.New("This CPF Already is used")

		inputInit := InputInit{
			CreateMock: func(ctx context.Context, account entities.Account) error {
				return nil
			},

			FindByCPFMock: func(ctx context.Context, cpf string) (entities.Account, error) {
				return entities.Account{}, ErrorCPFAlreadyUsed
			},
		}

		accountUseCases := Init(inputInit)

		secret := "secret-password"
		input := &entities.CreateAccountInput{
			Name:   "name test",
			CPF:    "451.856.400-65",
			Secret: secret,
		}

		_, err := accountUseCases.CreateAccount(context.Background(), *input)

		assert.ErrorIs(t, err, ErrorCPFAlreadyUsed)

	})

	t.Run("Erro: To save acccount", func(t *testing.T) {
		t.Parallel()
		ErrorCreateAccont := errors.New("Error to save acccount")

		inputInit := InputInit{
			CreateMock: func(ctx context.Context, account entities.Account) error {
				return ErrorCreateAccont
			},

			FindByCPFMock: func(ctx context.Context, cpf string) (entities.Account, error) {
				return entities.Account{}, nil
			},
		}

		accountUseCases := Init(inputInit)

		secret := "secret-password"
		input := &entities.CreateAccountInput{
			Name:   "name test",
			CPF:    "451.856.400-65",
			Secret: secret,
		}

		_, err := accountUseCases.CreateAccount(context.Background(), *input)

		assert.ErrorIs(t, err, ErrorCreateAccont)

	})

}
