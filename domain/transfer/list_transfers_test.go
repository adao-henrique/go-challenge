package transfer

import (
	"context"
	"testing"

	"github.com/adao-henrique/go-challenge/domain/entities"
	"github.com/stretchr/testify/assert"
)

func TestListTransfers(t *testing.T) {
	t.Run("Should list transfers account", func(t *testing.T) {

		accountOrigin, _ := entities.NewAccount("account-origin", "762.337.520-27", "123456")
		accountDestination, _ := entities.NewAccount("account-destination", "762.337.520-27", "123456")
		transfer1, _ := entities.NewTransfer(accountOrigin.ID, accountDestination.ID, 100)
		transfer2, _ := entities.NewTransfer(accountOrigin.ID, accountDestination.ID, 50)
		expectTransfers := []entities.Transfer{transfer1, transfer2}

		listTransfers := func(ctx context.Context, accountOriginID string) ([]entities.Transfer, error) {
			return expectTransfers, nil
		}

		inputInit := InputInit{ListTransfersByAccIDMock: listTransfers}
		transferUseCases := Init(inputInit)

		transfers, err := transferUseCases.List(context.Background(), accountOrigin.ID)
		assert.Nil(t, err)
		if err == nil {
			assert.Equal(t, transfers, expectTransfers)
		}
	})
}
