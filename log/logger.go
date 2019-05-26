package log

import (
	"fmt"
	"github.com/hellgate75/general_utils/errors"
	"reflect"
	"strings"
	"time"
)

type LogLevel int

const (
	DEBUG LogLevel = 1
	INFO  LogLevel = 2
	WARN  LogLevel = 3
	ERROR LogLevel = 4
	FATAL LogLevel = 5
	NOLOG LogLevel = 6
)

func LogLevelFromString(text string) (LogLevel, error) {
	if strings.TrimSpace(text) == "" {
		return 0, errors.New("logger::LogLevelFromString::error : Empty input string")
	}
	value := strings.ToUpper(text)
	switch value {
	case "DEBUG":
		return DEBUG, nil
	case "INFO":
		return INFO, nil
	case "WARN":
		return WARN, nil
	case "ERROR":
		return ERROR, nil
	case "FATAL":
		return FATAL, nil
	case "NOLOG":
		return NOLOG, nil
	}
	return INFO, nil
}

type loggerStruct struct {
	level LogLevel
}

type Logger interface {
	Log(LogLevel, interface{})
	Fatal(error)
	FatalS(interface{})
	Error(error)
	ErrorS(interface{})
	Warning(value interface{})
	Info(value interface{})
	Debug(value interface{})
}

func (l *loggerStruct) init() {

}

func (l *loggerStruct) Log(level LogLevel, value interface{}) {
	switch level {
	case DEBUG:
		l.Debug(value)
	case INFO:
		l.Info(value)
	case WARN:
		l.Warning(value)
	case ERROR:
		l.ErrorS(value)
	case FATAL:
		l.FatalS(value)

	}
}

func (l *loggerStruct) Fatal(err error) {
	if l.level <= FATAL {
		var LogDate time.Time = time.Now()
		fmt.Printf("[%s] FATAL %s\n", LogDate.Format("2006-01-02 15:04:05.000"), err.Error())
	}
}

func (l *loggerStruct) FatalS(value interface{}) {
	if l.level <= FATAL {
		var LogDate time.Time = time.Now()
		var text string = ""
		if _, ok := reflect.TypeOf(value).MethodByName("Error"); ok {
			text = value.(error).Error()
		} else {
			text = fmt.Sprintf("%v", value)
		}
		fmt.Printf("[%s] FATAL %s\n", LogDate.Format("2006-01-02 15:04:05.000"), text)
	}
}

func (l *loggerStruct) Error(err error) {
	if l.level <= ERROR {
		var LogDate time.Time = time.Now()
		fmt.Printf("[%s] ERROR %s\n", LogDate.Format("2006-01-02 15:04:05.000"), err.Error())
	}
}

func (l *loggerStruct) ErrorS(value interface{}) {
	if l.level <= ERROR {
		var LogDate time.Time = time.Now()
		var text string = ""
		if _, ok := reflect.TypeOf(value).MethodByName("Error"); ok {
			text = value.(error).Error()
		} else {
			text = fmt.Sprintf("%v", value)
		}
		fmt.Printf("[%s] ERROR %s\n", LogDate.Format("2006-01-02 15:04:05.000"), text)
	}
}

func (l *loggerStruct) Warning(value interface{}) {
	if l.level <= WARN {
		var LogDate time.Time = time.Now()
		fmt.Printf("[%s] WARN  %v\n", LogDate.Format("2006-01-02 15:04:05.000"), value)
	}
}

func (l *loggerStruct) Info(value interface{}) {
	if l.level <= INFO {
		var LogDate time.Time = time.Now()
		fmt.Printf("[%s] INFO  %v\n", LogDate.Format("2006-01-02 15:04:05.000"), value)
	}
}

func (l *loggerStruct) Debug(value interface{}) {
	if l.level <= DEBUG {
		var LogDate time.Time = time.Now()
		fmt.Printf("[%s] DEBUG %v\n", LogDate.Format("2006-01-02 15:04:05.000"), value)
	}
}

func New(verbosity LogLevel) Logger {
	return &loggerStruct{
		level: verbosity,
	}
}

var _logger Logger = nil

func InitStatic(verbosity LogLevel) {
	_logger = New(verbosity)
}

func GetLogger() Logger {
	if _logger == nil {
		_logger = New(DEBUG)
	}
	return _logger
}
