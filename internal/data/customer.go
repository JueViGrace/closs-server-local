package data

import (
	"context"

	"github.com/JueViGrace/clo-backend/internal/db"
	"github.com/JueViGrace/clo-backend/internal/types"
)

type CustomerStore interface {
	GetCustomers() (customers []types.CustomerResponse, err error)
	GetCustomerByCode(code string) (customer *types.CustomerResponse, err error)
	GetCustomersByManager(code string) (customers []types.CustomerResponse, err error)
	GetCustomersBySalesman(code string) (customers []types.CustomerResponse, err error)
	CreateCustomer(r *types.CreateCustomerRequest) (customer *types.CustomerResponse, err error)
	UpdateCustomer(r *types.UpdateCustomerRequest) (customer *types.CustomerResponse, err error)
}

func (s *storage) CustomerStore() CustomerStore {
	return NewCustomerStore(s.ctx, s.queries)
}

type customerStore struct {
	ctx context.Context
	db  *db.Queries
}

func NewCustomerStore(ctx context.Context, db *db.Queries) CustomerStore {
	return &customerStore{
		ctx: ctx,
		db:  db,
	}
}

func (s *customerStore) GetCustomers() ([]types.CustomerResponse, error) {
	customers := make([]types.CustomerResponse, 0)

	dbCustomers, err := s.db.GetCustomers(s.ctx)
	if err != nil {
		return nil, err
	}

	for _, dbCustomer := range dbCustomers {
		customers = append(customers, *types.DbCustomerToCustomer(&dbCustomer))
	}

	return customers, nil
}

func (s *customerStore) GetCustomerByCode(code string) (*types.CustomerResponse, error) {
	dbCustomer, err := s.db.GetCustomerByCode(s.ctx, code)
	if err != nil {
		return nil, err
	}

	return types.DbCustomerToCustomer(&dbCustomer), nil
}

func (s *customerStore) GetCustomersByManager(code string) ([]types.CustomerResponse, error) {
	customers := make([]types.CustomerResponse, 0)

	dbCustomers, err := s.db.GetCustomersByManager(s.ctx, code)
	if err != nil {
		return nil, err
	}

	for _, dbCustomer := range dbCustomers {
		customers = append(customers, *types.DbCustomerToCustomer(&dbCustomer))
	}

	return customers, nil
}

func (s *customerStore) GetCustomersBySalesman(code string) ([]types.CustomerResponse, error) {
	customers := make([]types.CustomerResponse, 0)

	dbCustomers, err := s.db.GetCustomersBySalesman(s.ctx, code)
	if err != nil {
		return nil, err
	}

	for _, dbCustomer := range dbCustomers {
		customers = append(customers, *types.DbCustomerToCustomer(&dbCustomer))
	}

	return customers, nil
}

func (s *customerStore) CreateCustomer(r *types.CreateCustomerRequest) (*types.CustomerResponse, error) {
	dbCustomer, err := s.db.CreateCustomer(s.ctx, *types.CreateCustomerToDb(r))
	if err != nil {
		return nil, err
	}

	return types.DbCustomerToCustomer(&dbCustomer), nil
}

func (s *customerStore) UpdateCustomer(r *types.UpdateCustomerRequest) (*types.CustomerResponse, error) {
	dbCustomer, err := s.db.UpdateCustomer(s.ctx, *types.UpdateCustomerToDb(r))
	if err != nil {
		return nil, err
	}

	return types.DbCustomerToCustomer(&dbCustomer), nil
}
