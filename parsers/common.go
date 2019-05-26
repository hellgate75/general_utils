package parsers

import (
	"fmt"
	"github.com/hellgate75/general_utils/errors"
	"github.com/hellgate75/general_utils/log"
)

var logger log.Logger = log.GetLogger()

type Type interface{}

type Encoding int

const (
	JSON Encoding = 1
	XML  Encoding = 2
	YAML Encoding = 3
)

type XmlParserStruct struct {
}

type JsonParserStruct struct {
}

type YamlParserStruct struct {
}

type Parser interface {
	DeserializeFromFile(filePath string, mask *Type) error

	DeserializeFromBytes(bytes []byte, mask *Type) error

	SerializeToFile(filePath string, mask *Type) error

	SerializeToBytes(mask *Type) ([]byte, error)
}

func New(enc Encoding) (Parser, error) {
	switch enc {
	case JSON:
		return &JsonParserStruct{}, nil
	case XML:
		return &XmlParserStruct{}, nil
	case YAML:
		return &YamlParserStruct{}, nil
	}
	return nil, errors.New(fmt.Sprintf("Uknown Parser : %v", enc))
}
