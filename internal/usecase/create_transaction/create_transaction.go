package createTransaction

import (
	"github.com/brunoan99/fc-ms-wallet/internal/entity"
	"github.com/brunoan99/fc-ms-wallet/internal/gateway"
)

type CreateTransactionInputDTO struct {
	IDFrom string
	IDTo   string
	Amount float64
}

type CreateTransactionOutputDTO struct {
	ID string
}

type CreateTransactionUseCase struct {
	TrasactionGateway gateway.TransactionGateway
	AccountGateway    gateway.AccountGateway
}

func NewCreateTransactionUseCase(trasactionGateway gateway.TransactionGateway, accountGateway gateway.AccountGateway) *CreateTransactionUseCase {
	return &CreateTransactionUseCase{
		TrasactionGateway: trasactionGateway,
		AccountGateway:    accountGateway,
	}
}

func (uc *CreateTransactionUseCase) Execute(input *CreateTransactionInputDTO) (*CreateTransactionOutputDTO, error) {
	accountFrom, err := uc.AccountGateway.FindByID(input.IDFrom)
	if err != nil {
		return nil, err
	}
	accountTo, err := uc.AccountGateway.FindByID(input.IDTo)
	if err != nil {
		return nil, err
	}
	transaction, err := entity.NewTransaction(accountFrom, accountTo, input.Amount)
	if err != nil {
		return nil, err
	}
	err = uc.TrasactionGateway.Create(transaction)
	if err != nil {
		return nil, err
	}
	return &CreateTransactionOutputDTO{
		ID: transaction.ID,
	}, nil
}
