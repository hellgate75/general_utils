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
	// Default log Path
	logPath string = "config"
	// Default log file name
	logFileName string = "log4go"
)

// Yellow color for command screen text
var yellowColor func(interface{}) color.Value = color.Yellow

// Red color for command screen text
var redColor func(interface{}) color.Value = color.Red

var (
	// Default log configration file extensions
	logFileExtensions []string = []string{"yaml", "json", "xml"}
)

// Enables Log to default Channel instead of target configuration
func EnableLogChanMode() *chan interface{} {
	log.LogChanEnabled = true
	log.LogOutChan = make(chan interface{})
	return &log.LogOutChan
}

// Disables Log to default Channel and restore log to target configuration
func DisableLogChanMode() {
	log.LogChanEnabled = false
	log.LogOutChan = nil
}

// Finds a file into a folder and look forward a file matching multiple file extensions
// Parameters:
//    folder (string) folder containing files
//    fileName (string) file name without extension
//    extensions ([]string) list of suitable file extensions
// Returns:
//    ( string absolute file path, stirng file extension, error Any error arisen during computation )
func FindFileInPath(folder string, fileName string, extensions []string) (string, string, error) {
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

// Decompose a file path to recover most relevant information
// Parameters:
//    filePath (string) file absolute path (with extension)
//    fileName (string) file name without extension
//    extensions ([]string) list of suitable file extensions
// Returns:
//    ( string absolute file folder path, string absolute file path, []stirng list of suitable extensions, error Any error arisen during computation )
func DecomposeFilePath(filePath string) (string, string, []string, error) {
	if strings.TrimSpace(filePath) == "" {
		return "", "", []string{}, errors.New("general_utils::error : Received in input an Empty String")
	}
	var extension string = "json"
	var path string = ""
	var file string = ""
	separator := fmt.Sprintf("%c", os.PathSeparator)

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

// Destroys and unregister Logger Engine
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
	filePath, extension, fErr := FindFileInPath(logPath, logFileName, logFileExtensions)
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
	lLogPath, lLogFileName, lLogFileExtensions, lErr := DecomposeFilePath(filePath)
	if lErr != nil {
		return lErr
	}

	filePath, extension, fErr := FindFileInPath(lLogPath, lLogFileName, lLogFileExtensions)
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
