package general_utils

//package main

import (
	"fmt"
	"github.com/hellgate75/general_utils/errors"
	"github.com/hellgate75/general_utils/log"
	"github.com/hellgate75/general_utils/parsers"
	logger "github.com/hellgate75/general_utils/parsers/logger"
	streams "github.com/hellgate75/general_utils/streams"
	color "github.com/logrusorgru/aurora"
	"os"
	"strings"
)

const (
	logPath     string = "config"
	logFileName string = "log4go"
)

var yellowColor func(interface{}) color.Value = color.Yellow

var (
	logFileExtensions []string = []string{"yaml", "json", "xml"}
)

func EnableLogChanMode() *chan interface{} {
	log.LogChanEnabled = true
	log.LogOutChan = make(chan interface{})
	return &log.LogOutChan
}

func DisableLogChanMode() {
	log.LogChanEnabled = false
	log.LogOutChan = nil
}

var redColor func(interface{}) color.Value = color.Red

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

	fmt.Println("original path: ", filePath)
	fmt.Println("separator: ", separator)

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

// Destroys and unregisters Logger.
func DestroyLoggerEngine() {
	log.ResetStaticLoggerEngine()
	log.LogOutChan = nil
	log.LogChanEnabled = false
}

// Initializes and registers Logger for StdOut.
//
// Params:
//   verbosity (log.LogLevel) Log verbosity level
func InitSimpleLoggerEngine(verbosity log.LogLevel) {
	if log.LogChanEnabled && log.LogOutChan != nil {
		fmt.Println(yellowColor("Logger :: Entering in test mode ..."))
	}
	log.InitStaticLoggerEngine(verbosity)
}

// Initializes and registers full Logger from base configuration file.
//
// Returns:
//   error Any suitable error risen during code execution
func InitDeviceLoggerEngine() error {
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
	if log.LogChanEnabled && log.LogOutChan != nil {
		fmt.Println(yellowColor("Logger :: Entering in test mode ..."))
	}
	log.InitStaticLoggerEngineFromConfig(*conf)
	return nil
}

// Initializes and registers full Logger from custom configuration file.
//
// Parameters:
//   filePath (string) Absolute file path
//
// Returns:
//   error Any suitable error risen during code execution
func InitCustomLoggerEngine(filePath string) error {
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
	fmt.Println(yellowColor(fmt.Sprintf("Logger :: LogOutChan : %v", log.LogOutChan)))
	if log.LogChanEnabled && log.LogOutChan != nil {
		fmt.Println(yellowColor("Logger :: Entering in test mode ..."))
	}
	log.InitStaticLoggerEngineFromConfig(*conf)
	return nil
}

// Initialize loggers for all sub packages in global_utils. To run after the logger initialization only
func InitializeLoggers() {
	parsers.InitLogger()
	streams.InitLogger()
}

//func main() {
//	//InitSimpleLogger(log.DEBUG)
//	//InitializeLoggers()
//	//	logManager, _ := log.New("main")
//	//	logManager.Debug("I will survive")
//	WriteTestResources()
//}
