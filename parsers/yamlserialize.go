package parsers

import (
	"fmt"
	"github.com/hellgate75/general_utils/streams"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"strconv"
)

func (this *YamlParserStruct) DeserializeFromFile(filePath string, mask *Type) error {
	var bytes []byte
	var err error
	if bytes, err = streams.LoadFileBytes(filePath); err == nil {
		return this.DeserializeFromBytes(bytes, mask)
	} else {
		return err
	}
}

func (this *YamlParserStruct) DeserializeFromBytes(bytes []byte, mask *Type) error {
	var err error
	if err = yaml.Unmarshal(bytes, mask); err == nil {
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

func (this *YamlParserStruct) SerializeToFile(filePath string, mask *Type) error {
	var bytes []byte
	var err error
	if bytes, err = this.SerializeToBytes(mask); err == nil {
		if err = ioutil.WriteFile(filePath, bytes, 666); err == nil {
			logger.Debug(fmt.Sprintf("Yaml Parser :: File written: %s", filePath))
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

func (this *YamlParserStruct) SerializeToBytes(mask *Type) ([]byte, error) {
	var bytes []byte
	var err error
	if bytes, err = yaml.Marshal(mask); err == nil {
		return bytes, nil
	} else {
		logger.Error(err)
		return nil, err
	}
}
