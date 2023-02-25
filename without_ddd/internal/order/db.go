package order

import (
	"context"
	"fmt"
)

type InMemoryOrderDB struct {
	ordersById     map[string]*Order
	ordersByCartId map[string]*Order
}

func NewOrderDB() *InMemoryOrderDB {
	return &InMemoryOrderDB{
		ordersById: map[string]*Order{},
	}
}

func (db *InMemoryOrderDB) GetById(ctx context.Context, ID string) (*Order, error) {
	order, ok := db.ordersById[ID]
	if !ok {
		return nil, fmt.Errorf("order with ID %q does not exist", ID)
	}
	return order, nil
}

func (db *InMemoryOrderDB) Save(ctx context.Context, order *Order) error {
	db.ordersById[order.ID] = order
	return nil
}
