package rest

import (
	"github.com/hellgate75/general_utils/log"
)

var logger log.Logger

func InitLogger() {
	currentLogger, err := log.New("net/rest")
	if err != nil {
		panic(err.Error())
	}
	logger = currentLogger
}
