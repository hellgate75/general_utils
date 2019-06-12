package parsers

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/hellgate75/general_utils/common"
	"github.com/hellgate75/general_utils/streams"
	"io/ioutil"
	"strconv"
)

func (this *base64ParserStruct) DeserializeFromFile(filePath string, mask common.Type) error {
	var bytes []byte
	var err error
	if bytes, err = streams.LoadFileBytes(filePath); err == nil {
		return this.DeserializeFromBytes(bytes, mask)
	} else {
		return err
	}
}

func (this *base64ParserStruct) DeserializeFromBytes(byteArray []byte, mask common.Type) error {
	buff := bytes.NewBuffer(byteArray)
	enc := base64.NewDecoder(base64.StdEncoding, buff)
	if no, err := enc.Read(byteArray); err == nil && no > 0 {
		var length interface{} = "<null>"
		byteArray = buff.Bytes()
		if byteArray != nil {
			length = strconv.Itoa(len(byteArray))
			if len(byteArray) > 0 {
				this.DeserializeFromBytes(byteArray, mask)
			}
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

func (this *base64ParserStruct) SerializeToFile(filePath string, mask common.Type) error {
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

func (this *base64ParserStruct) SerializeToBytes(mask common.Type) ([]byte, error) {
	byteArray, err := this.internalParser.SerializeToBytes(mask)
	if err != nil {
		if logger != nil {
			logger.Error(err)
		}
		return nil, err
	}
	buf := bytes.NewBuffer(byteArray)
	dec := base64.NewEncoder(base64.StdEncoding, buf)
	defer func() {
		dec.Close()
	}()
	if _, err := dec.Write(byteArray); err == nil {
		byteArray = buf.Bytes()
		return byteArray, nil
	} else {
		if logger != nil {
			logger.Error(err)
		}
		return nil, err
	}
}

func (this *base64ParserStruct) GetEncoding() Encoding {
	return BASE64
}
