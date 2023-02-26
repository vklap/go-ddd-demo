package shipment

type Request struct {
	CustomerID string     `json:"customer_id"`
	OrderID    string     `json:"order_id"`
	Products   []*Product `json:"products"`
}
