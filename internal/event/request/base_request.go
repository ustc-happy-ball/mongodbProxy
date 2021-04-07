package request

type BaseRequest struct {
	*UpdateItemById
	*UpdateItemByKey
	*AddItem
	*FindItemById
	*FindItemByKey
	*DeleteItemById
	*DeleteItemByKey
}