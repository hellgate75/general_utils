package logger

import (
	"fmt"
	"github.com/hellgate75/general_utils/common"
	"github.com/hellgate75/general_utils/log"
	"github.com/hellgate75/general_utils/parsers"
	"os"
	"testing"
)

func TestSetLogConfig(t *testing.T) {
	var expected string = "DEBUG"
	SetLogConfig(log.LogConfig{
		Verbosity: expected,
	})
	if conf.Verbosity != expected {
		t.Fatal(fmt.Sprintf("TestSetLogConfig::error : Expected <%s> but Given <%s>", expected, conf.Verbosity))
	}
}
func TestGetLogConfig(t *testing.T) {
	var expected string = "INFO"
	conf = log.LogConfig{
		Verbosity: expected,
	}
	config := GetLogConfig()
	if config.Verbosity != expected {
		t.Fatal(fmt.Sprintf("TestGetLogConfig::error : Expected <%s> but Given <%s>", expected, config.Verbosity))
	}
}

var testFolder = fmt.Sprintf("%s%c%s%c%s%c%s", "..", os.PathSeparator, "..", os.PathSeparator, "test", os.PathSeparator, "sample")
var testFile1 = fmt.Sprintf("%s%c%s", testFolder, os.PathSeparator, "sample_file_1.json")

const defaultDateFormat string = "2006-01-02 15:04:05.000"

func _createLogConfig() log.LogConfig {
	return log.LogConfig{
		Appenders: []log.LogAppender{
			log.LogAppender{
				AppenderName: "StdOutAppender",
				DateFormat:   defaultDateFormat,
				Verbosity:    "INFO",
			},
		},
		Writers: []log.LogWriter{
			log.LogWriter{
				WriterName:     "StdOutWriter",
				WriterType:     common.StdOutWriter,
				WriterEncoding: common.PlainTextFormat,
				Destination:    "",
			},
		},
		LoggerName: "StdOutLogger",
		Loggers: []log.LoggerInfo{
			log.LoggerInfo{
				AppenderName: "StdOutAppender",
				WriterName:   "StdOutWriter",
				Filters:      []log.LogFilter{},
			},
		},
	}
}

func TestLoadLoggerConfigFromFile(t *testing.T) {
	logConf, err := LoadLoggerConfigFromFile(parsers.Encoding(0), testFile1)
	if err == nil {
		t.Fatal("Expected error but not arisen!!")
	}
	expected := _createLogConfig()
	p, errP := parsers.New(parsers.JSON)
	if errP != nil {
		t.Fatal(fmt.Sprintf("Parser Creating - Arisen unexpected error : %s", errP.Error()))
	}
	os.MkdirAll(testFolder, 666)
	defer func(testFolder string) {
		os.RemoveAll(testFolder)
	}(testFolder)
	errS := p.SerializeToFile(testFile1, expected)
	if errS != nil {
		t.Fatal(fmt.Sprintf("LogConfig Serialization - Arisen unexpected error : %s", errS.Error()))
	}
	logConf, err = LoadLoggerConfigFromFile(parsers.JSON, testFile1)
	if err != nil {
		t.Fatal(fmt.Sprintf("LogConf Load - Arisen unexpected error : %s", err.Error()))
	}
	if logConf.Verbosity != expected.Verbosity {
		t.Fatal(fmt.Sprintf("LogConf Load - Wrong result in LoadLoggerConfigFromFile, Expected <%v> but Given <%v>", expected, logConf))

	}
}
