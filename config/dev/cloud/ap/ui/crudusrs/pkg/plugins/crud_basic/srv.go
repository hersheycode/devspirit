package phenomena

import (
	"context"
)

//PhenomenaService
type PhenomenaService interface {
	Create(ctx context.Context, req CreateIntentReq) (CreateIntentRes, error)
	Phenomenon(ctx context.Context, req PhenomenonReq) (PhenomenonRes, error)
	Update(ctx context.Context, req UpdatePhenomenonReq) (UpdatePhenomenonRes, error)
	Delete(ctx context.Context, req DeletePhenomenonReq) (DeletePhenomenonRes, error)
}

type CacheService[T any] interface {
	Get(key string) (any, bool)
	Set(key string, value T, cost int64)
	Clear(T)
}

type CreatePhenomenonReq struct {
	Name string
}

type CreatePhenomenonRes struct {
	Status string
}

type PhenomenonReq struct {
	ID string
}

type PhenomenonRes struct {
	Name string
}

type UpdatePhenomenonReq struct {
	ID string
}

type UpdatePhenomenonRes struct {
	Status string
}

type DeletePhenomenonReq struct {
	ID string
}

type DeletePhenomenonRes struct {
	Status string
}
