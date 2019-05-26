package parsers

import (
	"fmt"
	"github.com/hellgate75/general_utils/common"
	"github.com/hellgate75/general_utils/errors"
	"github.com/hellgate75/general_utils/log"
	"strings"
)

var logger log.Logger = log.GetLogger()

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

// Transform text to representing Encoding element.
//
// Parameters:
//   text (string) Text to parse
//
// Returns:
//   parsers.Encoding File Encoding type (JSON, XML, YAML)
//   error Any suitable error risen during code execution
func EncodingFromString(text string) (Encoding, error) {
	if strings.TrimSpace(text) == "" {
		return 0, errors.New("parsers::EncodingFromString::error : Empty input string")
	}
	value := strings.ToUpper(text)
	switch value {
	case "JSON":
		return JSON, nil
	case "XML":
		return XML, nil
	case "YAML":
		return YAML, nil
	}
	return JSON, nil
}

type xmlParserStruct struct {
}

type jsonParserStruct struct {
}

type yamlParserStruct struct {
}

// Define Generic Parser Features
type Parser interface {
	// Provides Deserialization from a File.
	//
	// Parameters:
	//   filePath (string) File full path
	//   mask (common.Type) Generic Element to be deserialized
	//
	// Returns:
	// error Any suitable error risen during code execution
	DeserializeFromFile(filePath string, mask common.Type) error

	// Provides Deserialization from a byte array.
	//
	// Parameters:
	//   bytes ([]byte) Bytes to be parsed
	//   mask (common.Type) Generic Element to be deserialized
	//
	// Returns:
	// error Any suitable error risen during code execution
	DeserializeFromBytes(bytes []byte, mask common.Type) error

	// Provides Serialization to a File.
	//
	// Parameters:
	//   filePath (string) File full path
	//   mask (common.Type) Generic Element to be serialized
	//
	// Returns:
	// error Any suitable error risen during code execution
	SerializeToFile(filePath string, mask common.Type) error

	// Provides Serialization from a byte array.
	//
	// Parameters:
	//   mask common.Type) Generic Element to be serialized
	//
	// Returns:
	// []byte Object serialization Byte array
	// error Any suitable error risen during code execution
	SerializeToBytes(mask common.Type) ([]byte, error)
}

// Creates new Parser.
//
// Parameters:
//   enc (parser.Encoding) Parser encoding type
//
// Returns:
// parser.Parser Required parser or Nil if not available
// error Any suitable error risen during code execution
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
