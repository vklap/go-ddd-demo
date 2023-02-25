package order

import (
	"github.com/vklap/go-ddd-demo/without_ddd/internal/order/status"
	"time"
)

type Order struct {
	ID         string
	Items      []*Item
	CustomerID string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Item struct {
	ProductID          string
	ProductPrice       int
	Quantity           int
	DiscountPercentage int
	Status             string
}

func NewOrder(cart *Cart) *Order {
	now := time.Now().UTC()
	items := make([]*Item, 0, len(cart.Items))
	for _, cartItem := range cart.Items {
		item := &Item{
			ProductID:          cartItem.ProductID,
			ProductPrice:       cartItem.ProductPrice,
			Quantity:           cartItem.Quantity,
			DiscountPercentage: cartItem.DiscountPercentage,
			Status:             status.FromCart,
		}
		items = append(items, item)
	}
	order := &Order{
		ID:         cart.ID,
		CustomerID: cart.CustomerID,
		Items:      items,
		CreatedAt:  now,
		UpdatedAt:  now,
	}
	return order
}

func (o *Order) Validate() error {
	return nil
}
