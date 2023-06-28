package account

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (h Handler) GetBalanceFromAccount(w http.ResponseWriter, r *http.Request) {

	accountID := chi.URLParam(r, "date")

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
