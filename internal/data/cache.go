package data

import "context"

type CacheStorage interface {
	SessionStorage() SessionStorage
}

type cacheStorage struct {
	ctx context.Context
}

func NewCacheStorage() CacheStorage {
	return &cacheStorage{
		ctx: context.Background(),
	}
}
