package product

import (
	"github.com/sercanakmaz/go-boilerplate-v3/models/shared"
	use_case "github.com/sercanakmaz/go-boilerplate-v3/pkg/ddd"
)

type CreateProductCommand struct {
	Sku          string                    `json:"sku"`
	Name         string                    `json:"name"`
	Brand        string                    `json:"brand"`
	InitialStock int                       `json:"initialStock"`
	CategoryID   int                       `json:"categoryID"`
	Price        shared.Money              `json:"price"`
	identity     *use_case.UseCaseIdentity `json:"_"`
}

func (self *CreateProductCommand) GetIdentity() *use_case.UseCaseIdentity {
	return self.identity
}
