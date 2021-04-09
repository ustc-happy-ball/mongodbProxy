package configs

const (
	// mongodb
	MongoUri string = "mongodb://localhost:27017"
	MongoPoolSize uint64 = 100
	MongoDatabase string = "happyball"

	// tcp port
	TcpPort string = ":50051"

	// enum类型
	responseStatusSuccess int32 = 0
	responseStatusFail int32 = 1
)