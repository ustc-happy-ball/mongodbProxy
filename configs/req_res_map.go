package configs

var ReqResMap = map[int32]int32 {
	RequestFindById: ResponseFind,
	RequestFindByKey: ResponseFind,
	RequestDeleteById: ResponseDelete,
	RequestDeleteByKey: ResponseDelete,
	RequestUpdateById: ResponseUpdate,
	RequestUpdateByKey: ResponseUpdate,
	RequestAddRequest: ResponseAdd,
}
