package ristretto

import (
	"github.com/dgraph-io/ristretto"
)

type CacheService[T any] struct{ cache *ristretto.Cache }

func NewCacheService[T any]() *CacheService[T] {
	cache, err := ristretto.NewCache(&ristretto.Config{NumCounters: 1e7, MaxCost: 1 << 30, BufferItems: 64})
	if err != nil {
		panic(err)
	}
	return &CacheService[T]{cache}
}
func (cs CacheService[T]) Get(key string) (any, bool) {
	t, found := cs.cache.Get(key)
	return t, found
}
func (cs CacheService[T]) Set(key string, value T, cost int64) {
	cs.cache.Set(key, value, cost)
	cs.cache.Wait()
}
func (cs CacheService[T]) Clear(data T) {
}
