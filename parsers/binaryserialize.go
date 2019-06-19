package parsers

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/hellgate75/general_utils/common"
	"github.com/hellgate75/general_utils/streams"
	"io/ioutil"
	"strconv"
)

func (this *binaryParserStruct) DeserializeFromFile(filePath string, out interface{}) error {
	var byteArray []byte
	var err error
	if byteArray, err = streams.LoadFileBytes(filePath); err == nil {
		buff := bytes.NewBuffer(byteArray)
		if err := binary.Read(buff, binary.BigEndian, out); err == nil {
			var length interface{} = "<null>"
			byteArray = buff.Bytes()
			if byteArray != nil && len(byteArray) > 0 {
				length = strconv.Itoa(len(byteArray))
			}
			if logger != nil {
				logger.Debug(fmt.Sprintf("Binary Parser :: Successful Deserialized bytes : %v", length))
			} else {
				return errors.New("Binary Parser :: Successful Deserialized null or empty set of bytes!!!")
			}
			return nil
		} else {
			if logger != nil {
				logger.Error(err)
			}
			return err
		}
	} else {
		return nil
	}
}

func (this *binaryParserStruct) DeserializeFromBytes(byteArray []byte, out interface{}) error {
	buff := bytes.NewBuffer(byteArray)
	if err := binary.Read(buff, binary.BigEndian, out); err == nil {
		var length interface{} = "<null>"
		byteArray = buff.Bytes()
		if byteArray != nil && len(byteArray) > 0 {
			length = strconv.Itoa(len(byteArray))
		} else {
			return errors.New("Binary Parser :: Successful Deserialized null or empty set of bytes!!!")
		}
		if logger != nil {
			logger.Debug(fmt.Sprintf("Binary Parser :: Successful Deserialized bytes : %v", length))
		}
		return nil
	} else {
		if logger != nil {
			logger.Error(err)
		}
		return err
	}
}

func (this *binaryParserStruct) SerializeToFile(filePath string, mask common.Type) error {
	var bytes []byte
	var err error
	if bytes, err = this.SerializeToBytes(mask); err == nil {
		if err = ioutil.WriteFile(filePath, bytes, 666); err == nil {
			if logger != nil {
				logger.Debug(fmt.Sprintf("Base64 Parser :: File written: %s", filePath))
			}
			return nil
		} else {
			if logger != nil {
				logger.Error(err)
			}
			return err
		}
	} else {
		if logger != nil {
			logger.Error(err)
		}
		return err
	}
}

func (this *binaryParserStruct) SerializeToBytes(mask common.Type) ([]byte, error) {
	var writer LocalWriter = NewLocalWriter()
	if err := binary.Write(writer, binary.BigEndian, mask); err == nil {
		byteArray, err1 := writer.GetBytes()
		return byteArray, err1
	} else {
		if logger != nil {
			logger.Error(err)
		}
		return nil, err
	}
}

func (this *binaryParserStruct) GetEncoding() Encoding {
	return BINARY
}
