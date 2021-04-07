package response

import "github.com/TianqiS/database_for_happyball/framework"

type AddResponse struct {
	*framework.BaseEvent
}

func (addResponse *AddResponse) ToMessage() interface{} {
	return nil
}

func (addResponse *AddResponse) FromMessage(interface{}) {

}