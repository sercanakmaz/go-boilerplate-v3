package products

type StockIncreased struct {
	Id    string `json:"id"`
	Stock int    `json:"Stock"`
}

func (s *StockIncreased) ExchangeName() string {
	return "Products:StockIncreased"
}
