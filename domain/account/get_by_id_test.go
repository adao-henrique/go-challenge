package account

import (
	"context"
	"errors"
	"testing"

	"github.com/adao-henrique/go-challenge/domain/entities"
	"github.com/stretchr/testify/assert"
)

func TestGetAccountByID(t *testing.T) {
	t.Run("Should create an account", func(t *testing.T) {
		t.Parallel()
		secretPassword := "secret_password"
		expectedAcc, err := entities.NewAccount("name test", "208.856.780-10", secretPassword)
		assert.Nil(t, err)

		inputInit := InputInit{
			GetByIDMock: func(ctx context.Context, accountID string) (*entities.Account, error) {
				return &expectedAcc, nil
			},
		}

		accountUseCases := Init(inputInit)

		account, err := accountUseCases.GetByID(context.Background(), expectedAcc.ID)
		assert.Nil(t, err)
		assert.Equal(t, expectedAcc, *account)
	})

	t.Run("Erros should be the same", func(t *testing.T) {
		t.Parallel()
		expectErr := errors.New("unable to fetch account")
		inputInit := InputInit{
			GetByIDMock: func(ctx context.Context, accountID string) (*entities.Account, error) {
				return nil, expectErr
			},
		}

		accountUseCases := Init(inputInit)
		_, err := accountUseCases.GetByID(context.Background(), "")

		assert.Equal(t, err, expectErr, "Erros should be the same")
	})
}
