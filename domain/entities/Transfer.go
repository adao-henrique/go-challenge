package entities

import (
	"time"

	"github.com/adao-henrique/go-challenge/domain/vo"
)

type Transfer struct {
	ID          string
	AccOriginId string
	AccDestId   string
	Amount      float64
	CreatedAt   time.Time
}

func NewTransfer(accOriginId string, accDestId string, amount float64) (Transfer, error) {
	transfer := Transfer{
		ID:          vo.UUID(),
		AccOriginId: accOriginId,
		AccDestId:   accDestId,
		Amount:      amount,
		CreatedAt:   time.Now(),
	}
	return transfer, nil
}

type CreateTransferInput struct {
	AccOriginId string
	AccDestId   string
	Amount      float64
}

type InputTransfer struct {
	AccOrigin Account
	AccDest   Account
	Transfer  Transfer
}
