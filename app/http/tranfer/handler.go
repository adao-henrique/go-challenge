package tranfer

import (
	"github.com/adao-henrique/go-challenge/domain/transfer"
)

type Handler struct {
	transferUseCase transfer.TransferUseCases
}

func NewHandler(transferUseCase transfer.TransferUseCases) Handler {
	return Handler{transferUseCase: transferUseCase}
}
