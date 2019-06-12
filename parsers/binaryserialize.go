package parsers

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/hellgate75/general_utils/common"
	"github.com/hellgate75/general_utils/streams"
	"io/ioutil"
	"strconv"
)

func (this *binaryParserStruct) DeserializeFromFile(filePath string, mask common.Type) error {
	var bytes []byte
	var err error
	if bytes, err = streams.LoadFileBytes(filePath); err == nil {
		return this.DeserializeFromBytes(bytes, mask)
	} else {
		return err
	}
}

func (this *binaryParserStruct) DeserializeFromBytes(byteArray []byte, mask common.Type) error {
	buff := bytes.NewBuffer(byteArray)
	if err := binary.Read(buff, binary.BigEndian, mask); err == nil {
		var length interface{} = "<null>"
		byteArray = buff.Bytes()
		if byteArray != nil {
			length = strconv.Itoa(len(byteArray))
		}
		if logger != nil {
			logger.Debug(fmt.Sprintf("Base64 Parser :: Successful Deserialized bytes : %v", length))
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
	return BASE64
}
