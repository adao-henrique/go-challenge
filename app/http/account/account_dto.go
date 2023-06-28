package account

import "time"

type (
	AccountRequest struct {
		Name   string `json:"name" validate:"required"`
		CPF    string `json:"cpf" validate:"required"`
		Secret string `json:"secret" validate:"required"`
	}

	AccountResponse struct {
		ID        string    `json:"id"`
		Name      string    `json:"name"`
		CPF       string    `json:"cpf"`
		Balance   float64   `json:"balance"`
		CreatedAt time.Time `json:"created_at"`
	}

	BalanceResponse struct {
		Balance float64 `json:"balance"`
	}
)
