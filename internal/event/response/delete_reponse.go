package response

import "github.com/TianqiS/database_for_happyball/framework"

type DeleteResponse struct {
	*framework.BaseEvent
	responseStatus int32
	error string
}

func (deleteResponse *DeleteResponse) ToMessage() interface{} {
	return nil
}

func (deleteResponse *DeleteResponse) FromMessage(interface{}) {

}