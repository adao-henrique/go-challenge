package tranfer

import (
	"encoding/json"
	"net/http"

	"github.com/adao-henrique/go-challenge/domain/entities"
)

func (h Handler) MakeTransfer(w http.ResponseWriter, r *http.Request) {
	var reqBody TransferRequest

	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(reqBody)
		return
	}

	accountID, _ := r.Context().Value("account_origin_id").(string)

	transfer, err := h.transferUseCase.Transfer(r.Context(), entities.CreateTransferInput{
		AccOriginId: accountID,
		AccDestId:   reqBody.AccountDestinationID,
		Amount:      reqBody.Amount,
	})

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(reqBody)
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
