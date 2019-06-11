package log

import (
	"errors"
	"fmt"
	"github.com/hellgate75/general_utils/common"
	"strings"
	"testing"
	"time"
)

func _testLogVerbosityFromText(t *testing.T, text string, expected LogLevel) {
	logLev, err := StringToLogLevel(text)
	if err != nil {
		t.Fatal(fmt.Sprintf("Error converting string %s to LogLevel - error : %s", text, err.Error()))
	}
	if logLev != expected {
		t.Fatal(fmt.Sprintf("Error converting string to LogLevel - Expected <%v> Given <%v> : ", expected, logLev))
	}
}

func TestStringToLogLevel(t *testing.T) {
	_testLogVerbosityFromText(t, "DEBUG", DEBUG)
	_testLogVerbosityFromText(t, "INFO", INFO)
	_testLogVerbosityFromText(t, "WARN", WARN)
	_testLogVerbosityFromText(t, "ERROR", ERROR)
	_testLogVerbosityFromText(t, "FATAL", FATAL)
	_testLogVerbosityFromText(t, "NOLOG", NOLOG)
	_testLogVerbosityFromText(t, "DEFAULT", DEFAULT_VERBOSITY)
	_, err := StringToLogLevel("")
	if err == nil {
		t.Fatal("No error arisen converting empty string to LogLevel")
	}
}

func _testLogVerbosityToText(t *testing.T, expected string, logLevel LogLevel) {
	logText, err := LogLevelToString(logLevel)
	if err != nil {
		t.Fatal(fmt.Sprintf("Error converting log level %v to LogLevel - error : %s", logLevel, err.Error()))
	}
	if logText != expected {
		t.Fatal(fmt.Sprintf("Error converting LogLevel to String - Expected <%v> Given <%v> : ", expected, logText))
	}
}

func TestLogLevelToString(t *testing.T) {
	_testLogVerbosityToText(t, "DEBUG", DEBUG)
	_testLogVerbosityToText(t, "INFO ", INFO)
	_testLogVerbosityToText(t, "WARN ", WARN)
	_testLogVerbosityToText(t, "ERROR", ERROR)
	_testLogVerbosityToText(t, "FATAL", FATAL)
	_testLogVerbosityToText(t, "NOLOG", NOLOG)
	_testLogVerbosityToText(t, DEFAULT_VERBOSITY_TEXT, DEFAULT_VERBOSITY)
	_, err := LogLevelToString(LogLevel(0))
	if err == nil {
		t.Fatal("No error arisen converting out of bound LogLevel to Text")
	}
}

func TestInitStaticLoggerEngine(t *testing.T) {
	InitStaticLoggerEngine(INFO)
	LogChanEnabled = true
	LogOutChan = make(chan interface{})
	defer func() {
		_loggerEng.Stop()
		LogChanEnabled = false
		close(LogOutChan)
	}()
	var message interface{} = "I'm logging"
	go func() {
		time.Sleep(350 * time.Millisecond)
		_loggerEng.Log(INFO, message, nil, "log")
	}()
	select {
	case msg := <-LogOutChan:
		var text string = fmt.Sprintf("%v", msg)
		var expected string = fmt.Sprintf("%v", message)
		if strings.Index(text, expected) < 0 {
			t.Fatal(fmt.Sprintf("Error logging - Expected contained <%v> Given <%v> : ", message, msg))

		}
	}
}

func _createLogConfig() LogConfig {
	return LogConfig{
		Appenders: []LogAppender{
			LogAppender{
				AppenderName: "StdOutAppender",
				DateFormat:   defaultDateFormat,
				Verbosity:    "INFO",
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
	}
}

func TestInitStaticLoggerEngineFromConfig(t *testing.T) {
	var conf LogConfig = _createLogConfig()
	InitStaticLoggerEngineFromConfig(conf)
	LogChanEnabled = true
	LogOutChan = make(chan interface{})
	defer func() {
		_loggerEng.Stop()
		LogChanEnabled = false
		close(LogOutChan)
	}()
	var message interface{} = "I'm logging"
	go func() {
		time.Sleep(350 * time.Millisecond)
		_loggerEng.Log(INFO, message, nil, "log")
	}()
	select {
	case msg := <-LogOutChan:
		var text string = fmt.Sprintf("%v", msg)
		var expected string = fmt.Sprintf("%v", message)
		if strings.Index(text, expected) < 0 {
			t.Fatal(fmt.Sprintf("Error logging - Expected contained <%v> Given <%v> : ", message, msg))

		}
	}
}

func TestNewLogger(t *testing.T) {
	var conf LogConfig = _createLogConfig()
	InitStaticLoggerEngineFromConfig(conf)
	InitStaticLoggerEngineFromConfig(conf)
	LogChanEnabled = true
	LogOutChan = make(chan interface{})
	logger, err := New("log")
	if err != nil {
		t.Fatal(fmt.Sprintf("Error creating new Logger - error : %s", err.Error()))

	}
	defer func() {
		_loggerEng.Stop()
		LogChanEnabled = false
		close(LogOutChan)
	}()
	var message interface{} = "I'm logging"
	go func() {
		time.Sleep(350 * time.Millisecond)
		logger.Info(message)
	}()
	select {
	case msg := <-LogOutChan:
		var text string = fmt.Sprintf("%v", msg)
		var expected string = fmt.Sprintf("%v", message)
		if strings.Index(text, expected) < 0 {
			t.Fatal(fmt.Sprintf("Error logging - Expected contained <%v> Given <%v> : ", message, msg))

		}
	}
}

func _testSingleLogLevel(t *testing.T, level string, asString bool) {
	var conf LogConfig = _createLogConfig()
	conf.Verbosity = strings.ToUpper(level)
	ResetStaticLoggerEngine()
	logLevel, _ := StringToLogLevel(level)
	InitStaticLoggerEngine(logLevel)
	InitStaticLoggerEngine(logLevel)
	logger, err := New("log")
	if err != nil {
		t.Fatal(fmt.Sprintf("Error creating new Logger - error : %s", err.Error()))

	}

	var message interface{} = "I'm logging"
	go func() {
		time.Sleep(500 * time.Millisecond)
		switch strings.ToUpper(level) {
		case "DEBUG":
			logger.Debug(message)
		case "INFO":
			logger.Info(message)
		case "WARN":
			logger.Warning(message)
		case "ERROR":
			if asString {
				logger.ErrorS(message)

			} else {
				err := errors.New(fmt.Sprintf("%v", message))
				fmt.Println(fmt.Sprintf("Sending error : %v", err))
				logger.Error(err)
			}
		case "FATAL":
			if asString {
				logger.FatalS(message)

			} else {
				err := errors.New(fmt.Sprintf("%v", message))
				fmt.Println(fmt.Sprintf("Sending fatal : %v", err))
				logger.Fatal(err)
			}
		default:
			t.Fatal(fmt.Sprintf("Unwanted test case with log level : %s!!", level))

		}
	}()
	select {
	case msg := <-LogOutChan:
		var text string = fmt.Sprintf("%v", msg)
		var expected string = fmt.Sprintf("%v", message)
		if strings.Index(text, expected) < 0 {
			t.Fatal(fmt.Sprintf("Error logging Level: %s - Expected contained <%v> Given <%v> : ", level, message, msg))
		}
	case <-time.After(6 * time.Second):
		t.Fatal(fmt.Sprintf("Testing level %s timeout waiting log reached!!", level))
	}
}

func TestNewLoggerDebugLevel(t *testing.T) {
	LogChanEnabled = true
	LogOutChan = make(chan interface{})
	defer func() {
		if _loggerEng != nil {
			if _loggerEng.IsRunning() {
				_loggerEng.Stop()
			}
			if _loggerEng.IsSimple() {
				fmt.Println("Stopping simple engine!!")
			} else {
				fmt.Println("Stopping complex engine!!")
			}
		}
		LogChanEnabled = false
		close(LogOutChan)
		ResetStaticLoggerEngine()
	}()
	_testSingleLogLevel(t, "DEBUG", true)
}
func TestNewLoggerInfoLevel(t *testing.T) {
	LogChanEnabled = true
	LogOutChan = make(chan interface{})
	defer func() {
		if _loggerEng != nil {
			if _loggerEng.IsRunning() {
				_loggerEng.Stop()
			}
			if _loggerEng.IsSimple() {
				fmt.Println("Stopping simple engine!!")
			} else {
				fmt.Println("Stopping complex engine!!")
			}
		}
		LogChanEnabled = false
		close(LogOutChan)
		ResetStaticLoggerEngine()
	}()
	_testSingleLogLevel(t, "INFO", true)
}
func TestNewLoggerWarnLevel(t *testing.T) {
	LogChanEnabled = true
	LogOutChan = make(chan interface{})
	defer func() {
		if _loggerEng != nil {
			if _loggerEng.IsRunning() {
				_loggerEng.Stop()
			}
			if _loggerEng.IsSimple() {
				fmt.Println("Stopping simple engine!!")
			} else {
				fmt.Println("Stopping complex engine!!")
			}
		}
		LogChanEnabled = false
		close(LogOutChan)
	}()
	_testSingleLogLevel(t, "WARN", true)
}

func TestNewLoggerErrorLevelAsString(t *testing.T) {
	LogChanEnabled = true
	LogOutChan = make(chan interface{})
	defer func() {
		if _loggerEng != nil {
			if _loggerEng.IsRunning() {
				_loggerEng.Stop()
			}
			if _loggerEng.IsSimple() {
				fmt.Println("Stopping simple engine!!")
			} else {
				fmt.Println("Stopping complex engine!!")
			}
		}
		LogChanEnabled = false
		close(LogOutChan)
		ResetStaticLoggerEngine()
	}()
	_testSingleLogLevel(t, "ERROR", true)
}

func TestNewLoggerErrorLevelAsErrorObject(t *testing.T) {
	LogChanEnabled = true
	LogOutChan = make(chan interface{})
	defer func() {
		if _loggerEng != nil {
			if _loggerEng.IsRunning() {
				_loggerEng.Stop()
			}
			if _loggerEng.IsSimple() {
				fmt.Println("Stopping simple engine!!")
			} else {
				fmt.Println("Stopping complex engine!!")
			}
		}
		LogChanEnabled = false
		close(LogOutChan)
		ResetStaticLoggerEngine()
	}()
	_testSingleLogLevel(t, "ERROR", false)
}

func TestNewLoggerFatalLevelAsString(t *testing.T) {
	LogChanEnabled = true
	LogOutChan = make(chan interface{})
	defer func() {
		if _loggerEng != nil {
			if _loggerEng.IsRunning() {
				_loggerEng.Stop()
			}
			if _loggerEng.IsSimple() {
				fmt.Println("Stopping simple engine!!")
			} else {
				fmt.Println("Stopping complex engine!!")
			}
		}
		LogChanEnabled = false
		close(LogOutChan)
		ResetStaticLoggerEngine()
	}()
	_testSingleLogLevel(t, "FATAL", true)
}

func TestNewLoggerFatalLevelAsErrorObject(t *testing.T) {
	LogChanEnabled = true
	LogOutChan = make(chan interface{})
	defer func() {
		if _loggerEng != nil {
			if _loggerEng.IsRunning() {
				_loggerEng.Stop()
			}
			if _loggerEng.IsSimple() {
				fmt.Println("Stopping simple engine!!")
			} else {
				fmt.Println("Stopping complex engine!!")
			}
		}
		LogChanEnabled = false
		close(LogOutChan)
		ResetStaticLoggerEngine()
	}()
	_testSingleLogLevel(t, "FATAL", false)
}

func TestNewLoggerUnappyPath(t *testing.T) {
	ResetStaticLoggerEngine()
	_, err := New("log")
	if err == nil {
		t.Fatal("No error arisen creating Logger with uninitialized Logger Engine!!")
	}

}

func TestGetEngineFromConfig(t *testing.T) {
	var conf LogConfig = _createLogConfig()
	conf.Verbosity = ""
	_, err := getEngineFromConfig(conf)
	if err == nil {
		t.Fatal("No error arisen creating Logger engine by LogConfig with wrong Log Verbosity!!")
	}
	// Here we use default verbosity
	conf.Verbosity = "I-DUNNO"
	_, err = getEngineFromConfig(conf)
	if err != nil {
		t.Fatal(fmt.Sprintf("Error creating new Logger engine by LogConfig - error : %s", err.Error()))

	}
	conf.Verbosity = "INFO"
	_, err = getEngineFromConfig(conf)
	if err != nil {
		t.Fatal(fmt.Sprintf("Error creating new Logger engine by LogConfig - error : %s", err.Error()))

	}

}
