package data

import (
	"context"

	database "github.com/JueViGrace/closs-server-local/internal/database/mysql"
	"github.com/JueViGrace/closs-server-local/internal/types"
)

type ProductStore interface {
	GetProducts() (products []types.ProductResponse, err error)
	GetProduct(code string) (product *types.ProductResponse, err error)
}

func (s *storage) ProductStore() ProductStore {
	return NewProductStore(s.ctx, s.myStore.Queries)
}

type productStore struct {
	ctx context.Context
	db  *database.Queries
}

func NewProductStore(ctx context.Context, db *database.Queries) ProductStore {
	return &productStore{
		ctx: ctx,
		db:  db,
	}
}

func (s *productStore) GetProducts() ([]types.ProductResponse, error) {
	products := make([]types.ProductResponse, 0)

	dbProducts, err := s.db.GetProducts(s.ctx)
	if err != nil {
		return nil, err
	}

	for _, p := range dbProducts {
		products = append(products, *types.DbProductToProduct(&p))
	}

	return products, nil
}

func (s *productStore) GetProduct(code string) (*types.ProductResponse, error) {
	product := new(types.ProductResponse)

	dbProduct, err := s.db.GetProductByCode(s.ctx, code)
	if err != nil {
		return nil, err
	}

	product = types.DbProductToProduct(&dbProduct)

	return product, nil
}
