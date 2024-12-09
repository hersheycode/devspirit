package dgraph

import (
	"bytes"
	"context"
	"encoding/json"
	// "fmt"
	"github.com/dgraph-io/dgo/v200/protos/api"
	"log"
)

func (g *CPluginService) commit(ctx context.Context, data any) error {
	mu := &api.Mutation{CommitNow: true}
	buf := &bytes.Buffer{}
	err := json.NewEncoder(buf).Encode(data)
	if err != nil {
		return err
	}
	mu.SetJson = buf.Bytes()
	_, err = g.NewTxn().Mutate(ctx, mu)
	if err != nil {
		return err
	}
	// fmt.Printf("Database: %+v \n", response)
	return nil
}

func (g *CPluginService) query(ctx context.Context, query string, vars map[string]string, result any) error {
	resp, err := g.NewTxn().QueryWithVars(ctx, query, vars)
	if err != nil {
		log.Fatal("while querying: ", err)
	}
	return json.Unmarshal(resp.Json, result)
}
