package parser

import (
	"errors"
	"fmt"
	"github.com/hellgate75/general_utils/common"
	"time"
)

type LogStreamType int

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
}

type _logFileStreamStruct struct {
	file      string
	format    common.StreamInOutFormat
	policy    RotatePolicy
	converter ConverterFunc
}

type _logUrlStreamStruct struct {
	url       string
	format    common.StreamInOutFormat
	converter ConverterFunc
}

func (l *_logStdOutStruct) Open() error {
	var err error = nil
	defer func() {
		r := recover()
		err = r.(error)
	}()
	l.converter = getConverterByStreamInOutFormat(l._logFormat)
	return err
}

func (l *_logStdOutStruct) Close() error {
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
	fmt.Println(string(body))
	return nil
}

func New(logType LogStreamType, format common.StreamInOutFormat) (LogStream, error) {

	switch logType {
	case StdOutStreamType:
		return &_logStdOutStruct{
			_logFormat: format,
		}, nil
	default:
		return nil, errors.New("log::parser::error : Unknown or Not Implemented Log Stream Type")
	}
}
