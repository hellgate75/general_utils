package parser

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"encoding/gob"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/hellgate75/general_utils/common"
	"github.com/hellgate75/general_utils/errors"
	"gopkg.in/yaml.v2"
)

//Represent the converion function marked as follow:
//    func(interface{}) ([]byte, error)
type ConverterFunc func(interface{}) ([]byte, error)

func getConverterByStreamInOutFormat(format common.StreamInOutFormat) ConverterFunc {
	switch format {
	case common.PlainTextFormat:
		return _convertToPlainText
	case common.JsonFormat:
		return _convertToJson
	case common.YamlFormat:
		return _convertToYaml
	case common.XmlFormat:
		return _convertToXml
	case common.EncryptedFormat:
		return _convertToBase64
	case common.BinaryFormat:
		return _convertToBinary
	case common.GoStructFormat:
		return _convertToGoFormat
	default:
		return _convertToPlainText
	}
}

func _convertToPlainText(text interface{}) ([]byte, error) {
	return []byte(fmt.Sprintf("%v", text)), nil
}

func _convertToJson(text interface{}) ([]byte, error) {
	return json.Marshal(text)
}

func _convertToXml(text interface{}) ([]byte, error) {
	return xml.Marshal(text)
}

func _convertToYaml(text interface{}) ([]byte, error) {
	return yaml.Marshal(text)
}

func _convertToBinary(text interface{}) ([]byte, error) {
	writer := NewLocalWriter()
	err := binary.Write(writer, binary.BigEndian, text)
	if err != nil {
		return []byte{}, err
	}
	return writer.GetBytes()
}

func _convertToBase64(text interface{}) ([]byte, error) {
	writer := NewLocalWriter()
	b64Writer := base64.NewEncoder(base64.RawStdEncoding, writer)
	_, err := b64Writer.Write([]byte(fmt.Sprintf("%v", text)))
	if err != nil {
		return []byte{}, err
	}
	return writer.GetBytes()
}

func _convertToGoFormat(text interface{}) ([]byte, error) {
	var buff bytes.Buffer
	enc := gob.NewEncoder(&buff)
	err := enc.Encode(text)
	if err != nil {
		return []byte{}, err
	}
	return buff.Bytes(), nil
}

func NewLocalWriter() LocalWriter {
	return &_localWriterStruct{
		_buff: bytes.NewBuffer([]byte{}),
	}
}

type LocalWriter interface {
	Write(p []byte) (n int, err error)
	GetBytes() (b []byte, err error)
}

type _localWriterStruct struct {
	_buff *bytes.Buffer
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
