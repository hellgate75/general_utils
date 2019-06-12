package parsers

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/hellgate75/general_utils/common"
	"github.com/hellgate75/general_utils/log"
	"strings"
)

var logger log.Logger

func InitLogger() {
	currentLogger, err := log.New("parsers")
	if err != nil {
		panic(err.Error())
	}
	logger = currentLogger
}

// Parser Ecoding Type
type Encoding int

const (
	// JSON Encoding type
	JSON Encoding = 1
	// XML Encoding type
	XML Encoding = 2
	// YAML Encoding type
	YAML Encoding = 3
	// GO LANGUAGE (Gob) Encoding type
	GOLANG Encoding = 4
	// BASE64 Encoding type
	BASE64 Encoding = 5
	// BINARY Encoding type
	BINARY Encoding = 6
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
	case "GOF":
		return GOLANG, nil
	case "B64":
		return BASE64, nil
	case "BIN":
		return BINARY, nil
	}
	return JSON, nil
}

type xmlParserStruct struct {
}

type jsonParserStruct struct {
}

type yamlParserStruct struct {
}

type gobParserStruct struct {
}
type base64ParserStruct struct {
	internalParser Parser
}
type binaryParserStruct struct {
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

	// Provides implemented encoding format.
	//
	// Returns:
	// Encoding Encoding type
	GetEncoding() Encoding
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
	case GOLANG:
		return &gobParserStruct{}, nil
	case BASE64:
		parser, _ := New(JSON)
		return &base64ParserStruct{
			internalParser: parser,
		}, nil
	case BINARY:
		return &binaryParserStruct{}, nil
	default:
		return nil, errors.New(fmt.Sprintf("Uknown encoding format %v!!", enc))
	}
	return nil, errors.New(fmt.Sprintf("Uknown Parser : %v", enc))
}

type LocalWriter interface {
	Write(p []byte) (n int, err error)
	GetBytes() (b []byte, err error)
}

type _localWriterStruct struct {
	_buff *bytes.Buffer
}

func NewLocalWriterCustom(buffer bytes.Buffer) LocalWriter {
	return &_localWriterStruct{
		_buff: &buffer,
	}
}

func NewLocalWriter() LocalWriter {
	var buff bytes.Buffer
	return &_localWriterStruct{
		_buff: &buff,
	}
}

func (lw *_localWriterStruct) Write(p []byte) (n int, err error) {
	if p == nil {
		return 0, errors.New("log::parser::LocalWriter : Undefined input/output byte array")
	}
	length := len(p)
	if length == 0 {
		return 0, errors.New("log::parser::LocalWriter : Zero-length input/output byte array")
	}
	if lw._buff == nil {
		lw._buff = bytes.NewBuffer([]byte{})
	}
	return lw._buff.Write(p)
}

func (lw *_localWriterStruct) GetBytes() (b []byte, err error) {
	if lw._buff == nil {
		return []byte{}, errors.New("log::parser::LocalWriter : no bytes added to the Writer")
	}
	return lw._buff.Bytes(), nil
}
