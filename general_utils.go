package general_utils

import (
	"fmt"
	logs "github.com/hellgate75/general_utils/log"
	parsers "github.com/hellgate75/general_utils/parsers"
	logger "github.com/hellgate75/general_utils/parsers/logger"
	"github.com/hellgate75/general_utils/streams"
	"os"
)

const (
	logPath     string = "config"
	logFileName string = "log4go"
)

var (
	logFileExtensions []string = []string{"yaml", "xml", "json"}
)

func findFileInPath(folder string, fileName string, extensions []string) (string, string, error) {
	currentPath := streams.GetCurrentPath()
	separator := os.PathSeparator
	var basePath string = fmt.Sprintf("%s%v%s", currentPath, separator, folder)
	_, err := os.Stat(basePath)
	if err != nil {
		return "", "", err
	}
	var found bool = false
	var filePath string = ""
	var extension string = ""
	var fileErr error
	for i := 0; i < len(extensions) && !found; i++ {
		extension = extensions[i]
		filePath = fmt.Sprintf("%s%v%s.%s", basePath, separator, fileName, extension)
		_, fileErr = os.Stat(filePath)
		if fileErr == nil {
			found = true
		} else {
			filePath = ""
			extension = ""
		}
	}
	return filePath, extension, fileErr
}

// Initialized, register and collect Logger from stadards path.
//
// Returns:
//   *logs.Logger Pointer to Initialized system logger
//   error Any suitable error risen during code execution
func InitLogger() (*logs.Logger, error) {
	filePath, extension, fErr := findFileInPath(logPath, logFileName, logFileExtensions)
	if fErr != nil {
		return nil, fErr
	}
	encoding, eErr := parsers.EncodingFromString(extension)
	if eErr != nil {
		return nil, eErr
	}
	conf, err := logger.LoadLoggerConfigFromFile(encoding, filePath)

	if err != nil {
		return nil, err
	}
	err = logs.InitStaticFromConfig(*conf)
	if err != nil {
		return nil, err
	}
	return logs.GetLoggerByRef(), nil
}

// Initialized, register and collect Logger from stadards path.
//
// Parameters:
//   filePath (string) Absolute file path
//
// Returns:
//   *logs.Logger Pointer to Initialized system logger
//   error Any suitable error risen during code execution
func InitLogger(filePath string) (*logs.Logger, error) {
	//TODO Complete!!
	filePath, extension, fErr := findFileInPath(logPath, logFileName, logFileExtensions)
	if fErr != nil {
		return nil, fErr
	}
	encoding, eErr := parsers.EncodingFromString(extension)
	if eErr != nil {
		return nil, eErr
	}
	conf, err := logger.LoadLoggerConfigFromFile(encoding, filePath)

	if err != nil {
		return nil, err
	}
	err = logs.InitStaticFromConfig(*conf)
	if err != nil {
		return nil, err
	}
	return logs.GetLoggerByRef(), nil
}
