package log

import (
	"fmt"
	"github.com/gookit/color"
	"github.com/hellgate75/general_utils/common"
	"github.com/hellgate75/general_utils/errors"
	parser "github.com/hellgate75/general_utils/log/parser"
	"reflect"
	"runtime"
	"strconv"
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
	level           LogLevel
	Config          LogConfig
	globalVerbosity LogLevel
	Running         bool
	TestChan        *chan interface{}
	Map             map[string]LogMapItem
}

type LogMapDistItem struct {
	Appender LogAppender
	Writer   LogWriter
	Stream   *parser.LogStream
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

const DEFAULT_VERBOSITY LogLevel = INFO

func finalizeLogMapItem(lmi *LogMapItem) {
	for _, stream := range lmi.Streams {
		stream.Close()
	}
}
func (l *_loggerEngineStruct) _writeLog(level LogLevel, value interface{}, err error, name string) {

	writeLogToStream := func(stream *parser.LogStream, level LogLevel, name string, logText string) {
		if l.TestChan == nil {
			(*stream).Write(logText)
			//			fmt.Println(logText)
		} else {
			*l.TestChan <- logText
		}
	}

	var defaultDateFormat string = "2006-01-02 15:04:05.000"
	var defaultLogEncoding common.StreamInOutFormat = common.PlainTextFormat
	var defaultLogSource parser.LogStreamType = parser.StdOutStreamType

	l.RLock()
	item, ok := l.Map[name]
	l.RUnlock()
	var streams []parser.LogStream = []parser.LogStream{}
	var dists []LogMapDistItem = []LogMapDistItem{}
	if !ok {
		//Key not found
		l.RLock()
		item, ok = l.Map["*"]
		l.RUnlock()
		if !ok {
			//No wildcat so proceeding with StdOut
			stream, err := parser.New(defaultLogSource, defaultLogEncoding)
			if err == nil {
				streams = append(streams, stream)
			}
			verb, errVerb := LogLevelToString(l.globalVerbosity)
			if errVerb != nil {
				verb, _ = LogLevelToString(DEFAULT_VERBOSITY)
			}
			dists = append(dists, LogMapDistItem{
				Appender: LogAppender{
					AppenderName: "defaultAppender",
					DateFormat:   defaultDateFormat,
					Verbosity:    verb,
				},
				Writer: LogWriter{
					WriterName:     "defaultWriter",
					WriterType:     common.StdOutWriter,
					Destination:    "",
					WriterEncoding: common.PlainTextFormat,
				},
				Stream: &stream,
			})
			item := LogMapItem{
				DefaultVerbosity: l.globalVerbosity,
				Dists:            dists,
				Streams:          streams,
			}
			runtime.SetFinalizer(item, finalizeLogMapItem)

			runtime.SetFinalizer(item, finalizeLogMapItem)
			l.Lock()
			l.Map["*"] = item
			l.Unlock()

		} else {
			//Found Wildcat Key
		}
	} else {
		//Key found
	}
	var dist LogMapDistItem
	for _, dist = range item.Dists {
		var LogDate time.Time = time.Now()
		logLevel, err := LogLevelToString(level)
		if err != nil {
			logLevel = "INFO "
		}

		//TODO: Verificare : livello log sia accettato. Per gli StdOutWriter usare colori per livello log

		if dist.Writer.WriterType == common.StdOutWriter {
			if value == nil {
				writeLogToStream(dist.Stream, level, name, fmt.Sprintf("[%s] %s - %s - %s", LogDate.Format("2006-01-02 15:04:05.000"), logLevel, name, err.Error()))
			} else if err == nil {
				writeLogToStream(dist.Stream, level, name, fmt.Sprintf("[%s] %s - %s - %v", LogDate.Format("2006-01-02 15:04:05.000"), logLevel, name, value))
			} else {
				writeLogToStream(dist.Stream, level, name, fmt.Sprintf("[%s] %s - %s - %v - error : %s", LogDate.Format("2006-01-02 15:04:05.000"), name, logLevel, value, err.Error()))
			}
		} else {
			if value == nil {
				writeLogToStream(dist.Stream, level, name, fmt.Sprintf("[%s] %s - %s - %s", LogDate.Format("2006-01-02 15:04:05.000"), logLevel, name, err.Error()))
			} else if err == nil {
				writeLogToStream(dist.Stream, level, name, fmt.Sprintf("[%s] %s - %s - %v", LogDate.Format("2006-01-02 15:04:05.000"), logLevel, name, value))
			} else {
				writeLogToStream(dist.Stream, level, name, fmt.Sprintf("[%s] %s - %s - %v - error : %s", LogDate.Format("2006-01-02 15:04:05.000"), name, logLevel, value, err.Error()))
			}
		}
	}

}

func (l *_loggerEngineStruct) Log(level LogLevel, value interface{}, err error, name string) {
	l._writeLog(level, value, err, name)
}

func (l *_loggerEngineStruct) Start(conf LogConfig) {
	l.Config = conf
	l._runConfigSetup()
}

func computePackageData(l *_loggerEngineStruct, writer LogWriter, appender LogAppender, pkgName string, pkgVerbosity LogLevel, warnings []string) []string {
	//	var defaultDateFormat string = "2006-01-02 15:04:05.000"
	var item LogMapItem
	if _, ok := l.Map[pkgName]; ok {
		//Update Item
		item, _ = l.Map[pkgName]
		delete(l.Map, pkgName)
	} else {
		//Create Item
		item = LogMapItem{
			DefaultVerbosity: l.globalVerbosity,
			Dists:            []LogMapDistItem{},
			Streams:          []parser.LogStream{},
		}
		runtime.SetFinalizer(item, finalizeLogMapItem)

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
		format, _ := common.StreamInOutFormatToString(writer.WriterEncoding)
		device, _ := common.WriterTypeToString(writer.WriterType)
		item.Dists = append(item.Dists, LogMapDistItem{
			Appender: appender,
			Writer:   writer,
			Format:   format,
			Stream:   &streamIO,
			Device:   device,
		})
		item.Streams = append(item.Streams, streamIO)
		item.DefaultVerbosity = pkgVerbosity
		runtime.SetFinalizer(item, finalizeLogMapItem)
		l.Map[pkgName] = item
	} else {
		//Log warning streamErr
		warnings = append(warnings, fmt.Sprintf("Lost Logger (%s) - throw clause : %s", writer.WriterName, streamErr.Error()))
	}
	return warnings
}

func (l *_loggerEngineStruct) _runConfigSetup() {
	l.Running = true
	red := color.FgRed.Render
	yellow := color.FgYellow.Render
	gray := color.FgGray.Render
	green := color.FgGreen.Render
	lightRed := color.FgLightRed.Render
	fmt.Println(gray("Logger for Go - Multiverse logger toolset. All right reserved. "))
	fmt.Println(gray("Loading configuration data ..."))

	var warnings []string

	var appendersMap map[string]LogAppender = make(map[string]LogAppender)

	for _, app := range l.Config.Appenders {
		appendersMap[app.AppenderName] = app
	}
	var writersMap map[string]LogWriter = make(map[string]LogWriter)

	for _, wtr := range l.Config.Writers {
		writersMap[wtr.WriterName] = wtr
	}

	loggerName := l.Config.LoggerName
	fmt.Println(gray("Logger name: "), lightRed(loggerName))
	globalVerb, gVErr := StringToLogLevel(l.Config.Verbosity)
	if gVErr != nil {
		globalVerb = DEFAULT_VERBOSITY
		defVerbStr, _ := LogLevelToString(DEFAULT_VERBOSITY)
		warnings = append(warnings, fmt.Sprintf("Missing Global Verbosity <%s> in Log Config, setting default %s", l.Config.Verbosity, defVerbStr))
	}
	l.globalVerbosity = globalVerb
	fmt.Println(gray("Logger Global Verbosity: "), lightRed(LogLevelToString(globalVerb)))

	for _, lgr := range l.Config.Loggers {
		appenderName := lgr.AppenderName
		writerName := lgr.WriterName
		fmt.Println(gray("Looking for Appender: "), lightRed(appenderName))
		fmt.Println(gray("Looking for Writer: "), lightRed(writerName))
		appender, okAppender := appendersMap[appenderName]
		writer, okWriter := writersMap[writerName]
		if !okAppender {
			//Missing appender
			warnings = append(warnings, fmt.Sprintf("Missing Appender <%s> in Log Config", appenderName))
			fmt.Println(red("Skipping logger: Appender : "), red(appenderName), red(" - Writer : "), red(writerName))
		} else {
			if !okWriter {
				//Missing writer
				warnings = append(warnings, fmt.Sprintf("Missing Writer <%s> in Log Config", writerName))
				fmt.Println(red("Skipping logger: Appender : "), red(appenderName), red(" - Writer : "), red(writerName))
			} else {
				//All is fine
				fmt.Println(gray("Logger/Appender state: "), green("ok!"))
				filterLen := len(lgr.Filters)
				fmt.Println(gray("Found package filers: "), lightRed(strconv.Itoa(filterLen)))
				if filterLen > 0 {
					for _, fltr := range lgr.Filters {
						pkgName := fltr.PackageName
						pkgVerbosity, pkgErr := StringToLogLevel(fltr.Verbosity)
						if pkgErr != nil {
							pkgVerbosity = DEFAULT_VERBOSITY
						}
						warnings = computePackageData(l, writer, appender, pkgName, pkgVerbosity, warnings)
					}

				} else {
					pkgName := "*"
					pkgVerbosity := globalVerb
					warnings = computePackageData(l, writer, appender, pkgName, pkgVerbosity, warnings)
				}

			}
		}
	}
	fmt.Println(gray("Configuration data loaded : "), green("Done!!"))
	if len(warnings) == 0 {
		fmt.Println(gray("Warnings : "), green("None!!"))
	} else {
		fmt.Println(gray("Warnings : "), red(strconv.Itoa(len(warnings))))
		for _, warn := range warnings {
			fmt.Println(yellow(warn))
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

const defaultDateFormat string = "2006-01-02 15:04:05.000"

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
		if testChan == nil {
			//			writerEncoding, _ := common.StreamInOutFormatToString(common.PlainTextFormat)
			//			writerType, _ := common.WriterTypeToString(common.StdOutWriter)
			verb, _ := LogLevelToString(verbosity)
			_loggerEng.Start(LogConfig{
				Appenders: []LogAppender{
					LogAppender{
						AppenderName: "StdOutAppender",
						DateFormat:   defaultDateFormat,
						Verbosity:    verb,
					},
				},
				Writers: []LogWriter{
					LogWriter{
						WriterName:     "StdOutWriter",
						WriterType:     common.StdOutWriter,
						WriterEncoding: common.PlainTextFormat,
						Destination:    "",
					},
				},
				LoggerName: "StdOutLogger",
				Loggers: []LoggerInfo{
					LoggerInfo{
						AppenderName: "StdOutAppender",
						WriterName:   "StdOutWriter",
						Filters:      []LogFilter{},
					},
				},
			})
		}
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
		if testChan == nil {
			_loggerEng.Start(config)
		}
	}
}
