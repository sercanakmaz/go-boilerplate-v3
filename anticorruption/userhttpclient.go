package anticorruption

import "github.com/sercanakmaz/go-boilerplate-v3/models/user"

type UserHttpClient struct {
}

func (self *UserHttpClient) GetUserById(id string) user.NewUserModel {

	// make a http call to http://user-api.hepsiglobal.com/v1/users/id/:id
	// deseralize response to NewUserModel

	return user.NewUserModel{}
}
