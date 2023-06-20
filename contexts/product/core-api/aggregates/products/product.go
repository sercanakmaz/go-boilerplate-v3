package products

import (
	"github.com/sercanakmaz/go-boilerplate-v3/events/product/products"
	"github.com/sercanakmaz/go-boilerplate-v3/models/shared"
	"github.com/sercanakmaz/go-boilerplate-v3/pkg/ddd"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	Id           primitive.ObjectID `json:"id" bson:"_id"`
	Sku          string             `json:"sku" bson:"Sku"`
	Name         string             `json:"name" bson:"Name"`
	Stock        int                `json:"stock" bson:"Stock"`
	Price        shared.Money       `json:"price" bson:"Price"`
	FinalPrice   shared.Money       `json:"finalPrice" bson:"FinalPrice"`
	Vat          float64            `json:"vat" bson:"vat"`
	CategoryId   int                `json:"categoryId" bson:"CategoryId"`
	domainEvents []ddd.IBaseEvent
}

func (u *Product) ClearDomainEvents() {
	u.domainEvents = nil
}

func (u *Product) GetDomainEvents() []ddd.IBaseEvent {
	return u.domainEvents
}

func (u *Product) RaiseEvent(event ddd.IBaseEvent) {
	u.domainEvents = append(u.domainEvents, event)
}

func NewProduct(sku, name string, stock int, price shared.Money, categoryId int) *Product {

	var product = &Product{
		Id:         primitive.NewObjectID(),
		Sku:        sku,
		Name:       name,
		Stock:      stock,
		Price:      price,
		CategoryId: categoryId,
	}

	product.calculateFinalPrice()

	product.RaiseEvent(&products.Created{
		Id:         product.Id.Hex(),
		Sku:        product.Sku,
		Name:       product.Name,
		Stock:      product.Stock,
		Price:      product.Price,
		FinalPrice: product.FinalPrice,
		CategoryId: product.CategoryId,
	})

	return product
}

func (u *Product) calculateFinalPrice() {
	u.FinalPrice = shared.Money{
		Value:        u.Vat * u.Price.Value,
		CurrencyCode: u.Price.CurrencyCode,
	}
}

func (u *Product) IncreaseStock(amount int) {
	u.Stock += amount

	u.RaiseEvent(&products.StockIncreased{
		Id:    u.Id.Hex(),
		Stock: u.Stock,
	})
}

func (u *Product) DecreaseStock(amount int) {
	u.Stock -= amount

	u.RaiseEvent(&products.StockDecreased{
		Id:    u.Id.Hex(),
		Stock: u.Stock,
	})
}
