package product

import (
	use_case "github.com/sercanakmaz/go-boilerplate-v3/pkg/ddd"
)

type UpdateStockCommand struct {
	Sku      string                    `json:"sku"`
	Stock    int                       `json:"stock"`
	identity *use_case.UseCaseIdentity `json:"_"`
}

func (self *UpdateStockCommand) GetIdentity() *use_case.UseCaseIdentity {
	return self.identity
}
