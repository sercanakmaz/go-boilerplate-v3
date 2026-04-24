package product

type UpdateStockResponse struct {
	Sku   string `json:"sku"`
	Stock int    `json:"stock"`
}
