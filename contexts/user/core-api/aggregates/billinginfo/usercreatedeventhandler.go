package billinginfo

import (
	"github.com/sercanakmaz/go-boilerplate-v3/events/user/users"
)

// listen user created event in memory and update your datasource

type UserCreatedEventHandler struct {
}

func (self *UserCreatedEventHandler) Consume(userCreated users.UserCreated) {
	var service = BillingDomainService{}
	service.CreateBilling(userCreated)
}
