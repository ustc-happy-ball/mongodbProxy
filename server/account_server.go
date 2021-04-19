package server

import (
	"context"
	"github.com/TianqiS/database_for_happyball/database_grpc"
	"github.com/TianqiS/database_for_happyball/db"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RpcServer struct {
	databaseGrpc.UnimplementedAccountServiceServer
}

func GetServer() *RpcServer {
	return &RpcServer{}
}

var (
	timeoutErr = errors.New("grpc handle request time out!")
)

func (*RpcServer) AccountFindByPhone(ctx context.Context, req *databaseGrpc.AccountFindByPhoneRequest) (*databaseGrpc.AccountFindByPhoneResponse, error) {
	phone := req.Phone
	db.GetAccountCollection().FindItemsByKey([]*db.MatchItem{
		{
			Key: "phone",
			MatchVal: phone,
		},
	})
	return nil, nil
}
func (*RpcServer) AccountAdd(ctx context.Context, req *databaseGrpc.AccountAddRequest) (*databaseGrpc.AccountAddResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AccountAdd not implemented")
}
