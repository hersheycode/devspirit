package sms

import (
	"context"
)

//SMSService
type SMSService interface {
	Send(ctx context.Context, req SendReq) (SendRes, error)
}

type CacheService[T any] interface {
	Get(key string) (any, bool)
	Set(key string, value T, cost int64)
	Clear(T)
}

type Message struct {
	Body string
}

type SendReq struct {
	PhoneNum string
	Email    string
	Message
}

type SendRes struct {
	Status string
}
