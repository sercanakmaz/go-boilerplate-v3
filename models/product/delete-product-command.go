package product

import (
	use_case "github.com/sercanakmaz/go-boilerplate-v3/pkg/ddd"
)

type DeleteProductCommand struct {
	Sku      string                    `json:"sku"`
	identity *use_case.UseCaseIdentity `json:"_"`
}

func (self *DeleteProductCommand) GetIdentity() *use_case.UseCaseIdentity {
	return self.identity
}
