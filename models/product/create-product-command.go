package product

type CreateProductCommand struct {
	Sku          string `json:"sku"`
	Name         string `json:"name"`
	InitialStock int    `json:"initialStock"`
	CategoryID   int    `json:"categoryID"`
}
