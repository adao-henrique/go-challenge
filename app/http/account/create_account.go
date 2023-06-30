package account

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/adao-henrique/go-challenge/domain/entities"
)

// ShowAccount godoc
// @Summary      Create account
// @Description  Create account
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Success      200  {object}  AccountResponse
// @Failure      400  {string}	string
// @Failure      404  {string}	string
// @Failure      500  {string}	string
// @Router       /account/ [post]
func (h Handler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var reqBody AccountRequest

	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("error to create account: %s", err)))
		return
	}

	acc, err := h.accountUseCase.CreateAccount(r.Context(), entities.CreateAccountInput{
		Name:   reqBody.Name,
		CPF:    reqBody.CPF,
		Secret: reqBody.Secret,
	})
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("error to create account: %s", err)))
		return
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
