package net

import (
	"github.com/hellgate75/general_utils/log"
)

var logger log.Logger

func InitLogger() {
	currentLogger, err := log.New("net")
	if err != nil {
		panic(err.Error())
	}
	logger = currentLogger
}
