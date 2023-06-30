package account

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// ShowAccount godoc
// @Summary      Show accounts
// @Description  get list of accounts
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Success      200  {object}  AccountResponse
// @Failure      400  {string}	string
// @Failure      404  {string}	string
// @Failure      500  {string}	string
// @Router       /account [get]
func (h Handler) GetAccounts(w http.ResponseWriter, r *http.Request) {

	accounts, err := h.accountUseCase.GetAccounts(r.Context())
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("error fetching accounts. %s", err)))
		return
	}

	response := make([]AccountResponse, len(accounts))

	for i, acc := range accounts {
		response[i] = AccountResponse{
			ID:        acc.ID,
			Name:      acc.Name,
			CPF:       acc.Cpf,
			Balance:   acc.Balance,
			CreatedAt: acc.CreatedAt,
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
