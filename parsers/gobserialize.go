package parsers

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"github.com/hellgate75/general_utils/common"
	"github.com/hellgate75/general_utils/streams"
	"io/ioutil"
	"strconv"
)

func (this *gobParserStruct) DeserializeFromFile(filePath string, out interface{}) error {
	var byteArray []byte
	var err error
	if byteArray, err = streams.LoadFileBytes(filePath); err == nil {
		var length interface{} = "<null>"
		if byteArray != nil && len(byteArray) > 0 {
			length = strconv.Itoa(len(byteArray))
		} else {
			return errors.New("GoLang Parser :: Null or empty set of bytes!!!")
		}
		buff := bytes.NewBuffer(byteArray)
		enc := gob.NewDecoder(buff)
		if err := enc.Decode(out); err == nil {

			if logger != nil {
				logger.Debug(fmt.Sprintf("GoLang Parser :: Successful Deserialized bytes : %v", length))
			}
			return err
		} else {
			if logger != nil {
				logger.Error(err)
			}
			return err
		}
	} else {
		return err
	}
}

func (this *gobParserStruct) DeserializeFromBytes(byteArray []byte, out interface{}) error {
	var length interface{} = "<null>"
	if byteArray != nil && len(byteArray) > 0 {
		length = strconv.Itoa(len(byteArray))
	} else {
		return errors.New("GoLang Parser :: Input null or empty set of bytes!!!")
	}
	buff := bytes.NewBuffer(byteArray)
	enc := gob.NewDecoder(buff)
	if err := enc.Decode(out); err == nil {
		if logger != nil {
			logger.Debug(fmt.Sprintf("GoLang Parser :: Successful Deserialized bytes : %v", length))
		}
		return err
	} else {
		if logger != nil {
			logger.Error(err)
		}
		return err
	}
}

func (this *gobParserStruct) SerializeToFile(filePath string, mask common.Type) error {
	var bytes []byte
	var err error
	if bytes, err = this.SerializeToBytes(mask); err == nil {
		if err = ioutil.WriteFile(filePath, bytes, 0777); err == nil {
			if logger != nil {
				logger.Debug(fmt.Sprintf("GoLang Parser :: File written: %s", filePath))
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

func (this *gobParserStruct) SerializeToBytes(mask common.Type) ([]byte, error) {
	var writer LocalWriter = NewLocalWriter()
	dec := gob.NewEncoder(writer)
	if err := dec.Encode(mask); err == nil {
		byteArray, err1 := writer.GetBytes()
		return byteArray, err1
	} else {
		if logger != nil {
			logger.Error(err)
		}
		return nil, err
	}
}

func (this *gobParserStruct) GetEncoding() Encoding {
	return GOLANG
}
