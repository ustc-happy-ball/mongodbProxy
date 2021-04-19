package server

import (
	"github.com/TianqiS/database_for_happyball/database_grpc"
	"github.com/pkg/errors"
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

//func (s *RpcServer) SendRequest(ctx context.Context, in *databaseGrpc.DbMessage) (*databaseGrpc.DbMessage, error) {
//	var timeout <-chan time.Time
//	timeout = time.After(5 * time.Second)
//	dbMessageEvent := &event.DbMessage{
//		ResResult: make(chan *event.DbMessage),
//		Error: make(chan error),
//	}
//	dbMessageEvent.FromMessage(in)
//	go framework.BaseDispatcher.FireEvent(dbMessageEvent)
//	select {
//	case res := <- dbMessageEvent.ResResult:
//		return res.ToMessage().(*databaseGrpc.DbMessage), nil
//	case err := <- dbMessageEvent.Error:
//		errMessage := event.NewErrResMsg(configs.ReqResMap[dbMessageEvent.MessageCode], configs.ResponseStatusUnexpectedError, err.Error())
//		return errMessage.ToMessage().(*databaseGrpc.DbMessage), nil
//	case <-timeout:
//		errMessage := event.NewErrResMsg(configs.ReqResMap[dbMessageEvent.MessageCode], configs.ResponseStatusTimeOut, "超时了")
//		return errMessage.ToMessage().(*databaseGrpc.DbMessage), nil
//	}
//}
