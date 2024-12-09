package dgraph

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/dgraph-io/dgo/v200/protos/api"
	"log"
)

func (is *IntentService) query(ctx context.Context, query string, vars map[string]string, result any) error {
	resp, err := is.NewTxn().QueryWithVars(ctx, query, vars)
	if err != nil {
		log.Fatal("while querying: ", err)
	}
	return json.Unmarshal(resp.Json, result)
}
func upsert[T any](ctx context.Context, is *IntentService, query string, mu ...*api.Mutation) ([]T, error) {
	request := &api.Request{Query: query, Mutations: mu, CommitNow: true}
	results, err := is.NewTxn().Do(context.Background(), request)
	r := &struct {
		Data []T `json:"q"`
	}{}
	if results == nil {
		var s []T
		return s, err
	}
	err = json.Unmarshal(results.Json, r)
	return r.Data, err
}
func (is *IntentService) mutate(ctx context.Context, data any) (string, error) {
	mu := &api.Mutation{CommitNow: true}
	buf := &bytes.Buffer{}
	err := json.NewEncoder(buf).Encode(data)
	if err != nil {
		return "", err
	}
	mu.SetJson = buf.Bytes()
	response, err := is.NewTxn().Mutate(ctx, mu)
	if err != nil {
		return "", err
	}
	return response.Uids["id"], nil
}
