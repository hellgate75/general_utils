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
}

type _logFileStreamStruct struct {
	file   string
	format common.StreamInOutFormat
	policy RotatePolicy
}

type _logUrlStreamStruct struct {
	url    string
	format common.StreamInOutFormat
}

func (l *_logStdOutStruct) Open() error {
	return nil
}

func (l *_logStdOutStruct) Close() error {
	return nil
}

func (l *_logStdOutStruct) Write(data interface{}) error {
	var function ConverterFunc = getConverterByStreamInOutFormat(l._logFormat)
	body, err := function(data)
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
