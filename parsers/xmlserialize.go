package parsers

import (
	"encoding/xml"
	"fmt"
	"github.com/hellgate75/general_utils/streams"
	"io/ioutil"
	"strconv"
)

func (this *XmlParserStruct) DeserializeFromFile(filePath string, mask *Type) error {
	var bytes []byte
	var err error
	if bytes, err = streams.LoadFileBytes(filePath); err == nil {
		return this.DeserializeFromBytes(bytes, mask)
	} else {
		return err
	}
}

func (this *XmlParserStruct) DeserializeFromBytes(bytes []byte, mask *Type) error {
	var err error
	if err = xml.Unmarshal(bytes, mask); err == nil {
		var length interface{} = "<null>"
		if bytes != nil {
			length = strconv.Itoa(len(bytes))
		}
		logger.Debug(fmt.Sprintf("Successful Deserialized bytes : %v", length))
		return nil
	} else {
		logger.Error(err)
		return err
	}
}

func (this *XmlParserStruct) SerializeToFile(filePath string, mask *Type) error {
	var bytes []byte
	var err error
	if bytes, err = this.SerializeToBytes(mask); err == nil {
		if err = ioutil.WriteFile(filePath, bytes, 666); err == nil {
			logger.Debug(fmt.Sprintf("Xml Parser :: File written: %s", filePath))
			return nil
		} else {
			logger.Error(err)
			return err
		}
	} else {
		logger.Error(err)
		return err
	}
}

func (this *XmlParserStruct) SerializeToBytes(mask *Type) ([]byte, error) {
	var bytes []byte
	var err error
	if bytes, err = xml.Marshal(mask); err == nil {
		return bytes, nil
	} else {
		logger.Error(err)
		return nil, err
	}
}
