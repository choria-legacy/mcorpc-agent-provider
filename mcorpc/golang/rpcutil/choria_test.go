package rpcutil

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/choria-io/go-config"
	natsd "github.com/nats-io/nats-server/v2/server"

	"github.com/choria-io/go-choria/choria"
	"github.com/choria-io/go-choria/server"
)

type TestInstance struct {
	ChoriaServer *server.Instance
	NatsServer   *natsd.Server
	Choria       *choria.Framework

	Ctx    context.Context
	cancel func()
}

func StartBrokerAndServer(cfg *config.Config, wg *sync.WaitGroup) (test *TestInstance, err error) {
	test = &TestInstance{}

	test.NatsServer, err = natsd.NewServer(&natsd.Options{Port: -1})
	if err != nil {
		return nil, err
	}
	go test.NatsServer.Start()

	if !test.NatsServer.ReadyForConnections(10 * time.Second) {
		return nil, fmt.Errorf("nats server did not start in 10 seconds")
	}

	cfg.Choria.MiddlewareHosts = []string{test.NatsServer.ClientURL()}

	test.Choria, err = choria.NewWithConfig(cfg)
	if err != nil {
		return nil, err
	}

	test.ChoriaServer, err = server.NewInstance(test.Choria)
	if err != nil {
		return nil, err
	}

	test.Ctx, test.cancel = context.WithCancel(context.Background())
	wg.Add(1)
	test.ChoriaServer.Run(test.Ctx, wg)

	return test, nil
}

func (t *TestInstance) Stop(wg *sync.WaitGroup) {
	t.cancel()
	wg.Wait()
	t.NatsServer.Shutdown()
	t.ChoriaServer = nil
	t.NatsServer = nil
}
