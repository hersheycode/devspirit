package ristretto

import (
	"github.com/dgraph-io/ristretto"
)

// CacheService represents a service for caching nodes.
type CacheService[T any] struct {
	cache *ristretto.Cache
}

// NewCacheService returns a new instance of CacheService.
func NewCacheService[T any]() *CacheService[T] {
	cache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,     // number of keys to track frequency of (10M).
		MaxCost:     1 << 30, // maximum cost of cache (1GB).
		BufferItems: 64,      // number of keys per Get buffer.
	})
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
	// set a value with a cost of 1
	cs.cache.Set(key, value, cost)

	// wait for value to pass through buffers
	cs.cache.Wait()
}

func (cs CacheService[T]) Clear(data T) {

}
