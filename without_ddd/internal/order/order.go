package order

import (
	"github.com/google/uuid"
	"time"
)

type Order struct {
	ID         string
	Items      []*Item
	CustomerID string
	CartID     string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Item struct {
	ProductID          string
	ProductPrice       int
	Quantity           int
	DiscountPercentage int
}

func New(cart *Cart) *Order {
	now := time.Now().UTC()
	items := make([]*Item, 0, len(cart.Items))
	for _, cartItem := range cart.Items {
		item := &Item{
			ProductID:          cartItem.ProductID,
			ProductPrice:       cartItem.ProductPrice,
			Quantity:           cartItem.Quantity,
			DiscountPercentage: cartItem.DiscountPercentage,
		}
		items = append(items, item)
	}
	order := &Order{
		ID:         uuid.NewString(),
		CustomerID: cart.CustomerID,
		CartID:     cart.ID,
		Items:      items,
		CreatedAt:  now,
		UpdatedAt:  now,
	}
	return order
}
