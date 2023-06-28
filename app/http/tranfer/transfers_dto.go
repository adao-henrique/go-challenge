package tranfer

import "time"

type (
	TransferRequest struct {
		AccountDestinationID string  `json:"account_destination_id" validate:"required"`
		Amount               float64 `json:"amount" validate:"required"`
	}

	TransferResponse struct {
		ID               string    `json:"id"`
		AccDestinationID string    `json:"account_destination_id"`
		AccOriginID      string    `json:"account_origin_id"`
		Amount           float64   `json:"amount"`
		CreatedAt        time.Time `json:"created_at"`
	}
)
