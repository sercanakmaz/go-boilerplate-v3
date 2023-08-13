package order

import (
	use_case "github.com/sercanakmaz/go-boilerplate-v3/pkg/ddd"
)

type RejectOrderPaymentCommand struct {
	OrderNumber  string                    `json:"orderNumber"`
	RejectReason string                    `json:"rejectReason"`
	identity     *use_case.UseCaseIdentity `json:"_"`
}

func (self *RejectOrderPaymentCommand) GetIdentity() *use_case.UseCaseIdentity {
	return self.identity
}
