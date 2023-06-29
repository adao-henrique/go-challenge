package tranfer

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/adao-henrique/go-challenge/domain/entities"
	"github.com/adao-henrique/go-challenge/extensions"
)

func (h Handler) MakeTransfer(w http.ResponseWriter, r *http.Request) {
	var reqBody TransferRequest

	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("error parse request")))
		return
	}

	bearerToken := r.Header.Clone().Get("Authorization")
	token := strings.Split(bearerToken, " ")[1]
	claims, err := extensions.GetClaims(token)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("error to parse token %s, %s", token, err)
		w.Write([]byte(fmt.Sprintf("error parse request")))
		return
	}

	transfer, err := h.transferUseCase.Transfer(r.Context(), entities.CreateTransferInput{
		AccOriginId: fmt.Sprintf("%v", claims["account_origin_id"]),
		AccDestId:   reqBody.AccountDestinationID,
		Amount:      reqBody.Amount,
	})

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("error make trasfer %s, %s, %s", claims["account_origin_id"], transfer, reqBody)
		w.Write([]byte("error make trasfer"))
		return
	}

	response := TransferResponse{
		ID:               transfer.ID,
		AccDestinationID: transfer.AccDestId,
		AccOriginID:      transfer.AccOriginId,
		Amount:           transfer.Amount,
		CreatedAt:        transfer.CreatedAt,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
