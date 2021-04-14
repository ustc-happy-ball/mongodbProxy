package configs

const (
	// mongodb
	MongoUri string = "mongodb://localhost:27017"
	MongoPoolSize uint64 = 100
	MongoDatabase string = "happyball"

	// tcp port
	TcpPort string = ":50051"

	// enum类型
	RequestFindById int32 = 0
	RequestFindByKey int32 = 1
	RequestAddRequest int32 = 2
	RequestDeleteById int32 = 3
	RequestDeleteByKey int32 = 4
	RequestUpdateById int32 = 5
	RequestUpdateByKey int32 = 6

	ResponseFind int32 = 7
	ResponseAdd int32 = 8
	ResponseDelete int32 = 9
	ResponseUpdate int32 = 10

	ResponseStatusSuccess int32 = 0
	ResponseStatusFail int32 = 1

	AddItemPlayer int32 = 0
)