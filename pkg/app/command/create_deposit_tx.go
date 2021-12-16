package command

import (
	"context"

	"github.com/hsequeda/money_processing_service_example/pkg/domain"
)

type CreateDepositTx struct {
	To     string
	Amount uint
}

type CreateDepositTxHandler struct {
	txRepository domain.TxService
}

func (h CreateDepositTxHandler) Handle(ctx context.Context, cmd CreateDepositTx) error {
	return h.txRepository.Deposit(ctx, cmd.To, cmd.Amount)
}
