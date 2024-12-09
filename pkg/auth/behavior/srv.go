package behavior

import (
	"context"
	"io"
)

//BehaviorService
type BehaviorService interface {
	LogCmd(ctx context.Context, rw io.ReadWriter) error
}

type CacheService[T any] interface {
	Get(key string) (any, bool)
	Set(key string, value T, cost int64)
	Clear(T)
}

type LogCmdReq struct {
	Command string
}

type LogCmdRes struct {
	Status string
}

type ReqReader interface {
	LogCmdReq
}

type ResWriter interface {
	LogCmdRes
}

type HandleFunc[r ReqReader, w ResWriter] func(r, *w) error
