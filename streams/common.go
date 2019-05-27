package streams

import (
	"github.com/hellgate75/general_utils/log"
)

var logger log.Logger = init_logger()

func init_logger() log.Logger {
	logger, _ := log.New("streams")
	return logger
}
