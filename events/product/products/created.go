package products

import (
	"github.com/sercanakmaz/go-boilerplate-v3/models/shared"
	"github.com/sercanakmaz/go-boilerplate-v3/pkg/ddd"
)

type Created struct {
	Id         string       `json:"id"`
	Sku        string       `json:"sku"`
	Name       string       `json:"Name"`
	Stock      int          `json:"Stock"`
	Price      shared.Money `json:"Price"`
	CategoryId int          `json:"CategoryId"`
	ddd.IBaseEvent
}

func (s *Created) ExchangeName() string {
	return "Products:Created"
}
