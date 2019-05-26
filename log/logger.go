package log

import (
	"fmt"
	"github.com/hellgate75/general_utils/errors"
	"reflect"
	"strings"
	"time"
)

// Log Verbosity Level, describes logging level
type LogLevel int

const (
	// DEBUG Verbosity level
	DEBUG LogLevel = 1
	// INFO Verbosity level
	INFO LogLevel = 2
	// WARN Verbosity level
	WARN LogLevel = 3
	// ERROR Verbosity level
	ERROR LogLevel = 4
	// FATAL Verbosity level
	FATAL LogLevel = 5
	// NOLOG Verbosity level
	NOLOG LogLevel = 6
)

// Transform text to representing Log Verbosity Level element.
//
// Parameters:
//   text (string) Text to parse
//
// Returns:
//   LogConfig Loaded Log Configuration or Nil in case of  error
//   error Any suitable error risen during code execution
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
	level    LogLevel
	_config  LogConfig
	_running bool
}

// Iterface describes main logger features
type Logger interface {
	// Log element in required verbosity level.
	//
	// Parameters:
	//   level (log.LogLevel) Required logging level
	//   value (interface{}) Element to log
	Log(LogLevel, interface{})
	// Log error in FATAL verbosity level.
	//
	// Parameters:
	//   err (error) Error to log
	Fatal(error)
	// Log element in FATAL verbosity level.
	//
	// Parameters:
	//   value (interface{}) Element to log
	FatalS(interface{})
	// Log error in ERROR verbosity level.
	//
	// Parameters:
	//   err (error) Error to log
	Error(error)
	// Log element in FATAL verbosity level.
	//
	// Parameters:
	//   value (interface{}) Element to log
	ErrorS(interface{})
	// Log element in WARN verbosity level.
	//
	// Parameters:
	//   value (interface{}) Element to log
	Warning(value interface{})
	// Log element in INFO verbosity level.
	//
	// Parameters:
	//   value (interface{}) Element to log
	Info(value interface{})
	// Log element in DEBUG verbosity level.
	//
	// Parameters:
	//   value (interface{}) Element to log
	Debug(value interface{})
	// Start logger deamons.
	//
	// Parameters:
	//   conf (log.LogConfig) Element to log
	Start(LogConfig)
	// Stop logger deamons.
	Stop()
	// Start logger deamons.
	//
	// Returns:
	//   bool is running state
	IsRunning() bool
}

func (l *loggerStruct) Start(conf LogConfig) {
	l._config = conf
	l.runConfig()
}

func (l *loggerStruct) runConfig() {
	l._running = true
	//TODO Prepare Config and run logger tasks
}

func (l *loggerStruct) Stop() {
	l._running = false
}

func (l *loggerStruct) IsRunning() bool {
	return l._running
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

// Create new Logger from verbosity.
//
// Parameters:
//   verbosity (log.LogLevel) Level for logging
//
// Returns:
//   Logger Default simple logger
func New(verbosity LogLevel) Logger {
	return &loggerStruct{
		level: verbosity,
	}
}

var _logger Logger = nil

// Istantiate global simple Logger from verbosity.
//
// Parameters:
//   verbosity (log.LogLevel) Level for logging
func InitStatic(verbosity LogLevel) {
	_logger = New(verbosity)
}

// Istantiate global full lifecycle Logger from Cofiguration.
//
// Parameters:
//   config (log.LogConfig) Logging exteded configuration
//
// Returns:
//   error Any suitable error risen during code execution
func InitStaticFromConfig(config LogConfig) error {
	level, err := LogLevelFromString(config.verbosity)
	if err != nil {
		return err
	}
	_logger = New(level)
	return nil
}

func GetLogger() Logger {
	if _logger == nil {
		_logger = New(DEBUG)
	}
	return _logger
}

func GetLoggerByRef() *Logger {
	if _logger == nil {
		_logger = New(DEBUG)
	}
	return &_logger
}
