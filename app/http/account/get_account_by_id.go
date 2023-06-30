package account

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// ShowAccount godoc
// @Summary      Show balance from account
// @Description  get string by ID
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        account_id   path      int  true  "Account ID"
// @Success      200  {object}  BalanceResponse
// @Failure      400  {string}	string
// @Failure      404  {string}	string
// @Failure      500  {string}	string
// @Router       /account/{account_id}/balance [get]
func (h Handler) GetBalanceFromAccount(w http.ResponseWriter, r *http.Request) {

	accountID := chi.URLParam(r, "account_id")

	account, err := h.accountUseCase.GetByID(r.Context(), accountID)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("error fetching account %s. %s", accountID, err)))
		return
	}

	response := BalanceResponse{
		account.Balance,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
