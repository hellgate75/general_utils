package parser

import (
	"errors"
	"fmt"
	"github.com/hellgate75/general_utils/common"
	errs "github.com/hellgate75/general_utils/errors"
	"runtime"
	"time"
)

type LogStreamType int

var testFlag bool = false
var testChan chan interface{}

const (
	StdOutStreamType LogStreamType = 101
	FileStreamType   LogStreamType = 102
	UrlStreamType    LogStreamType = 103
)

func WriterTypeToLogStreamType(wTp common.WriterType) (LogStreamType, error) {
	switch wTp {
	case common.StdOutWriter:
		return StdOutStreamType, nil
	case common.FileWriter:
		return FileStreamType, nil
	case common.UrlWriter:
		return UrlStreamType, nil
	default:
		return 0, errors.New("Unkown Writer Type")
	}
}

type LogStream interface {
	Open() error
	Close() error
	Write(interface{}) error
}

type RotatePolicy struct {
	size      int
	unit      common.StorageUnit
	frequency int
	timeUnit  time.Duration
	history   int
}

type _logStdOutStruct struct {
	_logFormat common.StreamInOutFormat
	converter  ConverterFunc
	open       bool
}

type _logFileStreamStruct struct {
	file      string
	format    common.StreamInOutFormat
	policy    RotatePolicy
	converter ConverterFunc
	open      bool
}

type _logUrlStreamStruct struct {
	url       string
	format    common.StreamInOutFormat
	converter ConverterFunc
	open      bool
}

func (l *_logStdOutStruct) Open() error {
	var err error
	defer func() {
		r := recover()
		if errs.IsError(r) {
			err = r.(error)
		} else {
			err = errors.New("Unknown Error Occured")
		}
	}()
	l.converter = getConverterByStreamInOutFormat(l._logFormat)
	l.open = (err == nil)
	return err
}

func (l *_logStdOutStruct) Close() error {
	if !l.open {
		return nil
	}
	l.open = false
	return nil
}

func (l *_logStdOutStruct) Write(data interface{}) error {
	if l.converter == nil {
		return errors.New("Start Log Stream before write content")
	}
	body, err := l.converter(data)
	if err != nil {
		return err
	}
	if testFlag && testChan != nil {
		testChan <- string(body)
	} else {
		fmt.Println(string(body))

	}
	return nil
}

func New(logType LogStreamType, format common.StreamInOutFormat) (LogStream, error) {
	//TODO: Complete File and Url Stream Implementations and place into this method

	switch logType {
	case StdOutStreamType:
		var los _logStdOutStruct = _logStdOutStruct{
			_logFormat: format,
		}
		runtime.SetFinalizer(&los, func(los *_logStdOutStruct) {
			los.Close()
		})
		return &los, nil
	default:
		return nil, errors.New("log::parser::error : Unknown or Not Implemented Log Stream Type")
	}
}
