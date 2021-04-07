package response

import "github.com/TianqiS/database_for_happyball/framework"

type UpdateResponse struct {
	*framework.BaseEvent
}

func (updateResponse *UpdateResponse) ToMessage() interface{} {
	return nil
}

func (updateResponse *UpdateResponse) FromMessage(interface{}) {

}