package recomodelcli

import (
	"context"
	"fmt"
)

// RecoModelCli represents ...
type RecoModelCli struct {
	KubeConfig string
	ctx        context.Context // background context
	cancel     func()          // cancel background context
}

// New
func New() *RecoModelCli {
	ctx, cancel := context.WithCancel(context.Background())
	return &RecoModelCli{
		ctx:    ctx,
		cancel: cancel,
	}
}

func (rm *RecoModelCli) Open() error {
	return nil
}

// RecoModelCliService is a wrapper around the docker client
type RecoModelCliService struct {
	*RecoModelCli
}

func (rm *RecoModelCli) Close() error {
	return nil
}

func NewRecoModelCliService(rm *RecoModelCli) *RecoModelCliService {
	return &RecoModelCliService{RecoModelCli: rm}
}

func (cli *RecoModelCliService) Recommendation(behavior string) error {
	fmt.Println("Gatting a recommendation for how to respond to this behavior...")
	return nil
}
