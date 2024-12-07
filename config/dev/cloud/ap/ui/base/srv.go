package base

import (
	"context"
	"io"
)

//CPluginService
type CPluginService interface {
	Create(ctx context.Context, rw io.ReadWriter) error
}

type CacheService[T any] interface {
	Get(key string) (any, bool)
	Set(key string, value T, cost int64)
	Clear(T)
}

type CreateReq struct {
	Name    string
	Content []byte
	Opts    *CreateOpts //optional
}

type CreateOpts struct {
	BaseImage     string
	ImageName     string
	ContainerName string
	DstDirName    string
}

type CreateRes struct {
	Status string
}

type ReqReader interface {
	CreateReq
}

type ResWriter interface {
	CreateRes
}

type HandleFunc[r ReqReader, w ResWriter] func(r, *w) error
