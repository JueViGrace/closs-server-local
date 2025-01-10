package data

import (
	"context"

	"github.com/JueViGrace/clo-backend/internal/db"
	"github.com/JueViGrace/clo-backend/internal/types"
)

type ConfigStore interface {
	GetConfigs() (configs []types.ConfigResponse, err error)
	GetConfigsByUser(username string) (configs []types.ConfigResponse, err error)
	CreateConfig(r *types.CreateConfigRequest) (config *types.ConfigResponse, err error)
	UpdateConfig(r *types.UpdateConfigRequest) (config *types.ConfigResponse, err error)
	DeleteConfig(r *types.DeleteConfigRequest) (err error)
}

func (s *storage) ConfigStore() ConfigStore {
	return NewConfigStore(s.ctx, s.queries)
}

type configStore struct {
	ctx context.Context
	db  *db.Queries
}

func NewConfigStore(ctx context.Context, db *db.Queries) ConfigStore {
	return &configStore{
		ctx: ctx,
		db:  db,
	}
}

func (s *configStore) GetConfigs() ([]types.ConfigResponse, error) {
	configs := make([]types.ConfigResponse, 0)

	dbConfigs, err := s.db.GetConfigs(s.ctx)
	if err != nil {
		return nil, err
	}

	for _, dbConfig := range dbConfigs {
		configs = append(configs, *types.DbConfigToConfig(&dbConfig))
	}

	return configs, nil
}

func (s *configStore) GetConfigsByUser(username string) ([]types.ConfigResponse, error) {
	configs := make([]types.ConfigResponse, 0)

	dbConfigs, err := s.db.GetConfigsByUsername(s.ctx, username)
	if err != nil {
		return nil, err
	}

	for _, dbConfig := range dbConfigs {
		configs = append(configs, *types.DbConfigToConfig(&dbConfig))
	}

	return configs, nil
}

func (s *configStore) CreateConfig(r *types.CreateConfigRequest) (*types.ConfigResponse, error) {
	dbConfig, err := s.db.CreateConfig(s.ctx, *types.CreateConfigToDb(r))
	if err != nil {
		return nil, err
	}

	return types.DbConfigToConfig(&dbConfig), nil
}

func (s *configStore) UpdateConfig(r *types.UpdateConfigRequest) (*types.ConfigResponse, error) {
	dbConfig, err := s.db.UpdateConfig(s.ctx, *types.UpdateConfigToDb(r))
	if err != nil {
		return nil, err
	}

	return types.DbConfigToConfig(&dbConfig), nil
}

func (s *configStore) DeleteConfig(r *types.DeleteConfigRequest) error {
	err := s.db.DeleteConfig(s.ctx, db.DeleteConfigParams{
		CnfgIdconfig: r.CnfgIdconfig,
		Username:     r.Username,
	})
	if err != nil {
		return err
	}

	return nil
}
