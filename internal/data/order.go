package data

import (
	"context"

	database "github.com/JueViGrace/closs-server-local/internal/database/mysql"
	"github.com/JueViGrace/closs-server-local/internal/types"
)

type OrderStore interface {
	GetOrders(a *types.AuthData) (orders []types.OrderWithLinesResponse, err error)
	GetOrderByCode(code string, a *types.AuthData) (order *types.OrderWithLinesResponse, err error)
	UpdateOrderCart(r *types.UpdateOrderCartRequest, a *types.AuthData) (order *types.OrderWithLinesResponse, err error)
	UpdateOrder(r *types.UpdateOrderRequest, a *types.AuthData) (order *types.OrderWithLinesResponse, err error)
}

func (s *storage) OrderStorage() OrderStore {
	return NewOrderStorage(s.ctx, s.myStore.Queries)
}

type orderStore struct {
	ctx context.Context
	db  *database.Queries
}

func NewOrderStorage(ctx context.Context, db *database.Queries) OrderStore {
	return &orderStore{
		ctx: ctx,
		db:  db,
	}
}

func (s *orderStore) GetOrders(a *types.AuthData) ([]types.OrderWithLinesResponse, error) {
	orders := make([]types.OrderWithLinesResponse, 0)

	rows, err := s.db.GetOrdersByUser(s.ctx, a.Username)
	if err != nil {
		return nil, err
	}

	orders, err = types.GroupOrdersByUserRow(rows)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (s *orderStore) GetOrderByCode(code string, a *types.AuthData) (*types.OrderWithLinesResponse, error) {
	order := new(types.OrderWithLinesResponse)

	rows, err := s.db.GetOrderByCode(s.ctx, database.GetOrderByCodeParams{
		Upickup:   a.Username,
		Documento: code,
	})
	if err != nil {
		return nil, err
	}

	order, err = types.GroupOrderByCodeRow(rows)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (s *orderStore) UpdateOrderCart(r *types.UpdateOrderCartRequest, a *types.AuthData) (*types.OrderWithLinesResponse, error) {
	// TODO: do doc update

	return s.GetOrderByCode(r.Documento, a)
}

func (s *orderStore) UpdateOrder(r *types.UpdateOrderRequest, a *types.AuthData) (*types.OrderWithLinesResponse, error) {
	// TODO: do doc update
	return s.GetOrderByCode(r.Documento, a)
}
