package products

import (
	"github.com/sercanakmaz/go-boilerplate-v3/models/shared"
	"github.com/sercanakmaz/go-boilerplate-v3/pkg/ddd"
)

type Created struct {
	Id             string       `json:"id"`
	Sku            string       `json:"sku"`
	Name           string       `json:"name"`
	Stock          int          `json:"stock"`
	CompanyId      int          `json:"company_id"`
	Price          shared.Money `json:"price"`
	FinalPrice     shared.Money `json:"finalPrice"`
	CategoryId     int          `json:"categoryId"`
	ddd.IBaseEvent `json:"-"`
}

func (s *Created) EventName() string {
	return "Product:Created"
}
