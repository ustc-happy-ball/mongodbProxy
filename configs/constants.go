package configs

const (
	// mongodb
	MongoUri string = "mongodb://localhost:27017"
	MongoPoolSize uint64 = 100
	MongoDatabase string = "happyball"

	// tcp port
	TcpPort string = ":50051"

	// enum类型
	ResponseStatusSuccess int32 = 0
	ResponseStatusFail int32 = 1

	AddItemPlayer int32 = 0
)