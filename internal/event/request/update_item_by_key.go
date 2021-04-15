package request

import "github.com/TianqiS/database_for_happyball/framework"

type UpdateItemByKey struct {
	*framework.BaseEvent
	UpdateItem int32 // 更新的item的种类
	Key string
	Value interface{}
	Items map[string]interface{}
}

func (updateItemByKey *UpdateItemByKey) ToMessage() interface{} {
	return nil
}

func (updateItemByKey *UpdateItemByKey) FromMessage(interface{}) {

}
