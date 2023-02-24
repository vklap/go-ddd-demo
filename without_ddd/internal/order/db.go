package order

import (
	"context"
	"fmt"
)

type InMemoryDB struct {
	orders []*Order
}

func (db *InMemoryDB) GetById(ctx context.Context, ID string) (*Order, error) {
	for _, order := range db.orders {
		if order.ID == ID {
			return order, nil
		}
	}
	return nil, fmt.Errorf("order with ID %q does not exist", ID)
}

func (db *InMemoryDB) Save(ctx context.Context, order *Order) error {
	for i, o := range db.orders {
		if o.CartID == order.CartID {
			db.orders[i] = order
			return nil
		}
	}
	db.orders = append(db.orders, order)
	return nil
}
