package data

import (
	"context"
	"database/sql"
	"time"

	"github.com/JueViGrace/clo-backend/internal/db"
	"github.com/JueViGrace/clo-backend/internal/types"
)

type ProductStore interface {
	GetProducts() (products []types.ProductResponse, err error)
	GetExistingProducts() (products []types.ProductResponse, err error)
	GetProductByCode(code string) (product *types.ProductResponse, err error)
	GetExistingProductByCode(code string) (product *types.ProductResponse, err error)
	CreateProduct(r *types.CreateProductRequest) (product *types.ProductResponse, err error)
	UpdateProduct(r *types.UpdateProductRequest) (product *types.ProductResponse, err error)
	SoftDeleteProduct(code string) (err error)
	DeleteProduct(code string) (err error)
}

func (s *storage) ProductStore() ProductStore {
	return NewProductStore(s.ctx, s.queries)
}

type productStore struct {
	ctx context.Context
	db  *db.Queries
}

func NewProductStore(ctx context.Context, db *db.Queries) ProductStore {
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

	for _, dbProduct := range dbProducts {
		products = append(products, *types.DbProductToProduct(&dbProduct))
	}

	return products, nil
}

func (s *productStore) GetExistingProducts() ([]types.ProductResponse, error) {
	products := make([]types.ProductResponse, 0)

	dbProducts, err := s.db.GetExistingProducts(s.ctx)
	if err != nil {
		return nil, err
	}

	for _, dbProduct := range dbProducts {
		products = append(products, *types.DbProductToProduct(&dbProduct))
	}

	return products, nil
}

func (s *productStore) GetProductByCode(code string) (*types.ProductResponse, error) {
	dbProduct, err := s.db.GetProductByCode(s.ctx, code)
	if err != nil {
		return nil, err
	}

	return types.DbProductToProduct(&dbProduct), nil
}

func (s *productStore) GetExistingProductByCode(code string) (*types.ProductResponse, error) {
	dbProduct, err := s.db.GetExistingProductByCode(s.ctx, code)
	if err != nil {
		return nil, err
	}

	return types.DbProductToProduct(&dbProduct), nil
}

func (s *productStore) CreateProduct(r *types.CreateProductRequest) (*types.ProductResponse, error) {
	dbProduct, err := s.db.CreateProduct(s.ctx, *types.CreateProductToDb(r))
	if err != nil {
		return nil, err
	}

	return types.DbProductToProduct(&dbProduct), nil
}

func (s *productStore) UpdateProduct(r *types.UpdateProductRequest) (*types.ProductResponse, error) {
	dbProduct, err := s.db.UpdateProduct(s.ctx, *types.UpdateProductToDb(r))
	if err != nil {
		return nil, err
	}

	return types.DbProductToProduct(&dbProduct), nil
}

func (s *productStore) SoftDeleteProduct(code string) error {
	err := s.db.SoftDeleteProduct(s.ctx, db.SoftDeleteProductParams{
		UpdatedAt: time.Now().String(),
		DeletedAt: sql.NullString{
			String: time.Now().String(),
			Valid:  true,
		},
		Codigo: code,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *productStore) DeleteProduct(code string) error {
	err := s.db.DeleteProduct(s.ctx, code)
	if err != nil {
		return err
	}

	return nil
}
