package gateway

import "github.com/brunoan99/fc-ms-wallet/internal/entity"

type AccountGateway interface {
	FindByID(id string) (*entity.Account, error)
	Save(account *entity.Account) error
}
