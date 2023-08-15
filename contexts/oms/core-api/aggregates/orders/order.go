package orders

import (
	"github.com/sercanakmaz/go-boilerplate-v3/events/oms/orders"
	"github.com/sercanakmaz/go-boilerplate-v3/models/shared"
	"github.com/sercanakmaz/go-boilerplate-v3/pkg/ddd"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	Id                  primitive.ObjectID `json:"id" bson:"_id"`
	OrderNumber         string             `json:"orderNumber" bson:"OrderNumber"`
	Price               shared.Money       `json:"price" bson:"Price"`
	FinalPrice          shared.Money       `json:"finalPrice" bson:"FinalPrice"`
	PaymentStatus       string             `json:"paymentStatus" bson:"PaymentStatus"`
	PaymentRejectReason string             `json:"paymentRejectReason" bson:"PaymentRejectReason"`
	domainEvents        []ddd.IBaseEvent
}

func (u *Order) ClearDomainEvents() {
	u.domainEvents = nil
}

func (u *Order) GetDomainEvents() []ddd.IBaseEvent {
	return u.domainEvents
}

func (u *Order) RaiseEvent(event ddd.IBaseEvent) {
	u.domainEvents = append(u.domainEvents, event)
}

func NewOrder(orderNumber string, price shared.Money) *Order {

	var order = &Order{
		Id:          primitive.NewObjectID(),
		OrderNumber: orderNumber,
		Price:       price,
	}

	order.RaiseEvent(&orders.Created{
		Id:          order.Id.Hex(),
		OrderNumber: orderNumber,
		Price:       order.Price,
		FinalPrice:  order.FinalPrice,
	})

	return order
}
