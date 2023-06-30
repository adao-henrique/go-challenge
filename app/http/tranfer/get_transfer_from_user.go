package tranfer

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// ShowTransfer godoc
// @Summary      Show all transfer from user
// @Description  Show all transfer from user
// @Tags         transfer
// @Accept       json
// @Produce      json
// @Success      200  {array}  	TransferResponse
// @Failure      400  {string}	string
// @Failure      404  {string}	string
// @Failure      500  {string}	string
// @Router       /transfer/ 		[get]
func (h Handler) GetTransferFromUSer(w http.ResponseWriter, r *http.Request) {
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
