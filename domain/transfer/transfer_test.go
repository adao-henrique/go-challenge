package transfer

import (
	"context"
	"testing"

	"github.com/adao-henrique/go-challenge/domain/entities"
	"github.com/stretchr/testify/assert"
)

func TestCreateTransfers(t *testing.T) {
	t.Run("Should create an account transfer", func(t *testing.T) {
		t.Parallel()

		accountOrigin, err := entities.NewAccount("account-origin", "208.990.000-88", "secret-account--origin")
		accountOrigin.Balance = 1000.0
		assert.Nil(t, err)

		accountDestination, err := entities.NewAccount("account-destination", "208.990.000-88", "secret-account-destination")
		assert.Nil(t, err)

		amount := 500.00

		inputTransfer := entities.CreateTransferInput{
			AccOriginId: accountOrigin.ID,
			AccDestId:   accountDestination.ID,
			Amount:      amount,
		}

		transferMock := func(ctx context.Context, input entities.InputTransfer) (entities.Transfer, error) {
			return input.Transfer, nil
		}

		getByIDMock := func(ctx context.Context, ID string) (*entities.Account, error) {
			if ID == accountOrigin.ID {
				return &accountOrigin, nil
			}
			return &accountDestination, nil
		}

		inputInit := InputInit{TransferMock: transferMock, GetByIDMock: getByIDMock}
		transferUseCases := Init(inputInit)

		_, err = transferUseCases.transfer(context.Background(), inputTransfer)
		assert.Nil(t, err)

		assert.Equal(t, amount, accountOrigin.Balance)
		assert.Equal(t, amount, accountDestination.Balance)
	})
}
