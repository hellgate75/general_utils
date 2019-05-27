package log

import (
	//	"encoding/xml"
	"fmt"
	"github.com/hellgate75/general_utils/errors"
	"strings"
)

// Type WriterType describe any Writer Option in the cofiguration
type WriterType string

const (
	// Stadard Output Writer Type
	StdOutWriter WriterType = "stdOut"
	// File Output Writer Type
	FileWriter WriterType = "file"
	// URL Output Writer Type
	UrlWriter WriterType = "url"
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
		return StdOutWriter, nil
	case "FILE":
		return FileWriter, nil
	case "URL":
		return UrlWriter, nil
	}
	return "", errors.New(fmt.Sprintf("logger::WriterTypeFromString::error : Invalid WriterType '%s'", text))
}

// Log File Appeder Element, describes Writer default log verbosity
type LogAppender struct {
	AppenderName string `yaml:"appenderName" json:"appenderName" xml:"appender-name"`
	Verbosity    string `yaml:"defaulVerbosity" json:"defaulVerbosity,omitempty" xml:"defaul-verbosity"`
	DateFormat   string `yaml:"dateFormat" json:"dateFormat,omitempty" xml:"date-format"`
}

// Log File Write, describes support where to write the logs
type LogWriter struct {
	WriterName  string     `yaml:"writerName" json:"writerName" xml:"writer-name"`
	WriterType  WriterType `yaml:"writerType" json:"writerType" xml:"writer-type"`
	Destination string     `yaml:"destination" json:"destination,omitempty" xml:"destination"`
}

//Log File filter, describes filters for packages and wildcat represent default verbosity for any package
type LogFilter struct {
	PackageName string `yaml:"packageName" json:"packageName" xml:"package-name"`
	Verbosity   string `yaml:"verbosity" json:"verbosity" xml:"verbosity"`
}

// Logger Configuration represents the defiition of a logger poiter
type LoggerInfo struct {
	AppenderName string      `yaml:"appenderName" json:"appenderName,omitempty" xml:"appender-name"`
	WriterName   string      `yaml:"writerName" json:"writerName" xml:"writer-name"`
	Filters      []LogFilter `yaml:"filter" json:"filters,omitempty" xml:"filter"`
}

// Log Main Configuration represents the file layout, and actions recognition
type LogConfig struct {
	LoggerName string        `yaml:"loggerName" json:"loggerName" xml:"logger-name"`
	Verbosity  string        `yaml:"rootVerbosity" json:"rootVerbosity" xml:"rootVerbosity"`
	Appenders  []LogAppender `yaml:"appender" json:"appenders" xml:"appender"`
	Writers    []LogWriter   `yaml:"writer" json:"writers" xml:"writer"`
	Loggers    []LoggerInfo  `yaml:"logger" json:"loggers" xml:"logger"`
}

//	XMLName    xml.Name      `xml:"http://www.git/XMLSchema/standard/2019 logCofig"`
