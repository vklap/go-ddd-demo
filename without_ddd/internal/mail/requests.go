package mail

type OrderNotificationRequest struct {
	OrderID             string
	CustomerID          string
	UnavailableProducts []string
	OrderStatus         string
}
