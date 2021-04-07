package request

import "github.com/TianqiS/database_for_happyball/framework"

type FindItemByKey struct {
	*framework.BaseEvent
}

func (findItemByKey *FindItemByKey) ToMessage() interface{} {
	return nil
}

func (findItemByKey *FindItemByKey) FromMessage(interface{}) {

}
