package tranfer

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (h Handler) GetTransferFromUSer(w http.ResponseWriter, r *http.Request) {
	// token := r.Header.Get("authorization")
	// token = strings.Replace(token, "Bearer ", "", -1)

	// accountID, err :=h.loginService.Autenticate(token)
	// if err != nil {
	// 	w.Header().Set("Content-Type", "application/json")
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	w.Write([]byte(fmt.Sprintf("error fetching account %s. %s", accountID, err)))
	// 	return
	// }

	accountID, _ := r.Context().Value("account_origin_id").(string)

	transfers, err := h.transferUseCase.List(r.Context(), accountID)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("error fetching accounts. %s", err)))
		return
	}

	response := make([]TransferResponse, len(transfers))

	for i, transfer := range transfers {
		response[i] = TransferResponse{
			ID:               transfer.ID,
			AccDestinationID: transfer.AccDestId,
			AccOriginID:      transfer.AccOriginId,
			Amount:           transfer.Amount,
			CreatedAt:        transfer.CreatedAt,
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
