package intentsys

import (
	"context"
)

//IntentSysService
type IntentSysService interface {
	SetIntent(ctx context.Context, req SetIntentReq) (SetIntentRes, error)
}

type CacheService[T any] interface {
	Get(key string) (any, bool)
	Set(key string, value T, cost int64)
	Clear(T)
}

type SetIntentReq struct {
	*Schedule
	SMSInfo
	Intent
}

type Schedule struct {
	Time string
}

type Intent struct {
	Name string
}

type SMSInfo struct {
	Recipient
	Message
}

type Message struct {
	Body string
}

type Recipient struct {
	PhoneNum string
	Email    string
}

type SetIntentRes struct {
	Status string
}
