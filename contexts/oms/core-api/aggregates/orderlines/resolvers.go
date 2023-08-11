package orderlines

import (
	"github.com/sercanakmaz/go-boilerplate-v3/pkg/mongo"
)

func NewOrderLineRepositoryResolve(mongoConfig mongo.Config) IOrderLineRepository {
	return newOrderLineRepository(mongo.NewMongoDb(mongoConfig))
}

func NewOrderLineServiceResolve(mongoConfig mongo.Config) IOrderLineService {
	return NewOrderLineService(NewOrderLineRepositoryResolve(mongoConfig))
}
