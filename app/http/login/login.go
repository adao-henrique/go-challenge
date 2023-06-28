package login

import (
	"encoding/json"
	"net/http"

	"github.com/adao-henrique/go-challenge/domain/login"
)

func (h Handler) Login(w http.ResponseWriter, r *http.Request) {

	var reqBody LoginRequest

	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(reqBody)
		return
	}

	token, err := h.loginService.Login(r.Context(), login.LoginUserInput{
		Cpf:    reqBody.Cpf,
		Secret: reqBody.Secret,
	})
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(reqBody)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(token)
}
