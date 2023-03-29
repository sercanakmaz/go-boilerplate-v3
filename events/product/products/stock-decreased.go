package products

type StockDecreased struct {
	Id    string `json:"id"`
	Stock int    `json:"Stock"`
}

func (s *StockDecreased) ExchangeName() string {
	return "Products:StockDecreased"
}
