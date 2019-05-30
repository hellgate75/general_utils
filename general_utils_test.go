package general_utils

import (
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

func _getTestConfig1() log.LogConfig {
	path := fmt.Sprintf("%s%c%s", streams.GetCurrentPath(), os.PathSeparator, "temp")
	os.MkdirAll(path, os.ModeAppend)
	file := fmt.Sprintf("%s%cSampleLogger.log", path, os.PathSeparator)
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
						PackageName: "main",
						Verbosity:   "info",
					},
				},
			},
		},
	}
}

func WriteTestResources() bool {
	if preparedTestResources {
		return false
	}
	InitSimpleLoggerEngine(log.DEBUG)
	InitializeLoggers()
	//	logManager, _ := log.New("main")
	//	logManager.Debug("I will survive")
	var config log.LogConfig = _getTestConfig1()
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
	preparedTestResources = true
	return true
}

func TestInitSimpleLoggerEngine(t *testing.T) {
	WriteTestResources()
	DestroyLoggerEngine()
	_logTestState = true
	_logTestOutChan = make(chan interface{})
	InitSimpleLoggerEngine(log.DEBUG)
	InitializeLoggers()
	logger, _ := log.New("test")
	var testMessage string = "Test1"
	var val interface{}
	go func() {
		select {
		case val = <-_logTestOutChan:
			if strings.Index(fmt.Sprintf("%v", val), testMessage) < 0 {
				t.Fatal("Unable to read proper log")
			}
		case <-time.After(5 * time.Millisecond):
			val = "xxxxxxxxxxxx"
		default:
			val = "xxxxxxxxxxxx"
		}
	}()
	logger.Debug(testMessage)
	time.Sleep(5 * time.Millisecond)
	_logTestState = false
	close(_logTestOutChan)
}
