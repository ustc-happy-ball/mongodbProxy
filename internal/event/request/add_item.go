package request

import "github.com/TianqiS/database_for_happyball/framework"

type AddItem struct {
	*framework.BaseEvent
	addItem int // 添加的Item的类型
	item interface{}
}

func (addItem *AddItem) ToMessage() interface{} {
	return nil
}

func (addItem *AddItem) FromMessage(interface{}) {

}
