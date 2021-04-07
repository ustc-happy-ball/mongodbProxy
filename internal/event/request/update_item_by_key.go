package request

import "github.com/TianqiS/database_for_happyball/framework"

type UpdateItemByKey struct {
	*framework.BaseEvent
	updateItem int // 更新的item的种类
	key string
	value interface{}
	items map[string]interface{}
}

func (updateItemByKey *UpdateItemByKey) ToMessage() interface{} {
	return nil
}

func (updateItemByKey *UpdateItemByKey) FromMessage(interface{}) {

}
