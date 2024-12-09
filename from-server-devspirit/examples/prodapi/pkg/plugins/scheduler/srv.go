package scheduler

import (
	"context"
)

//SchedulerService
type SchedulerService interface {
	Register(ctx context.Context, req RegisterSchedulerReq) (RegisterSchedulerRes, error)
}

type CacheService[T any] interface {
	Get(key string) (any, bool)
	Set(key string, value T, cost int64)
	Clear(T)
}

type RegisterSchedulerReq struct {
	Time string
}

type RegisterSchedulerRes struct {
	Status string
}
