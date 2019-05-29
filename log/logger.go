package log

import (
	"fmt"
	"github.com/hellgate75/general_utils/common"
	"github.com/hellgate75/general_utils/errors"
	parser "github.com/hellgate75/general_utils/log/parser"
	"reflect"
	"strings"
	"sync"
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
func StringToLogLevel(text string) (LogLevel, error) {
	if strings.TrimSpace(text) == "" {
		return 0, errors.New("logger::StringToLogLevel::error : Empty input string")
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

type LogMapItem struct {
	DefaultVerbosity LogLevel
	Dists            []LogMapDistItem
	Streams          []parser.LogStream
}

type _loggerEngineStruct struct {
	sync.RWMutex
	level    LogLevel
	Config   LogConfig
	Running  bool
	TestChan *chan interface{}
	Map      map[string]LogMapItem
}

type LogMapDistItem struct {
	Appender LogAppender
	Writer   LogWriter
	Stream   parser.LogStream
	Format   string
	Device   string
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

	if l.TestChan == nil {
		fmt.Println(logText)
	} else {
		*l.TestChan <- logText
	}
}

func (l *_loggerEngineStruct) Log(level LogLevel, value interface{}, err error, name string) {
	var LogDate time.Time = time.Now()
	logLevel, err := LogLevelToString(level)
	if err != nil {
		logLevel = "INFO "
	}
	if value == nil {
		l._writeLog(level, name, fmt.Sprintf("[%s] %s - %s - %s", LogDate.Format("2006-01-02 15:04:05.000"), logLevel, name, err.Error()))
	} else if err == nil {
		l._writeLog(level, name, fmt.Sprintf("[%s] %s - %s - %v", LogDate.Format("2006-01-02 15:04:05.000"), logLevel, name, value))
	} else {
		l._writeLog(level, name, fmt.Sprintf("[%s] %s - %s - %v - error : %s", LogDate.Format("2006-01-02 15:04:05.000"), name, logLevel, value, err.Error()))
	}
}

func (l *_loggerEngineStruct) Start(conf LogConfig) {
	l.Config = conf
	l._runConfig()
}

func (l *_loggerEngineStruct) _runConfig() {
	l.Running = true

	var warnings []string

	var appendersMap map[string]LogAppender = make(map[string]LogAppender)

	for _, app := range l.Config.Appenders {
		appendersMap[app.AppenderName] = app
	}
	var writersMap map[string]LogWriter = make(map[string]LogWriter)

	for _, wtr := range l.Config.Writers {
		writersMap[wtr.WriterName] = wtr
	}

	//loggerName := l.Config.LoggerName
	//globalVerb, gVErr := StringToLogLevel(l.Config.Verbosity)

	for _, lgr := range l.Config.Loggers {
		appenderName := lgr.AppenderName
		writerName := lgr.WriterName
		appender, okAppender := appendersMap[appenderName]
		writer, okWriter := writersMap[writerName]
		if !okAppender {
			//Missing appender
			warnings = append(warnings, fmt.Sprintf("Missing Appender <%s> in Log Config", appenderName))
		} else {
			if !okWriter {
				//Missing writer
				warnings = append(warnings, fmt.Sprintf("Missing Writer <%s> in Log Config", writerName))
			} else {
				//All is fine
				for _, fltr := range lgr.Filters {
					pkgName := fltr.PackageName
					pkgVerbosity, _ := StringToLogLevel(fltr.Verbosity)
					var item LogMapItem
					if _, ok := l.Map[pkgName]; ok {
						//Update Item
						item, _ = l.Map[pkgName]
					} else {
						//Create Item
						item = LogMapItem{
							DefaultVerbosity: pkgVerbosity,
							Dists:            []LogMapDistItem{},
							Streams:          []parser.LogStream{},
						}

						l.Map[pkgName] = item
					}
					//Update Item
					lst, lstErr := parser.WriterTypeToLogStreamType(writer.WriterType)
					if lstErr != nil {
						//Log warning lstErr
						lst = parser.StdOutStreamType
						warnings = append(warnings, fmt.Sprintf("Lost In/Out Stream (%s), replaced with StdOut - throw clause : %s", writer.WriterType, lstErr.Error()))
					}
					streamIO, streamErr := parser.New(lst, writer.WriterEncoding)
					if streamErr == nil {
						//						type LogMapItem struct {
						//							DefaultVerbosity	LogLevel
						//							Dists           	[]LogMapDistItem
						//							Streams         	[]parser.LogStream
						//						}
						format, _ := common.StreamInOutFormatToString(writer.WriterEncoding)
						device, _ := common.WriterTypeToString(writer.WriterType)
						item.Dists = append(item.Dists, LogMapDistItem{
							Appender: appender,
							Writer:   writer,
							Format:   format,
							Stream:   streamIO,
							Device:   device,
						})
						item.Streams = append(item.Streams, streamIO)
						item.DefaultVerbosity = pkgVerbosity
						delete(l.Map, pkgName)
						l.Map[pkgName] = item
					} else {
						//Log warning streamErr
						warnings = append(warnings, fmt.Sprintf("Lost Logger (%s) - throw clause : %s", writer.WriterName, streamErr.Error()))
					}
				}

			}
		}
	}

	//TODO Prepare Config and run logger tasks

}

func (l *_loggerEngineStruct) Stop() {
	l.Running = false
}

func (l *_loggerEngineStruct) IsRunning() bool {
	return l.Running
}

func (l *_loggerEngineStruct) IsSimple() bool {
	return l.Config.Loggers == nil
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

var NULL_LOGConfig LogConfig = LogConfig{
	Loggers: nil,
}

var _loggerEng _loggerEngine = nil

var _logger Logger = nil

func getEngine(verbosity LogLevel, testChan *chan interface{}) _loggerEngine {
	return &_loggerEngineStruct{
		level:    verbosity,
		Running:  false,
		Config:   NULL_LOGConfig,
		TestChan: testChan,
		Map:      make(map[string]LogMapItem),
	}
}

func getEngineFromConfig(config LogConfig, testChan *chan interface{}) (_loggerEngine, error) {
	verbosity, err := StringToLogLevel(config.Verbosity)
	if err != nil {
		return nil, err
	}
	return &_loggerEngineStruct{
		level:    verbosity,
		Running:  false,
		Config:   config,
		TestChan: testChan,
		Map:      make(map[string]LogMapItem),
	}, nil
}

// Remove instance of global simple Logger.
func ResetStaticLoggerEngine() {
	if _loggerEng != nil && _loggerEng.IsRunning() {
		_loggerEng.Stop()
	}
	_loggerEng = nil
}

// Istantiate global simple Logger from verbosity, writing to StdOut.
//
// Parameters:
//   verbosity (log.LogLevel) Level for logging
func InitStaticLoggerEngine(verbosity LogLevel) {
	InitStaticTestLoggerEngine(verbosity, nil)
}

// Istantiate global simple test Logger from verbosity, writing to passed channel the log instead the StdOut
//
// Parameters:
//   verbosity (log.LogLevel) Level for logging
//	 testChan (*chan interface{}) Fake output, or nil in case of real Logger
func InitStaticTestLoggerEngine(verbosity LogLevel, testChan *chan interface{}) {
	if _loggerEng == nil || !_loggerEng.IsSimple() {
		if _loggerEng != nil && _loggerEng.IsRunning() {
			_loggerEng.Stop()
		}
		_loggerEng = getEngine(verbosity, testChan)
	}
}

// Istantiate global full lifecycle Logger from Configuration.
//
// Parameters:
//   config (log.LogConfig) Logging exteded configuration
func InitStaticLoggerEngineFromConfig(config LogConfig) {
	InitStaticTestLoggerEngineFromConfig(config, nil)
}

// Istantiate global full lifecycle Logger from Configuration.
//
// Parameters:
//   config (log.LogConfig) Logging exteded configuration
//	 testChan (*chan interface{}) Fake output, or nil in case of real Logger
func InitStaticTestLoggerEngineFromConfig(config LogConfig, testChan *chan interface{}) {
	if _loggerEng == nil || _loggerEng.IsSimple() {
		if _loggerEng != nil && _loggerEng.IsRunning() {
			_loggerEng.Stop()
		}
		_loggerEng, _ = getEngineFromConfig(config, testChan)
	}
}
