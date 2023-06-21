package account

import (
	"context"
	"errors"
	"testing"

	"github.com/adao-henrique/go-challenge/domain/entities"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

type InputInit struct {
	CreateMock    func(ctx context.Context, account *entities.Account) error
	FindByCPFMock func(ctx context.Context, cpf string) (*entities.Account, error)
}

type RepositoryMock struct {
	CreateMock    func(ctx context.Context, account *entities.Account) error
	FindByCPFMock func(ctx context.Context, cpf string) (*entities.Account, error)
}

func (r RepositoryMock) Create(ctx context.Context, account *entities.Account) error {
	return r.CreateMock(ctx, account)
}
func (r RepositoryMock) FindByCPF(ctx context.Context, cpf string) (*entities.Account, error) {
	return r.FindByCPFMock(ctx, cpf)
}

func Init(input InputInit) AccountUseCases {

	return AccountUseCases{RepositoryMock{
		CreateMock:    input.CreateMock,
		FindByCPFMock: input.FindByCPFMock,
	}}
}

func TestCreateAccouunt(t *testing.T) {
	t.Run("Should create an account", func(t *testing.T) {
		t.Parallel()
		inputInit := InputInit{
			CreateMock: func(ctx context.Context, account *entities.Account) error {
				return nil
			},

			FindByCPFMock: func(ctx context.Context, cpf string) (*entities.Account, error) {
				return nil, nil
			},
		}
		accountUseCases := Init(inputInit)

		secret := "secret-password"
		input := &entities.CreateAccountInput{
			Name:   "name test",
			CPF:    "451.856.400-65",
			Secret: secret,
		}

		hashed, _ := bcrypt.GenerateFromPassword([]byte(secret), 8)
		expect := &entities.Account{
			Name:    "name test",
			Cpf:     "451.856.400-65",
			Secret:  string(hashed),
			Balance: 0,
		}

		account, err := accountUseCases.CreateAccount(context.Background(), *input)
		assert.Nil(t, err)

		assert.Equal(t, expect.Name, account.Name, "Names shold be equals")
		assert.Equal(t, expect.Cpf, account.Cpf, "CPFs shold be equals")
		assert.Equal(t, expect.Balance, account.Balance, "Balances shold be equals")

		err = bcrypt.CompareHashAndPassword([]byte(expect.Secret), []byte(secret))
		if err != nil {
			assert.Fail(t, "Erro to compare secret")
		}
		assert.Nil(t, err, "Secrets shold be equals")

	})

	t.Run("Erro: Already is used this CPF", func(t *testing.T) {
		t.Parallel()
		ErrorCPFAlreadyUsed := errors.New("This CPF Already is used")

		inputInit := InputInit{
			CreateMock: func(ctx context.Context, account *entities.Account) error {
				return nil
			},

			FindByCPFMock: func(ctx context.Context, cpf string) (*entities.Account, error) {
				return nil, ErrorCPFAlreadyUsed
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

		assert.Equal(t, ErrorCPFAlreadyUsed, err)

	})

	t.Run("Erro: To save acccount", func(t *testing.T) {
		t.Parallel()
		ErrorCreateAccont := errors.New("Error to save acccount")

		inputInit := InputInit{
			CreateMock: func(ctx context.Context, account *entities.Account) error {
				return ErrorCreateAccont
			},

			FindByCPFMock: func(ctx context.Context, cpf string) (*entities.Account, error) {
				return nil, nil
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

		assert.Equal(t, ErrorCreateAccont, err)

	})

}
