package parsers // import "github.com/hellgate75/general_utils/parsers"


FUNCTIONS

func InitLogger()

TYPES

type Encoding int
    Parser Ecoding Type

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
func EncodingFromString(text string) (Encoding, error)
    Transform text to representing Encoding element.

    Parameters:

    text (string) Text to parse

    Returns:

    parsers.Encoding File Encoding type (JSON, XML, YAML)
    error Any suitable error risen during code execution

type LocalWriter interface {
	Write(p []byte) (n int, err error)
	GetBytes() (b []byte, err error)
}

func NewLocalWriter() LocalWriter
func NewLocalWriterCustom(buffer bytes.Buffer) LocalWriter
type Parser interface {
	// Provides Deserialization from a File.
	//
	// Parameters:
	//   filePath (string) File full path
	//   out (interface{}) Pointer to Element to parse
	//
	// Returns:
	// out (interface{}) Element to parse
	// error Any suitable error risen during code execution
	DeserializeFromFile(filePath string, out interface{}) error

	// Provides Deserialization from a byte array.
	//
	// Parameters:
	//   bytes ([]byte) Bytes to be parsed
	//   out (interface{}) Pointer to Element to parse
	//
	// Returns:
	// error Any suitable error risen during code execution
	DeserializeFromBytes(bytes []byte, out interface{}) error

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
    Define Generic Parser Features

func New(enc Encoding) (Parser, error)
    Creates new Parser.

    Parameters:

    enc (parser.Encoding) Parser encoding type

    Returns: parser.Parser Required parser or Nil if not available error Any
    suitable error risen during code execution

