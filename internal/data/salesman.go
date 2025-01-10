package data

import (
	"context"

	"github.com/JueViGrace/clo-backend/internal/db"
	"github.com/JueViGrace/clo-backend/internal/types"
)

type SalesmanStore interface {
	GetSalesmenByManager(code string) (salesmen []types.Salesman, err error)
	GetExistingSalesmenByManager(code string) (salesmen []types.Salesman, err error)
	GetSalesmanByCode(code string) (salesman *types.Salesman, err error)
	GetExistingSalesmanByCode(code string) (salesman *types.Salesman, err error)
	CreateSalesman(r *types.CreateSalesmanRequest) (salesman *types.Salesman, err error)
	UpdateSalesman(r *types.UpdateSalesmanRequest) (salesman *types.Salesman, err error)
}

func (s *storage) SalesmanStore() SalesmanStore {
	return NewSalesmanStore(s.ctx, s.queries)
}

type salesmanStore struct {
	ctx context.Context
	db  *db.Queries
}

func NewSalesmanStore(ctx context.Context, db *db.Queries) SalesmanStore {
	return &salesmanStore{
		ctx: ctx,
		db:  db,
	}
}

func (s *salesmanStore) GetSalesmenByManager(code string) ([]types.Salesman, error) {
	salesmen := make([]types.Salesman, 0)

	dbSalesmen, err := s.db.GetSalesmenByManager(s.ctx, code)
	if err != nil {
		return nil, err
	}

	for _, dbSalesman := range dbSalesmen {
		salesmen = append(salesmen, *types.DbSalesmanByManagerToSalesman(&dbSalesman))
	}

	return salesmen, nil
}

func (s *salesmanStore) GetExistingSalesmenByManager(code string) ([]types.Salesman, error) {
	salesmen := make([]types.Salesman, 0)

	dbSalesmen, err := s.db.GetExistingSalesmenByManager(s.ctx, code)
	if err != nil {
		return nil, err
	}

	for _, dbSalesman := range dbSalesmen {
		salesmen = append(salesmen, *types.DbExistingSalesmanByManagerToSalesman(&dbSalesman))
	}

	return salesmen, nil
}

func (s *salesmanStore) GetExistingSalesmanByCode(code string) (*types.Salesman, error) {
	dbSalesman, err := s.db.GetExistingSalesmanByCode(s.ctx, code)
	if err != nil {
		return nil, err
	}

	return types.DbExistingSalesmanToSalesman(&dbSalesman), nil
}

func (s *salesmanStore) GetSalesmanByCode(code string) (*types.Salesman, error) {
	dbSalesman, err := s.db.GetSalesmanByCode(s.ctx, code)
	if err != nil {
		return nil, err
	}

	return types.DbSalesmanToSalesman(&dbSalesman), nil
}

func (s *salesmanStore) CreateSalesman(r *types.CreateSalesmanRequest) (*types.Salesman, error) {
	dbSalesman, err := s.db.CreateSalesman(s.ctx, *types.CreateSalesmanToDb(r))
	if err != nil {
		return nil, err
	}

	return s.GetSalesmanByCode(dbSalesman.Codigo)
}

func (s *salesmanStore) UpdateSalesman(r *types.UpdateSalesmanRequest) (*types.Salesman, error) {
	dbSalesman, err := s.db.UpdateSalesman(s.ctx, *types.UpdateSalesmanToDb(r))
	if err != nil {
		return nil, err
	}

	return s.GetSalesmanByCode(dbSalesman.Codigo)
}
