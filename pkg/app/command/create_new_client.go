package command

import (
	"context"

	"github.com/hsequeda/money_processing_service_example/pkg/domain"
)

type CreateNewClient struct {
	ID string
}

type CreateNewClientHandler struct {
	clientRepository domain.ClientRepository
}

func NewCreateNewClientHandler(clientRepository domain.ClientRepository) CreateNewClientHandler {
	return CreateNewClientHandler{clientRepository}
}

func (h CreateNewClientHandler) Handle(ctx context.Context, cmd CreateNewClient) error {
	return h.clientRepository.NewClient(ctx, cmd.ID)
}
