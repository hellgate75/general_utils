package log

import (
	"encoding/xml"
	"fmt"
	"github.com/hellgate75/general_utils/errors"
	"strings"
)

// Type WriterType describe any Writer Option in the cofiguration
type WriterType string

const (
	// Stadard Output Writer Type
	stdOut WriterType = "stdOut"
	// File Output Writer Type
	file WriterType = "file"
	// URL Output Writer Type
	url WriterType = "url"
)

// Parse a String to Writer Type.
//
// Parameters:
//   text (string) Text to parse
//
// Returns:
// WriterType costant element
// error Any suitable error risen during code execution

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

// Log File Appeder Element, describes Writer default log verbosity
type LogAppender struct {
	appenderName string `yaml:"appenderName" json:"appenderName" xml:"appender-name"`
	verbosity    string `yaml:"defaulVerbosity" json:"defaulVerbosity,omitempty" xml:"defaul-verbosity"`
}

// Log File Write, describes support where to write the logs
type LogWriter struct {
	writerName  string     `yaml:"writerName" json:"writerName" xml:"writer-name"`
	writerType  WriterType `yaml:"writerType" json:"writerType" xml:"writer-type"`
	destination string     `yaml:"destination" json:"destination,omitempty" xml:"destination"`
}

//Log File filter, describes filters for packages and wildcat represent default verbosity for any package
type LogFilter struct {
	packageName string `yaml:"packageName" json:"packageName" xml:"package-name"`
	verbosity   string `yaml:"verbosity" json:"verbosity" xml:"verbosity"`
}

// Logger Configuration represents the defiition of a logger poiter
type LoggerInfo struct {
	appenderName string      `yaml:"appenderName" json:"appenderName,omitempty" xml:"appender-name"`
	writerName   string      `yaml:"writerName" json:"writerName" xml:"writer-name"`
	filters      []LogFilter `yaml:"filter" json:"filters,omitempty" xml:"filter"`
}

// Log Main Configuration represents the file layout, and actions recognition
type LogConfig struct {
	XMLName    xml.Name      `xml:"http://www.example.com/XMLSchema/standard/2019 logCofig"`
	loggerName string        `yaml:"loggerName" json:"loggerName" xml:"logger-name"`
	verbosity  string        `yaml:"rootVerbosity" json:"rootVerbosity" xml:"rootVerbosity"`
	appenders  []LogAppender `yaml:"appender" json:"appenders" xml:"appender"`
	writers    []LogWriter   `yaml:"writer" json:"writers" xml:"writer"`
	loggers    []LoggerInfo  `yaml:"logger" json:"loggers" xml:"logger"`
}
