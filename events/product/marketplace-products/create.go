package marketplace_products

import (
	"github.com/sercanakmaz/go-boilerplate-v3/pkg/ddd"
)

type Create struct {
	Sku           string `json:"sku"`
	CompanyId     int    `json:"company_id"`
	MarketplaceId int    `json:"marketplace_id"`
	ddd.IBaseEvent
}

func (s *Create) EventName() string {
	return "Marketplace:Create"
}
