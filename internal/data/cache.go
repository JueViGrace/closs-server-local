package data

import (
	"context"
	"database/sql"
)

type CacheStorage interface {
	SessionStorage() SessionStorage
}

type cacheStorage struct {
	db  *sql.DB
	ctx context.Context
}

func NewCacheStorage() CacheStorage {
	return &cacheStorage{
		ctx: context.Background(),
	}
}
