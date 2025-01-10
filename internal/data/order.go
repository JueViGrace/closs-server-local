package data

import (
	"context"

	"github.com/JueViGrace/clo-backend/internal/db"
	"github.com/JueViGrace/clo-backend/internal/types"
)

type OrderStore interface {
	GetOrders() (orders []types.OrderResponse, err error)
	GetOrdersWithLines() (orders []types.OrderWithLinesResponse, err error)
	GetOrderByCode(code string) (order *types.OrderResponse, err error)
	GetOrderByCodeWithLines(code string) (order *types.OrderWithLinesResponse, err error)
	GetAllOrdersByManager(code string) (orders []types.OrderResponse, err error)
	GetOrdersByManager(code string) (orders []types.OrderResponse, err error)
	GetAllOrdersBySalesman(code string) (orders []types.OrderResponse, err error)
	GetOrdersBySalesman(code string) (orders []types.OrderResponse, err error)
	GetAllOrdersByCustomer(code string) (orders []types.OrderResponse, err error)
	GetOrdersByCustomer(code string) (orders []types.OrderResponse, err error)
	CreateOrder(r *types.CreateOrderRequest) (order *types.OrderWithLinesResponse, err error)
	UpdateOrder(r *types.UpdateOrderRequest) (order *types.OrderWithLinesResponse, err error)
}

func (s *storage) OrderStore() OrderStore {
	return NewOrderStore(s.ctx, s.queries)
}

type orderStore struct {
	ctx context.Context
	db  *db.Queries
}

func NewOrderStore(ctx context.Context, db *db.Queries) OrderStore {
	return &orderStore{
		ctx: ctx,
		db:  db,
	}
}

func (s *orderStore) GetOrders() ([]types.OrderResponse, error) {
	orders := make([]types.OrderResponse, 0)

	dbOrders, err := s.db.GetOrders(s.ctx)
	if err != nil {
		return nil, err
	}

	for _, dbOrder := range dbOrders {
		orders = append(orders, *types.DbOrderToOrder(&dbOrder))
	}

	return orders, nil
}

func (s *orderStore) GetOrdersWithLines() ([]types.OrderWithLinesResponse, error) {
	dbOrders, err := s.db.GetOrdersWithLines(s.ctx)
	if err != nil {
		return nil, err
	}

	return types.GroupOrderWithLinesRow(dbOrders), nil
}

func (s *orderStore) GetOrderByCode(code string) (*types.OrderResponse, error) {
	dbOrder, err := s.db.GetOrderByCode(s.ctx, code)
	if err != nil {
		return nil, err
	}

	return types.DbOrderToOrder(&dbOrder), nil
}

func (s *orderStore) GetOrderByCodeWithLines(code string) (*types.OrderWithLinesResponse, error) {
	dbOrders, err := s.db.GetOrderWithLinesByCode(s.ctx, code)
	if err != nil {
		return nil, err
	}

	return types.GroupOrderWithLinesByCodeRow(dbOrders), nil
}

func (s *orderStore) GetAllOrdersByManager(code string) ([]types.OrderResponse, error) {
	orders := make([]types.OrderResponse, 0)

	dbOrders, err := s.db.GetAllOrdersByManager(s.ctx, code)
	if err != nil {
		return nil, err
	}

	for _, dbOrder := range dbOrders {
		orders = append(orders, *types.DbOrderToOrder(&dbOrder))
	}

	return orders, nil
}

func (s *orderStore) GetOrdersByManager(code string) ([]types.OrderResponse, error) {
	orders := make([]types.OrderResponse, 0)

	dbOrders, err := s.db.GetOrdersByManager(s.ctx, code)
	if err != nil {
		return nil, err
	}

	for _, dbOrder := range dbOrders {
		orders = append(orders, *types.DbOrderToOrder(&dbOrder))
	}

	return orders, nil
}

func (s *orderStore) GetAllOrdersBySalesman(code string) ([]types.OrderResponse, error) {
	orders := make([]types.OrderResponse, 0)

	dbOrders, err := s.db.GetAllOrdersBySalesman(s.ctx, code)
	if err != nil {
		return nil, err
	}

	for _, dbOrder := range dbOrders {
		orders = append(orders, *types.DbOrderToOrder(&dbOrder))
	}

	return orders, nil
}

func (s *orderStore) GetOrdersBySalesman(code string) ([]types.OrderResponse, error) {
	orders := make([]types.OrderResponse, 0)

	dbOrders, err := s.db.GetOrdersBySalesman(s.ctx, code)
	if err != nil {
		return nil, err
	}

	for _, dbOrder := range dbOrders {
		orders = append(orders, *types.DbOrderToOrder(&dbOrder))
	}

	return orders, nil
}

func (s *orderStore) GetAllOrdersByCustomer(code string) ([]types.OrderResponse, error) {
	orders := make([]types.OrderResponse, 0)

	dbOrders, err := s.db.GetAllOrdersByCustomer(s.ctx, code)
	if err != nil {
		return nil, err
	}

	for _, dbOrder := range dbOrders {
		orders = append(orders, *types.DbOrderToOrder(&dbOrder))
	}

	return orders, nil
}

func (s *orderStore) GetOrdersByCustomer(code string) ([]types.OrderResponse, error) {
	orders := make([]types.OrderResponse, 0)

	dbOrders, err := s.db.GetOrdersByCustomer(s.ctx, code)
	if err != nil {
		return nil, err
	}

	for _, dbOrder := range dbOrders {
		orders = append(orders, *types.DbOrderToOrder(&dbOrder))
	}

	return orders, nil
}

func (s *orderStore) CreateOrder(r *types.CreateOrderRequest) (*types.OrderWithLinesResponse, error) {
	lines := make([]types.OrderLineResponse, 0)

	dbOrder, err := s.db.CreateOrder(s.ctx, *types.CreateOrderToDb(r))
	if err != nil {
		return nil, err
	}

	head := types.DbOrderToOrder(&dbOrder)

	for _, line := range r.Lines {
		dbLine, err := s.db.CreateOrderLine(s.ctx, *types.CreateOrderLineToDb(&line))
		if err != nil {
			continue
		}

		lines = append(lines, *types.DbOrderLineToOrderLine(&dbLine))
	}

	return types.MapToOrderWithLines(head, lines), nil
}

func (s *orderStore) UpdateOrder(r *types.UpdateOrderRequest) (*types.OrderWithLinesResponse, error) {
	lines := make([]types.OrderLineResponse, 0)

	dbOrder, err := s.db.UpdateOrder(s.ctx, *types.UpdateOrderToDb(r))
	if err != nil {
		return nil, err
	}

	head := types.DbOrderToOrder(&dbOrder)

	for _, line := range r.Lines {
		dbLine, err := s.db.UpdateOrderLine(s.ctx, *types.UpdateOrderLineToDb(&line))
		if err != nil {
			continue
		}

		lines = append(lines, *types.DbOrderLineToOrderLine(&dbLine))
	}

	return types.MapToOrderWithLines(head, lines), nil
}
