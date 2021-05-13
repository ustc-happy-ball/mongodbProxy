package configs

var  (
	// mongodb
	MongoPoolSize uint64 = 100
	DBName string = "happyball"
	MongoURI string
	MongoURIForTest = "mongodb://localhost:27017"

	// tcp port
	TcpPort string = ":8890"

	// log file
	LogFilePathProduction string = "./log/logs_production.txt"
	LogFilePathDevelopment string = "./log/logs_development.txt"
	// log level
	LogLevel string
	LogToFile bool
)
