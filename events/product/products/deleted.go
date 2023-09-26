package products

import (
	"github.com/sercanakmaz/go-boilerplate-v3/pkg/ddd"
)

type Deleted struct {
	Id             string `json:"id"`
	Sku            string `json:"sku"`
	ddd.IBaseEvent `json:"-"`
}

func (s *Deleted) EventName() string {
	return "Product:Deleted"
}
