package shared

type Money struct {
	Value        float64 `json:"value" bson:"Value"`
	CurrencyCode string  `json:"currencyCode" bson:"CurrencyCode"`
}
