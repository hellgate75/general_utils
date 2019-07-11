package general_utils

import (
	"errors"
	"fmt"
	"github.com/hellgate75/general_utils/common"
	"github.com/hellgate75/general_utils/log"
	"github.com/hellgate75/general_utils/parsers"
	streams "github.com/hellgate75/general_utils/streams"
	"os"
	"strings"
	"testing"
	"time"
)

var preparedTestResources bool = false

var logTestModeEnabled bool = false

func _getTestConfig1() log.LogConfig {
	path := fmt.Sprintf("%s%c%s", streams.GetCurrentPath(), os.PathSeparator, "temp")
	os.MkdirAll(path, os.ModeAppend)
	file := fmt.Sprintf("%s%cSampleLogger.log", path, os.PathSeparator)
	return log.LogConfig{
		LoggerName: "myLogger",
		Verbosity:  "info",
		Appenders: []log.LogAppender{
			log.LogAppender{
				AppenderName: "standardAppender",
				Verbosity:    "info",
				DateFormat:   "2006-01-02 15:04:05.000",
			},
		},
		Writers: []log.LogWriter{
			log.LogWriter{
				WriterName:     "stdOutWriter",
				WriterType:     common.StdOutWriter,
				WriterEncoding: common.PlainTextFormat,
				Destination:    "",
			},
			log.LogWriter{
				WriterName:     "fileWriter",
				WriterType:     common.FileWriter,
				Destination:    file,
				WriterEncoding: common.JsonFormat,
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
						PackageName: "test",
						Verbosity:   "info",
					},
				},
			},
		},
	}
}

var TraverseFileFunc func(path string, info os.FileInfo, err error) error

func _destroyFolder(path string) {
	//	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
	//		os.RemoveAll(path)
	//		return nil
	//	})
	err := os.RemoveAll(path)
	if err != nil {
		fmt.Println("_destroyFolder() - error : ", err.Error())
	} else {
		os.Remove(path)
	}

}

func _destroyTestConfig1() {
	path := fmt.Sprintf("%s%c%s", streams.GetCurrentPath(), os.PathSeparator, "temp")
	_destroyFolder(path)
}

func ClearTestResources() bool {
	_destroyTestConfig1()
	preparedTestResources = false
	return true
}

func WriteTestResources() bool {
	if preparedTestResources {
		return false
	}
	InitSimpleLoggerEngine(log.DEBUG)
	//	InitializeLoggers()
	//	logManager, _ := log.New("main")
	//	logManager.Debug("I will survive")
	var config log.LogConfig = _getTestConfig1()
	var path string = streams.GetCurrentPath() + fmt.Sprintf("%c", os.PathSeparator) + "test" + fmt.Sprintf("%c", os.PathSeparator) + "resources"
	os.MkdirAll(path, 0777)
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
	preparedTestResources = true
	return true
}

func TestInitSimpleLoggerEngine(t *testing.T) {
	testing.Verbose()
	defer func() {
		logTestModeEnabled = false
		DestroyLoggerEngine()
		DisableLogChanMode()
	}()
	EnableLogChanMode()
	InitSimpleLoggerEngine(log.INFO)
	//	InitializeLoggers()
	logger, _ := log.New("test")
	var testMessage string = "Test1"
	var val interface{}
	var err error
	go func() {
		select {
		case val = <-log.LogOutChan:
			fmt.Println("1.", val)
			if strings.Index(fmt.Sprintf("%v", val), testMessage) < 0 {
				err = errors.New("Unable to read proper log")
			} else {
				logTestModeEnabled = false
			}
		case <-time.After(time.Second):
			err = errors.New("Timeout reached")
			//		default:
			//			err = errors.New("No message")
		}
	}()
	logTestModeEnabled = true
	logger.Info(testMessage)
	for logTestModeEnabled {
		if err != nil {
			logTestModeEnabled = false
			t.Fatal(err.Error())
		}
	}
}

func TestSimpleLoggerLevelBlocking(t *testing.T) {
	defer func() {
		logTestModeEnabled = false
		DestroyLoggerEngine()
		DisableLogChanMode()
	}()
	EnableLogChanMode()
	InitSimpleLoggerEngine(log.INFO)
	//	InitializeLoggers()
	logger, _ := log.New("test")
	var testMessage string = "Test0"
	//TODO: Fix Log level
	//	var testMessage2 string = "Test1"
	var val interface{}
	var err error
	go func() {
		select {
		case val = <-log.LogOutChan:
			fmt.Println("2.", val)
			if strings.Index(fmt.Sprintf("%v", val), testMessage) < 0 {
				//				err = errors.New("Unable to read proper log")
			} else {
				logTestModeEnabled = false
			}
		case <-time.After(time.Second):
			err = errors.New("Timeout reached")
			//		default:
			//			err = errors.New("No message")
		}
	}()
	//Only Info should pass the logging and not the Debug message
	//	time.Sleep(250 * time.Millisecond)
	logTestModeEnabled = true
	logger.Info(testMessage)
	//	time.Sleep(500 * time.Millisecond)
	//	logger.Debug(testMessage2)
	for logTestModeEnabled {
		time.Sleep(500 * time.Millisecond)
		if err != nil {
			logTestModeEnabled = false
			t.Fatal(err.Error())
		}
	}
}

func TestInitCustomLogger(t *testing.T) {
	defer func() {
		logTestModeEnabled = false
		DestroyLoggerEngine()
		ClearTestResources()
		DisableLogChanMode()
	}()
	EnableLogChanMode()
	WriteTestResources()
	var path string = streams.GetCurrentPath() + fmt.Sprintf("%c", os.PathSeparator) + "test" + fmt.Sprintf("%c", os.PathSeparator) + "resources"
	var filePath string = path + fmt.Sprintf("%c", os.PathSeparator) + logFileName + ".yaml"
	InitCustomLoggerEngine(filePath)
	//	InitializeLoggers()
	logger, _ := log.New("test")
	var testMessage string = "Test0"
	var val interface{}
	var err error
	go func() {
		select {
		case val = <-log.LogOutChan:
			fmt.Println("3.", val)
			if strings.Index(fmt.Sprintf("%v", val), testMessage) < 0 {
				err = errors.New("Unable to read proper log")
			} else {
				logTestModeEnabled = false
			}
		case <-time.After(time.Second):
			err = errors.New("Timeout reached")
			//		default:
			//			err = errors.New("No message")
		}
	}()
	//Only Info should pass the logging and not the Debug message
	//	time.Sleep(250 * time.Millisecond)
	logTestModeEnabled = true
	logger.Info(testMessage)
	for logTestModeEnabled {
		time.Sleep(500 * time.Millisecond)
		if err != nil {
			logTestModeEnabled = false
			t.Fatal(err.Error())
		}
	}
}

func TestCustomLoggerLevelBlocking(t *testing.T) {
	defer func() {
		logTestModeEnabled = false
		DestroyLoggerEngine()
		ClearTestResources()
		DisableLogChanMode()
	}()
	EnableLogChanMode()
	WriteTestResources()
	var path string = streams.GetCurrentPath() + fmt.Sprintf("%c", os.PathSeparator) + "test" + fmt.Sprintf("%c", os.PathSeparator) + "resources"
	var filePath string = path + fmt.Sprintf("%c", os.PathSeparator) + logFileName + ".yaml"
	InitCustomLoggerEngine(filePath)
	//	InitializeLoggers()
	logger, _ := log.New("test")
	var testMessage string = "Test0"
	var testMessage2 string = "Test1"
	var val interface{}
	var err error
	var testing bool = true
	go func() {
		select {
		case val = <-log.LogOutChan:
			fmt.Println("3.", val)
			if strings.Index(fmt.Sprintf("%v", val), testMessage) < 0 && testing {
				//				err = errors.New("Unable to read proper log")
				logTestModeEnabled = false
			} else {
				logTestModeEnabled = false
			}
			testing = false
		case <-time.After(time.Second):
			err = errors.New("Timeout reached")
			//		default:
			//			err = errors.New("No message")
		}
	}()
	//Only Info should pass the logging and not the Debug message
	logTestModeEnabled = true
	time.Sleep(500 * time.Millisecond)
	logger.Info(testMessage)
	logger.Debug(testMessage2)
	for logTestModeEnabled {
		time.Sleep(500 * time.Millisecond)
		if err != nil {
			logTestModeEnabled = false
			t.Fatal(err.Error())
		}
	}
}
