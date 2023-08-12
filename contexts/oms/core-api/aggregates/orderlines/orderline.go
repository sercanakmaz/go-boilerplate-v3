package orderlines

import (
	"github.com/sercanakmaz/go-boilerplate-v3/events/oms/orderlines"
	"github.com/sercanakmaz/go-boilerplate-v3/models/shared"
	"github.com/sercanakmaz/go-boilerplate-v3/pkg/ddd"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderLine struct {
	Id           primitive.ObjectID `json:"id" bson:"_id"`
	Sku          string             `json:"sku" bson:"Sku"`
	OrderNumber  string             `json:"orderNumber" bson:"OrderNumber"`
	Price        shared.Money       `json:"price" bson:"Price"`
	Status       string             `json:"status" bson:"Status"`
	CancelReason string             `json:"cancelReason" bson:"CancelReason"`
	domainEvents []ddd.IBaseEvent
}

func (u *OrderLine) ClearDomainEvents() {
	u.domainEvents = nil
}

func (u *OrderLine) GetDomainEvents() []ddd.IBaseEvent {
	return u.domainEvents
}

func (u *OrderLine) RaiseEvent(event ddd.IBaseEvent) {
	u.domainEvents = append(u.domainEvents, event)
}

func NewOrderLine(sku, orderNumber string, price shared.Money) *OrderLine {
	var order = &OrderLine{
		Id:          primitive.NewObjectID(),
		Sku:         sku,
		OrderNumber: orderNumber,
		Price:       price,
	}

	order.RaiseEvent(&orderlines.Created{
		Id:          order.Id.Hex(),
		Sku:         sku,
		OrderNumber: orderNumber,
		Price:       order.Price,
	})

	return order
}

func (u *OrderLine) Cancel(reason string) {
	u.CancelReason = reason
	u.Status = "Cancelled"

	u.RaiseEvent(&orderlines.Cancelled{
		Id:           u.Id.Hex(),
		OrderNumber:  u.OrderNumber,
		Status:       u.Status,
		CancelReason: u.CancelReason,
	})
}
