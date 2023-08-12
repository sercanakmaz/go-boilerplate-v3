package orders

type PaymentRejected struct {
	Id                  string `json:"id"`
	OrderNumber         string `json:"orderNumber"`
	PaymentStatus       string `json:"paymentStatus"`
	PaymentRejectReason string `json:"paymentRejectReason"`
}

func (s *PaymentRejected) ExchangeName() string {
	return "Orders:PaymentRejected"
}
