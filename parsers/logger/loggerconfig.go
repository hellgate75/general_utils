package parsers

import (
	"encoding/xml"
	"fmt"
	"github.com/hellgate75/general_utils/errors"
	"strings"
)

type WriterType string

const (
	stdOut WriterType = "stdOut"
	file   WriterType = "file"
	url    WriterType = "url"
)

func WriterTypeFromString(text string) (WriterType, error) {
	if strings.TrimSpace(text) == "" {
		return "", errors.New("logger::WriterTypeFromString::error : Empty input string")
	}
	value := strings.ToUpper(text)
	switch value {
	case "STDOUT":
		return stdOut, nil
	case "FILE":
		return file, nil
	case "URL":
		return url, nil
	}
	return "", errors.New(fmt.Sprintf("logger::WriterTypeFromString::error : Invalid WriterType '%s'", text))
}

type LogAppender struct {
	appenderName string `yaml:"appenderName" json:"appenderName" xml:"appender-name"`
	verbosity    string `yaml:"defaulVerbosity" json:"defaulVerbosity,omitempty" xml:"defaul-verbosity"`
}

type LogWriter struct {
	writerName  string     `yaml:"writerName" json:"writerName" xml:"writer-name"`
	writerType  WriterType `yaml:"writerType" json:"writerType" xml:"writer-type"`
	destination string     `yaml:"destination" json:"destination,omitempty" xml:"destination"`
}

type LogFilter struct {
	packageName string `yaml:"packageName" json:"packageName" xml:"package-name"`
	verbosity   string `yaml:"verbosity" json:"verbosity" xml:"verbosity"`
}

type Logger struct {
	appenderName string      `yaml:"appenderName" json:"appenderName" xml:"appender-name"`
	writerName   string      `yaml:"writerName" json:"writerName" xml:"writer-name"`
	filters      []LogFilter `yaml:"filter" json:"filters,omitempty" xml:"filter"`
}

type LogConfig struct {
	XMLName    xml.Name      `xml:"http://www.example.com/XMLSchema/standard/2019 logCofig"`
	loggerName string        `yaml:"loggerName" json:"loggerName" xml:"logger-name"`
	verbosity  string        `yaml:"rootVerbosity" json:"rootVerbosity" xml:"rootVerbosity"`
	appenders  []LogAppender `yaml:"appender" json:"appenders" xml:"appender"`
	writers    []LogWriter   `yaml:"writer" json:"writers" xml:"writer"`
	loggers    []Logger      `yaml:"logger" json:"loggers" xml:"logger"`
}
