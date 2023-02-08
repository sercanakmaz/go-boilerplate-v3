package billinginfo

import "go-boilerplate-v3/events/user"

// listen user created event in memory and update your datasource

type UserCreatedEventHandler struct {
}

func (self *UserCreatedEventHandler) Consume(userCreated user.UserCreated) {
	var service = BillingDomainService{}
	service.CreateBilling(userCreated)
}
