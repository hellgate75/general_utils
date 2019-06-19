package parsers

import (
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/hellgate75/general_utils/common"
	"github.com/hellgate75/general_utils/streams"
	"io/ioutil"
	"strconv"
)

func (this *xmlParserStruct) DeserializeFromFile(filePath string, out interface{}) error {
	var bytes []byte
	var err error
	if bytes, err = streams.LoadFileBytes(filePath); err == nil {
		if err = xml.Unmarshal(bytes, out); err == nil {
			var length interface{} = "<null>"
			if bytes != nil && len(bytes) > 0 {
				length = strconv.Itoa(len(bytes))
			} else {
				return errors.New("Xml Parser :: Input null or empty set of bytes!!!")
			}
			if logger != nil {
				logger.Debug(fmt.Sprintf("Xml Parser :: Successful Deserialized bytes : %v", length))
			}
			return nil
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

func (this *xmlParserStruct) DeserializeFromBytes(bytes []byte, out interface{}) error {
	var err error
	if err = xml.Unmarshal(bytes, out); err == nil {
		var length interface{} = "<null>"
		if bytes != nil && len(bytes) > 0 {
			length = strconv.Itoa(len(bytes))
		} else {
			return errors.New("Xml Parser :: Input null or empty set of bytes!!!")
		}
		if logger != nil {
			logger.Debug(fmt.Sprintf("Xml Parser :: Successful Deserialized bytes : %v", length))
		}
		return nil
	} else {
		if logger != nil {
			logger.Error(err)
		}
		return err
	}
}

func (this *xmlParserStruct) SerializeToFile(filePath string, mask common.Type) error {
	var bytes []byte
	var err error
	if bytes, err = this.SerializeToBytes(mask); err == nil {
		if err = ioutil.WriteFile(filePath, bytes, 666); err == nil {
			if logger != nil {
				logger.Debug(fmt.Sprintf("Xml Parser :: File written: %s", filePath))
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

func (this *xmlParserStruct) SerializeToBytes(mask common.Type) ([]byte, error) {
	var bytes []byte
	var err error
	if bytes, err = xml.Marshal(&mask); err == nil {
		return bytes, nil
	} else {
		if logger != nil {
			logger.Error(err)
		}
		return nil, err
	}
}

func (this *xmlParserStruct) GetEncoding() Encoding {
	return XML
}
