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
//   log.LogConfig representing Log level
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

// Transform log.LogLevel to representing string.
//
// Parameters:
//   level (log.LogLevel) Log level to transform
//
// Returns:
//   string representing string
//   error Any suitable error risen during code execution
func LogLevelToString(level LogLevel) (string, error) {
	if level < DEBUG || level > NOLOG {
		return "", errors.New("logger::LogLevelToString::error : Wrong log level passed")
	}
	switch level {
	case DEBUG:
		return "DEBUG", nil
	case INFO:
		return "INFO ", nil
	case WARN:
		return "WARN ", nil
	case ERROR:
		return "ERROR", nil
	case FATAL:
		return "FATAL", nil
	case NOLOG:
		return "NOLOG", nil
	}
	return "INFO", nil
}

type _loggerEngineStruct struct {
	level    LogLevel
	_config  LogConfig
	_running bool
}

type loggerStruct struct {
	name    string
	_engine _loggerEngine
}

type _loggerEngine interface {
	// Log element in required verbosity level.
	//
	// Parameters:
	//   level (log.LogLevel) Required logging level
	//   value (interface{}) Element to log
	//   err (error) Error element
	//   name (string) logger class/package/reference name
	Log(LogLevel, interface{}, error, string)
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
	// Start logger deamons.
	//
	// Returns:
	//   bool is SIMPLE, no LifeCycle Logger Engine
	IsSimple() bool
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
}

func (l *_loggerEngineStruct) _writeLog(level LogLevel, name string, logText interface{}) {
	fmt.Println(logText)
}

func (l *_loggerEngineStruct) Log(level LogLevel, value interface{}, err error, name string) {
	var LogDate time.Time = time.Now()
	logLevel, err := LogLevelToString(level)
	if err != nil {
		logLevel = "INFO "
	}
	if value == nil {
		l._writeLog(level, name, fmt.Sprintf("[%s] %s - %s - %s\n", LogDate.Format("2006-01-02 15:04:05.000"), logLevel, name, err.Error()))
	} else if err == nil {
		l._writeLog(level, name, fmt.Sprintf("[%s] %s - %s - %v\n", LogDate.Format("2006-01-02 15:04:05.000"), logLevel, name, value))
	} else {
		l._writeLog(level, name, fmt.Sprintf("[%s] %s - %s - %v - error : \n", LogDate.Format("2006-01-02 15:04:05.000"), name, logLevel, value, err.Error()))
	}
}

func (l *_loggerEngineStruct) Start(conf LogConfig) {
	l._config = conf
	l.runConfig()
}

func (l *_loggerEngineStruct) runConfig() {
	l._running = true
	//TODO Prepare Config and run logger tasks
}

func (l *_loggerEngineStruct) Stop() {
	l._running = false
}

func (l *_loggerEngineStruct) IsRunning() bool {
	return l._running
}

func (l *_loggerEngineStruct) IsSimple() bool {
	return l._config.loggers == nil
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

func (l *loggerStruct) _writeLog(logText interface{}) {
	fmt.Println(logText)
}

func (l *loggerStruct) Fatal(err error) {
	l._engine.Log(FATAL, nil, err, l.name)
}

func (l *loggerStruct) FatalS(value interface{}) {
	if _, ok := reflect.TypeOf(value).MethodByName("Error"); ok {
		l._engine.Log(FATAL, nil, value.(error), l.name)
	} else {
		l._engine.Log(FATAL, value, nil, l.name)
	}
}

func (l *loggerStruct) Error(err error) {
	l._engine.Log(ERROR, nil, err, l.name)
}

func (l *loggerStruct) ErrorS(value interface{}) {
	if _, ok := reflect.TypeOf(value).MethodByName("Error"); ok {
		l._engine.Log(ERROR, nil, value.(error), l.name)
	} else {
		l._engine.Log(ERROR, value, nil, l.name)
	}
}

func (l *loggerStruct) Warning(value interface{}) {
	l._engine.Log(WARN, value, nil, l.name)
}

func (l *loggerStruct) Info(value interface{}) {
	l._engine.Log(INFO, value, nil, l.name)
}

func (l *loggerStruct) Debug(value interface{}) {
	l._engine.Log(DEBUG, value, nil, l.name)
}

// Create new Logger from verbosity.
//
// Parameters:
//   verbosity (log.LogLevel) Level for logging
//
// Returns:
//   Logger Default simple logger
func New(name string) (Logger, error) {
	if _loggerEng == nil {
		return nil, errors.New("Log Engine has not been initialized")
	}
	return &loggerStruct{
		name:    name,
		_engine: _loggerEng,
	}, nil
}

var NULL_LOG_CONFIG LogConfig = LogConfig{
	loggers: nil,
}

var _loggerEng _loggerEngine = nil

var _logger Logger = nil

func getEngine(verbosity LogLevel) _loggerEngine {
	return &_loggerEngineStruct{
		level:    verbosity,
		_running: false,
		_config:  NULL_LOG_CONFIG,
	}
}

func getEngineFromConfig(config LogConfig) (_loggerEngine, error) {
	verbosity, err := LogLevelFromString(config.verbosity)
	if err != nil {
		return nil, err
	}
	return &_loggerEngineStruct{
		level:    verbosity,
		_running: false,
		_config:  config,
	}, nil
}

// Istantiate global simple Logger from verbosity.
//
// Parameters:
//   verbosity (log.LogLevel) Level for logging
func InitStatic(verbosity LogLevel) {
	if _loggerEng == nil || !_loggerEng.IsSimple() {
		if _loggerEng != nil && _loggerEng.IsRunning() {
			_loggerEng.Stop()
		}
		_loggerEng = getEngine(verbosity)
	}
}

// Istantiate global full lifecycle Logger from Cofiguration.
//
// Parameters:
//   config (log.LogConfig) Logging exteded configuration
//
// Returns:
//   error Any suitable error risen during code execution
func InitStaticFromConfig(config LogConfig) {
	if _loggerEng == nil || _loggerEng.IsSimple() {
		if _loggerEng != nil && _loggerEng.IsRunning() {
			_loggerEng.Stop()
		}
		_loggerEng, _ = getEngineFromConfig(config)
	}
}
