package order

import (
	"github.com/sercanakmaz/go-boilerplate-v3/models/shared"
	use_case "github.com/sercanakmaz/go-boilerplate-v3/pkg/ddd/use-case"
)

type CreateOrderCommand struct {
	OrderNumber string                    `json:"orderNumber"`
	Price       shared.Money              `json:"price"`
	identity    *use_case.UseCaseIdentity `json:"_"`
}

func (self *CreateOrderCommand) GetIdentity() *use_case.UseCaseIdentity {
	return self.identity
}
