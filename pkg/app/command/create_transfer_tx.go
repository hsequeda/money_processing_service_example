package command

import (
	"context"

	"github.com/hsequeda/money_processing_service_example/pkg/domain"
)

type CreateTransferTx struct {
	To     string
	From   string
	Amount uint
}

type CreateTransferTxHandler struct {
	txRepository domain.TxService
}

func (h CreateTransferTxHandler) Handle(ctx context.Context, cmd CreateTransferTx) error {
	return h.txRepository.Transfer(ctx, cmd.From, cmd.To, cmd.Amount)
}
