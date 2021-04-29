package server

import (
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
)

func ErrorHandler(err error) error {
	var encoderedError error
	if mongoWriteException, ok := err.(mongo.WriteException); ok {
		mongoWriteErrs := mongoWriteException.WriteErrors
		writeErr := mongoWriteErrs[0]
		switch writeErr.Code {
		case 11000:
			encoderedError = errors.New("预写入数据与数据库中原有的数据冲突")

			break
		default:
			encoderedError = errors.New("mongodb发生其他错误: " + mongoWriteException.Error())
		}
		return encoderedError
	}
	encoderedError = errors.New("发生其他错误: " + err.Error())
	return encoderedError
}