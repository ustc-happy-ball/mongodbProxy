package response

import "github.com/TianqiS/database_for_happyball/framework"

type FindResponse struct {
	*framework.BaseEvent
	responseStatus int32
	results interface{}
	error string
}

func (findResponse *FindResponse) ToMessage() interface{} {
	return nil
}

func (findResponse *FindResponse) FromMessage(interface{}) {

}
