package intent

import (
	"context"
)

//IntentService
type IntentService interface {
	Register(ctx context.Context, req RegisterIntentReq) (RegisterIntentRes, error)
}

type CacheService[T any] interface {
	Get(key string) (any, bool)
	Set(key string, value T, cost int64)
	Clear(T)
}

type RegisterIntentReq struct {
	Name string
}

type RegisterIntentRes struct {
	Status string
}
