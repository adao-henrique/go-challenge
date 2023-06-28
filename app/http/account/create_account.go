package account

import (
	"encoding/json"
	"net/http"

	"github.com/adao-henrique/go-challenge/domain/entities"
)

func (h Handler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var reqBody AccountRequest

	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(reqBody)
		return
	}

	acc, err := h.accountUseCase.CreateAccount(r.Context(), entities.CreateAccountInput{
		Name:   reqBody.Name,
		CPF:    reqBody.CPF,
		Secret: reqBody.Secret,
	})
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(reqBody)
	}

	response := AccountResponse{
		ID:        acc.ID,
		Name:      acc.Name,
		CPF:       acc.Cpf,
		Balance:   acc.Balance,
		CreatedAt: acc.CreatedAt,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)

}
