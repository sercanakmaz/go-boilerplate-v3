package order

import (
	"fmt"
	"github.com/sercanakmaz/go-boilerplate-v3/anticorruption"
)

type OrderDomainService struct {
}

func (self *OrderDomainService) GetOrdersWithUserInfo(id string) {
	var userHttpClient = anticorruption.UserHttpClient{}

	for true { // loop through orders
		var user = userHttpClient.GetUserById(id)
		fmt.Println(user.UserName)
	}
}
