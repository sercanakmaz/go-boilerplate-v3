package products

import (
	"github.com/sercanakmaz/go-boilerplate-v3/pkg/ddd"
)

type StockUpdated struct {
	Sku            string `json:"sku"`
	Stock          int    `json:"stock"`
	ddd.IBaseEvent `json:"-"`
}

func (s *StockUpdated) EventName() string {
	return "Product:StockUpdated"
}
