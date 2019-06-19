package parsers

import (
	"errors"
	"fmt"
	"github.com/hellgate75/general_utils/common"
	"github.com/hellgate75/general_utils/streams"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"strconv"
)

func (this *yamlParserStruct) DeserializeFromFile(filePath string, out interface{}) error {
	var bytes []byte
	var err error

	if bytes, err = streams.LoadFileBytes(filePath); err == nil {
		var length interface{} = "<null>"
		if bytes != nil && len(bytes) > 0 {
			length = strconv.Itoa(len(bytes))
		} else {
			return errors.New("Yaml Parser :: Input null or empty set of bytes!!!")
		}
		if err = yaml.Unmarshal(bytes, out); err == nil {
			if logger != nil {
				logger.Debug(fmt.Sprintf("Yaml Parser :: Successful Deserialized file %s length: %s", filePath, length))
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

func (this *yamlParserStruct) DeserializeFromBytes(bytes []byte, out interface{}) error {
	var err error
	var length interface{} = "<null>"
	if bytes != nil && len(bytes) > 0 {
		length = strconv.Itoa(len(bytes))
	} else {
		return errors.New("Yaml Parser :: Input null or empty set of bytes!!!")
	}
	//	var inInterface interface{}
	if err = yaml.Unmarshal(bytes, out); err == nil {
		if logger != nil {
			logger.Debug(fmt.Sprintf("Yaml Parser :: Successful Deserialized bytes : %v", length))
		}
		//outIface := structureCleaner.MapToInterface(inInterface.(map[interface{}]interface{}))
		//_assign(&out, &outIface)
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

func (this *yamlParserStruct) GetEncoding() Encoding {
	return YAML
}
