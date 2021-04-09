package request

import (
	"github.com/TianqiS/database_for_happyball/configs"
	databaseGrpc "github.com/TianqiS/database_for_happyball/database_grpc"
	"github.com/TianqiS/database_for_happyball/framework"
	proto2 "google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"log"
)

type AddItem struct {
	*framework.BaseEvent
	addItem int32 // 添加的Item的类型
	item interface{}
}

func (addItemReq *AddItem) ToMessage() interface{} {
	return nil
}

func (addItemReq *AddItem) FromMessage(message interface{}) {
	addItemReqPb := message.(*databaseGrpc.AddItem)
	addItemReq.addItem = int32(addItemReqPb.AddItem)
	addItemReq.item = itemDecoder(addItemReqPb.Item, addItemReq.addItem)
}

func itemDecoder(item *anypb.Any, addItem int32) interface{}{
	var err error
	defer func() {
		if err != nil {
			log.Println("itemDecoder error:", err)
		}
	}()
	switch addItem {
	case configs.AddItemPlayer:
		result := &databaseGrpc.AddItem{}
		err = anypb.UnmarshalTo(item, result, proto2.UnmarshalOptions{})
		return result
		break
	default:
	}
	return nil
}
