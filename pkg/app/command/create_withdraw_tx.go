package command

import (
	"context"

	"github.com/hsequeda/money_processing_service_example/pkg/domain"
)

type CreateWithdrawTx struct {
	From   string
	Amount uint
}

type CreateWithdrawTxHandler struct {
	txRepository domain.TxService
}

func (h CreateWithdrawTxHandler) Handle(ctx context.Context, cmd CreateWithdrawTx) error {
	return h.txRepository.Withdraw(ctx, cmd.From, cmd.Amount)
}
