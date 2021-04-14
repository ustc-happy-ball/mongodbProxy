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
	var responseMes chan *databaseGrpc.DbMessage
	var handleErr chan error
	timeout = time.After(500 * time.Millisecond)

	select {
	case res := <- responseMes:
		return res, nil
	case err := <-handleErr:
		return nil, err
	case <-timeout:
		return nil, errors.WithStack(timeoutErr)
	}
}

func handlerRequest(in *databaseGrpc.DbMessage, response chan *databaseGrpc.DbMessage, handleErr chan error) {
	dbMessageEvent := &event.DbMessage{}
	dbMessageEvent.FromMessage(in)
	framework.BaseDispatcher.FireEvent(dbMessageEvent)
}