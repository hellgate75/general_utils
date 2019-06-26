package common

import (
	"fmt"
	"github.com/hellgate75/general_utils/errors"
)

// Type that describes Server Log Level
type LogLevel int

const (
	// NO_LOG Log Level - Not Logging
	NO_LOG LogLevel = iota
	// TRACE Log Level - High Verbosity
	TRACE
	// DEBUG Log Level - Development Verbosity
	DEBUG
	// INFO Log Level - Standard Verbosity
	INFO
	// WARNING Log Level - Advice and Notifications Verbosity
	WARNING
	// ERROR Log Level - Application Errors Only Verbosity
	ERROR
	// FATAL Log Level - System Crash Verbosity
	FATAL
)

type ServerLogger interface {
	// Starts Logger
	// Returns:
	//    error Any error that can occurs during computation
	Open() error
	// Stops Logger
	// Returns:
	//    error Any error that can occurs during computation
	Close() error
	// Add Outbound channel
	// Parameters:
	//    channel (*chan interface{}) Output channel pointer for server logging activities to be added
	// Returns:
	//    error Any error that can occurs during computation
	AddOutChannel(channel *chan interface{}) error
	// Log an interface in the logging system
	// Parameters:
	//    logLevel (LogLevel) Capping Log Level
	//    message (interface{}) Message to be logged
	// Returns:
	//    error Any error that can occurs during computation
	Log(logLevel LogLevel, message interface{}) error
}

type __serverLogger struct {
	LogOutChanList []*chan interface{}
	Initialized    bool
	VerbosityLevel LogLevel
}

func (logManager *__serverLogger) Open() error {
	var err error
	defer func() {
		itf := recover()
		if errors.IsError(itf) {
			err = itf.(error)
		} else {
			err = errors.New(fmt.Sprintf("%v", itf))
		}
		if logger != nil {
			logger.Error(err)
		}
	}()
	logManager.LogOutChanList = make([]*chan interface{}, 0)
	logManager.Initialized = true
	return err
}

func (logManager *__serverLogger) Close() error {
	var err error
	defer func() {
		itf := recover()
		if errors.IsError(itf) {
			err = itf.(error)
		} else {
			err = errors.New(fmt.Sprintf("%v", itf))
		}
		if logger != nil {
			logger.Error(err)
		}
	}()
	for _, v := range logManager.LogOutChanList {
		close(*v)
	}
	logManager.LogOutChanList = make([]*chan interface{}, 0)
	logManager.Initialized = false
	return err
}

func (logManager *__serverLogger) AddOutChannel(channel *chan interface{}) error {
	if !logManager.Initialized {
		return errors.New("Unable to Add an Outbound channel to the Logger if it's not started")
	}
	logManager.LogOutChanList = append(logManager.LogOutChanList, channel)
	return nil
}

func (logManager *__serverLogger) Log(logLevel LogLevel, message interface{}) error {
	if !logManager.Initialized {
		return errors.New("Unable to Log Message to Outbound channels since the Logger is not started")
	}
	if len(logManager.LogOutChanList) == 0 {
		return errors.New("Unable to Log Message to Outbound channels since the Logger has no Outbound channels")
	}
	if logLevel >= logManager.VerbosityLevel {
		for _, ch := range logManager.LogOutChanList {
			*ch <- message
		}
	} else {
		return errors.New(fmt.Sprintf("Unable to Log Message to Outbound channels since the Logger has lower verbosity -> Given : <%v> that is less than Expected min verbosity: <%v> ", logLevel, logManager.VerbosityLevel))
	}
	return nil
}

func NewServerLogger(verbosity LogLevel) ServerLogger {
	return &__serverLogger{
		VerbosityLevel: verbosity,
	}
}
