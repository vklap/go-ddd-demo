package shipment

type Product struct {
	ID       string `json:"id"`
	Quantity int    `json:"quantity"`
	Price    int    `json:"price"`
}
