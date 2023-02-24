package order

type Cart struct {
	ID         string      `json:"id"`
	CustomerID string      `json:"customer_id"`
	Items      []*CartItem `json:"items"`
}

type CartItem struct {
	ProductID          string `json:"product_id"`
	ProductPrice       int    `json:"product_price"`
	Quantity           int    `json:"quantity"`
	DiscountPercentage int    `json:"discount_percentage"`
}
