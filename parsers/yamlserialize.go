package parsers

import (
	"fmt"
	"github.com/hellgate75/general_utils/common"
	"github.com/hellgate75/general_utils/streams"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"strconv"
)

func (this *yamlParserStruct) DeserializeFromFile(filePath string, mask common.Type) error {
	var bytes []byte
	var err error
	if bytes, err = streams.LoadFileBytes(filePath); err == nil {
		return this.DeserializeFromBytes(bytes, mask)
	} else {
		return err
	}
}

func (this *yamlParserStruct) DeserializeFromBytes(bytes []byte, mask common.Type) error {
	var err error
	if err = yaml.Unmarshal(bytes, &mask); err == nil {
		var length interface{} = "<null>"
		if bytes != nil {
			length = strconv.Itoa(len(bytes))
		}
		if logger != nil {
			logger.Debug(fmt.Sprintf("Yaml Parser :: Successful Deserialized bytes : %v", length))
		}
		return nil
	} else {
		if logger != nil {
			logger.Error(err)
		}
		return err
	}
}

func (this *yamlParserStruct) SerializeToFile(filePath string, mask common.Type) error {
	var bytes []byte
	var err error
	if bytes, err = this.SerializeToBytes(mask); err == nil {
		if err = ioutil.WriteFile(filePath, bytes, 666); err == nil {
			if logger != nil {
				logger.Debug(fmt.Sprintf("Yaml Parser :: File written: %s", filePath))
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

func (this *yamlParserStruct) SerializeToBytes(mask common.Type) ([]byte, error) {
	var bytes []byte
	var err error
	if bytes, err = yaml.Marshal(&mask); err == nil {
		return bytes, nil
	} else {
		if logger != nil {
			logger.Error(err)
		}
		return nil, err
	}
}
