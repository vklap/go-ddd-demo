package store

import (
	"context"
)

type InMemoryStore struct {
	productsById map[string]*Product
}

func NewStore(products []*Product) *InMemoryStore {
	store := &InMemoryStore{
		productsById: map[string]*Product{},
	}
	for _, p := range products {
		store.productsById[p.ID] = p
	}
	return store
}

func (s *InMemoryStore) Release(ctx context.Context, products []*Product) error {
	unavailableProducts := make([]*Product, 0)

	for _, p := range products {
		availableProduct, ok := s.productsById[p.ID]
		if !ok || availableProduct.Quantity < p.Quantity {
			unavailableProducts = append(unavailableProducts, p)
			continue
		}
		availableProduct.Quantity -= p.Quantity
	}

	if len(unavailableProducts) > 0 {
		return NewReleaseError("at least one product was not available", unavailableProducts)
	}

	return nil
}
