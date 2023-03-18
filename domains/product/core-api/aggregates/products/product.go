package products

import (
	"go-boilerplate-v3/pkg/ddd"
	"go-boilerplate-v3/pkg/mongo"
)

type Product struct {
	Sku          string           `json:"sku"`
	Name         string           `json:"Name"`
	Stock        int              `json:"Stock"`
	CategoryId   int              `json:"CategoryId"`
	domainEvents []ddd.IBaseEvent `json:"domain_events"`
	ddd.IAggregateRoot
	mongo.IEntity
}

func (u *Product) ClearDomainEvents() {
	u.domainEvents = nil
}

func (u *Product) GetDomainEvents() []ddd.IBaseEvent {
	return u.domainEvents
}

func (u *Product) AddEvent(event ddd.IBaseEvent) {
	u.domainEvents = append(u.domainEvents, event)
}

func NewProduct(sku, name string, stock, categoryId int) *Product {
	return &Product{
		Sku:        sku,
		Name:       name,
		Stock:      stock,
		CategoryId: categoryId,
	}
}

func (u *Product) IncreaseStock(amount int) {
	u.Stock += amount
}

func (u *Product) DecreaseStock(amount int) {
	u.Stock -= amount
}
