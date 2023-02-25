package order

import (
	"context"
	"errors"
	"github.com/vklap/go-ddd-demo/without_ddd/internal/store"
)

type StoreReleaser interface {
	Release(ctx context.Context, products []*store.Product) error
}

type DBSaver interface {
	Save(ctx context.Context, order *Order) error
}

type Service struct {
	storeReleaser StoreReleaser
	dbSaver       DBSaver
}

func NewService() *Service {
	return &Service{
		storeReleaser: store.NewStore(),
		dbSaver:       NewOrderDB(),
	}
}

func (s *Service) CreateOrder(ctx context.Context, cart *Cart) (*Order, error) {
	order := NewOrder(cart)
	if err := order.Validate(); err != nil {
		return nil, err
	}
	if err := s.dbSaver.Save(ctx, order); err != nil {
		return nil, err
	}
	products := make([]*store.Product, 0, len(order.Items))
	for _, item := range order.Items {
		p := &store.Product{ID: item.ProductID, Quantity: item.Quantity}
		products = append(products, p)
	}
	if err := s.storeReleaser.Release(ctx, products); err != nil {
		var releaseError *store.ReleaseError
		if errors.As(err, &releaseError) {
			itemsByProductID := map[string]*Item{}
			productsBytID := map[string]*store.Product{}
		} else {
			return nil, err
		}
	}

}
