package order

import (
	"context"
	"github.com/vklap/go-ddd-demo/without_ddd/internal/store"
)

type StoreReleaser interface {
	Release(products []*store.Product) error
}

type DBSaver interface {
	Save(ctx context.Context, order *Order) error
}
