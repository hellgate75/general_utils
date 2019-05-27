package general_utils

//package main

import (
	"fmt"
	"github.com/hellgate75/general_utils/errors"
	"github.com/hellgate75/general_utils/log"
	"github.com/hellgate75/general_utils/parsers"
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
	logFileExtensions []string = []string{"yaml", "json", "xml"}
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
//   *log.Logger Pointer to Initialized system logger
//   error Any suitable error risen during code execution
func InitSimpleLogger(verbosity log.LogLevel) {
	log.InitStatic(verbosity)
}

// Initialized, register and collect Logger from stadards path.
//
// Returns:
//   *log.Logger Pointer to Initialized system logger
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
	log.InitStaticFromConfig(*conf)
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
//   *log.Logger Pointer to Initialized system logger
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
	log.InitStaticFromConfig(*conf)
	return nil
}

// Initialize loggers for all packages. To run after the logger initialization only
func InitializeLoggers() {
	parsers.InitLogger()
	streams.InitLogger()
}

func _getTestConfig() log.LogConfig {
	return log.LogConfig{
		LoggerName: "myLogger",
		Verbosity:  "debug",
		Appenders: []log.LogAppender{
			log.LogAppender{
				AppenderName: "standardAppender",
				Verbosity:    "debug",
				DateFormat:   "2006-01-02 15:04:05.000",
			},
		},
		Writers: []log.LogWriter{
			log.LogWriter{
				WriterName: "stdOutWriter",
				WriterType: log.StdOutWriter,
			},
			log.LogWriter{
				WriterName:  "fileWriter",
				WriterType:  log.FileWriter,
				Destination: "C:\\sample-log.log",
			},
		},
		Loggers: []log.LoggerInfo{
			log.LoggerInfo{
				AppenderName: "standardAppender",
				WriterName:   "stdOutWriter",
				Filters:      []log.LogFilter{},
			},
			log.LoggerInfo{
				AppenderName: "standardAppender",
				WriterName:   "fileWriter",
				Filters: []log.LogFilter{
					log.LogFilter{
						PackageName: "main",
						Verbosity:   "info",
					},
				},
			},
		},
	}
}

func WriteTestResources() {
	InitSimpleLogger(log.DEBUG)
	InitializeLoggers()
	//	logManager, _ := log.New("main")
	//	logManager.Debug("I will survive")
	var config log.LogConfig = _getTestConfig()
	var path string = streams.GetCurrentPath() + fmt.Sprintf("%c", os.PathSeparator) + "test" + fmt.Sprintf("%c", os.PathSeparator) + "resources"
	os.MkdirAll(path, 666)
	var filePath string = path + fmt.Sprintf("%c", os.PathSeparator) + logFileName

	parser1, _ := parsers.New(parsers.YAML)
	err1 := parser1.SerializeToFile(filePath+".yaml", config)
	if err1 != nil {
		fmt.Printf("error 1 (yaml) : %v", err1)
	}

	parser2, _ := parsers.New(parsers.JSON)
	err2 := parser2.SerializeToFile(filePath+".json", config)
	if err1 != nil {
		fmt.Printf("error 2 (json) : %v", err2)
	}

	parser3, _ := parsers.New(parsers.XML)
	err3 := parser3.SerializeToFile(filePath+".xml", config)
	if err1 != nil {
		fmt.Printf("error 3 (xml) : %v", err3)
	}
}

//func main() {
//	//InitSimpleLogger(log.DEBUG)
//	//InitializeLoggers()
//	//	logManager, _ := log.New("main")
//	//	logManager.Debug("I will survive")
//	WriteTestResources()
//}
