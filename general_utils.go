//package general_utils
package main

import (
	"fmt"
	"github.com/hellgate75/general_utils/errors"
	logs "github.com/hellgate75/general_utils/log"
	parsers "github.com/hellgate75/general_utils/parsers"
	logger "github.com/hellgate75/general_utils/parsers/logger"
	"github.com/hellgate75/general_utils/streams"
	"os"
	"strings"
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
	separator := fmt.Sprintf("%c", os.PathSeparator)
	var basePath string = fmt.Sprintf("%s%s%s", currentPath, separator, folder)
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
		filePath = fmt.Sprintf("%s%s%s.%s", basePath, separator, fileName, extension)
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

func decomposeFilePath(filePath string) (string, string, []string, error) {
	if strings.TrimSpace(filePath) == "" {
		return "", "", []string{}, errors.New("general_utils::error : Received in input an Empty String")
	}
	var extension string = "json"
	var path string = ""
	var file string = ""
	separator := fmt.Sprintf("%c", os.PathSeparator)

	fmt.Println("original path: " + filePath)
	fmt.Println("separator: " + separator)

	var idx1 int = strings.LastIndex(filePath, separator)
	var idx2 int = strings.LastIndex(filePath, ".")

	if idx1 < 0 || idx2 < 0 {
		return "", "", []string{}, errors.New("general_utils::error : Invalid path format, please pass an absolute path")
	}

	path = filePath[0:idx1]
	file = filePath[idx1+1 : idx2]
	extension = filePath[idx2+1:]

	return path, file, []string{extension}, nil
}

// Initialized, register and collect Logger from stadards path.
//
// Returns:
//   *logs.Logger Pointer to Initialized system logger
//   error Any suitable error risen during code execution
func InitSimpleLogger(verbosity logs.LogLevel) {
	logs.InitStatic(verbosity)
}

// Initialized, register and collect Logger from stadards path.
//
// Returns:
//   *logs.Logger Pointer to Initialized system logger
//   error Any suitable error risen during code execution
func InitDeviceLogger() error {
	filePath, extension, fErr := findFileInPath(logPath, logFileName, logFileExtensions)
	if fErr != nil {
		return fErr
	}
	encoding, eErr := parsers.EncodingFromString(extension)
	if eErr != nil {
		return eErr
	}
	conf, err := logger.LoadLoggerConfigFromFile(encoding, filePath)

	if err != nil {
		return err
	}
	logs.InitStaticFromConfig(*conf)
	if err != nil {
		return err
	}
	return nil
}

// Initialized, register and collect Logger from stadards path.
//
// Parameters:
//   filePath (string) Absolute file path
//
// Returns:
//   *logs.Logger Pointer to Initialized system logger
//   error Any suitable error risen during code execution
func InitDeviceLoggerFromFile(filePath string) error {
	//TODO Complete!!
	lLogPath, lLogFileName, lLogFileExtensions, lErr := decomposeFilePath(filePath)
	if lErr != nil {
		return lErr
	}

	filePath, extension, fErr := findFileInPath(lLogPath, lLogFileName, lLogFileExtensions)
	if fErr != nil {
		return fErr
	}
	encoding, eErr := parsers.EncodingFromString(extension)
	if eErr != nil {
		return eErr
	}
	conf, err := logger.LoadLoggerConfigFromFile(encoding, filePath)

	if err != nil {
		return err
	}
	logs.InitStaticFromConfig(*conf)
	return nil
}

func main() {
	InitSimpleLogger(logs.DEBUG)
	logManager, _ := logs.New("main")
	logManager.Debug("I will survive")
}
