package data

import (
	"context"

	"github.com/JueViGrace/clo-backend/internal/db"
	"github.com/JueViGrace/clo-backend/internal/types"
)

type StatisticStore interface {
	GetStatistics() ([]types.SalesmanStatistic, error)
	GetStatisticsByManager(code string) ([]types.SalesmanStatistic, error)
	GetStatisticBySalesman(code string) (*types.SalesmanStatistic, error)
	CreateStatistic(r *types.CreateStatisticRequest) (*types.SalesmanStatistic, error)
	UpdateStatistic(r *types.UpdateStatisticRequest) (*types.SalesmanStatistic, error)
}

func (s *storage) StatisticStore() StatisticStore {
	return NewStatisticStore(s.ctx, s.queries)
}

type statisticStore struct {
	ctx context.Context
	db  *db.Queries
}

func NewStatisticStore(ctx context.Context, db *db.Queries) StatisticStore {
	return &statisticStore{
		ctx: ctx,
		db:  db,
	}
}

func (s *statisticStore) GetStatistics() ([]types.SalesmanStatistic, error) {
	statistics := make([]types.SalesmanStatistic, 0)

	dbStatistics, err := s.db.GetStatistics(s.ctx)
	if err != nil {
		return nil, err
	}

	for _, dbStatistic := range dbStatistics {
		statistics = append(statistics, *types.DbStatisticToStatistc(&dbStatistic))
	}

	return statistics, nil
}

func (s *statisticStore) GetStatisticsByManager(code string) ([]types.SalesmanStatistic, error) {
	statistics := make([]types.SalesmanStatistic, 0)

	dbStatistics, err := s.db.GetStatisticsByManager(s.ctx, code)
	if err != nil {
		return nil, err
	}

	for _, dbStatistic := range dbStatistics {
		statistics = append(statistics, *types.DbStatisticToStatistc(&dbStatistic))
	}

	return statistics, nil
}

func (s *statisticStore) GetStatisticBySalesman(code string) (*types.SalesmanStatistic, error) {
	dbStatistic, err := s.db.GetStatisticBySalesman(s.ctx, code)
	if err != nil {
		return nil, err
	}

	return types.DbStatisticToStatistc(&dbStatistic), nil
}

func (s *statisticStore) CreateStatistic(r *types.CreateStatisticRequest) (*types.SalesmanStatistic, error) {
	dbStatistic, err := s.db.CreateStatistic(s.ctx, *types.CreateStatisticToDb(r))
	if err != nil {
		return nil, err
	}

	return types.DbStatisticToStatistc(&dbStatistic), nil
}

func (s *statisticStore) UpdateStatistic(r *types.UpdateStatisticRequest) (*types.SalesmanStatistic, error) {
	dbStatistic, err := s.db.UpdateStatistic(s.ctx, *types.UpdateStatisticToDb(r))
	if err != nil {
		return nil, err
	}

	return types.DbStatisticToStatistc(&dbStatistic), nil
}
