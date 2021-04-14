package server

import (
	"context"
	"github.com/TianqiS/database_for_happyball/database_grpc"
	"github.com/TianqiS/database_for_happyball/framework"
	"github.com/TianqiS/database_for_happyball/internal/event"
	"github.com/pkg/errors"
	"time"
)

type RpcServer struct {
	databaseGrpc.UnimplementedDbServiceServer
}

func GetServer() *RpcServer {
	return &RpcServer{}
}

var (
	timeoutErr = errors.New("grpc handle request time out!")
)

func (s *RpcServer) SendRequest(ctx context.Context, in *databaseGrpc.DbMessage) (*databaseGrpc.DbMessage, error) {
	var timeout <-chan time.Time
	timeout = time.After(500 * time.Millisecond)
	dbMessageEvent := &event.DbMessage{
		ResResult: make(chan *event.DbMessage),
		Error: make(chan error),
	}
	dbMessageEvent.FromMessage(in)
	go framework.BaseDispatcher.FireEvent(dbMessageEvent)
	select {
	case res := <- dbMessageEvent.ResResult:
		return res.ToMessage().(*databaseGrpc.DbMessage), nil
	case err := <- dbMessageEvent.Error:
		return nil, err
	case <-timeout:
		return nil, errors.WithStack(timeoutErr)
	}
}

func handlerRequest(req *event.DbMessage) {
}