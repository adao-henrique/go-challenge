package login

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/adao-henrique/go-challenge/domain/login"
)

func (h Handler) Login(w http.ResponseWriter, r *http.Request) {

	var reqBody LoginRequest

	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("error to read login request: %s", err)))
		return
	}

	token, err := h.loginService.Login(r.Context(), login.LoginUserInput{
		Cpf:    reqBody.Cpf,
		Secret: reqBody.Secret,
	})
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("login error: %s", err)))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(LoginResponse{Token: token})
}
