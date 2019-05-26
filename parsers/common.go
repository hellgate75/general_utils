package parsers

import (
	"fmt"
	"github.com/hellgate75/general_utils/errors"
	"github.com/hellgate75/general_utils/log"
)

var logger log.Logger = log.GetLogger()

// Geeric Type
type Type interface{}

// Parser Ecoding Type
type Encoding int

const (
	// JSON Encoding type
	JSON Encoding = 1
	// XML Encoding type
	XML Encoding = 2
	// YAML Encoding type
	YAML Encoding = 3
)

type xmlParserStruct struct {
}

type jsonParserStruct struct {
}

type yamlParserStruct struct {
}

// Define Generic Parser Features
type Parser interface {
	// Provides read access to log configuration and parse from a File.
	//
	// Parameters:
	//   encoding (parsers.Encoding) File Encoding type (JSON, XML, YAML)
	//   filePath (string) File full path
	//
	// Returns:
	// error Any suitable error during code execution
	DeserializeFromFile(filePath string, mask Type) error

	DeserializeFromBytes(bytes []byte, mask Type) error

	SerializeToFile(filePath string, mask Type) error

	SerializeToBytes(mask Type) ([]byte, error)
}

func New(enc Encoding) (Parser, error) {
	switch enc {
	case JSON:
		return &jsonParserStruct{}, nil
	case XML:
		return &xmlParserStruct{}, nil
	case YAML:
		return &yamlParserStruct{}, nil
	}
	return nil, errors.New(fmt.Sprintf("Uknown Parser : %v", enc))
}
