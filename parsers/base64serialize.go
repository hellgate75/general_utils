package parsers

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"errors"
	"fmt"
	"github.com/hellgate75/general_utils/common"
	"github.com/hellgate75/general_utils/streams"
	"io/ioutil"
	"strconv"
)

func (this *base64ParserStruct) DeserializeFromFile(filePath string, out interface{}) error {
	var byteArray []byte
	var err error
	var length interface{} = "<null>"
	if byteArray, err = streams.LoadFileBytes(filePath); err == nil {
		if byteArray != nil && len(byteArray) > 0 {
			length = strconv.Itoa(len(byteArray))
		} else {
			return errors.New("Base64 Parser :: Input null or empty set of bytes!!!")
		}
		inStr := bytes.NewBuffer(byteArray).String()
		outStr, err := base64.StdEncoding.DecodeString(inStr)
		if err != nil {
			if logger != nil {
				logger.Debug(fmt.Sprintf("Base64 Parser :: Successful Deserialized bytes : %v", length))
			}
			return err
		}
		var dstArray []byte = []byte(outStr)
		if err == nil {
			if len(dstArray) > 0 {
				//de-serialize object
				buff := bytes.NewBuffer(dstArray)
				enc := gob.NewDecoder(buff)
				if err = enc.Decode(out); err == nil {
					if logger != nil {
						logger.Debug(fmt.Sprintf("Base64 Parser :: GoLang Parser :: Successful Deserialized bytes : %v", length))
					}
				} else {
					if logger != nil {
						logger.Error(err)
					}
					return err
				}
				if logger != nil {
					logger.Debug(fmt.Sprintf("Base64 Parser :: Successful Deserialized bytes : %v", length))
				}
			} else {
				err = errors.New("Base64 Parser :: Empty decoded array!!!")
				if logger != nil {
					logger.Error(err)
				}
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

func (this *base64ParserStruct) DeserializeFromBytes(byteArray []byte, out interface{}) error {
	var length interface{} = "<null>"
	if byteArray != nil && len(byteArray) > 0 {
		length = strconv.Itoa(len(byteArray))
	} else {
		return errors.New("Base64 Parser :: Input null or empty set of bytes!!!")
	}
	inStr := bytes.NewBuffer(byteArray).String()
	outStr, err := base64.StdEncoding.DecodeString(inStr)
	if err != nil {
		if logger != nil {
			logger.Debug(fmt.Sprintf("Base64 Parser :: Successful Deserialized bytes : %v", length))
		}
		return err
	}
	var dstArray []byte = []byte(outStr)
	if len(dstArray) > 0 {
		//de-serialize object
		buff := bytes.NewBuffer(dstArray)
		enc := gob.NewDecoder(buff)
		if err = enc.Decode(out); err == nil {
			if logger != nil {
				logger.Debug(fmt.Sprintf("Base64 Parser :: GoLang Parser :: Successful Deserialized bytes : %v", length))
			}
		} else {
			if logger != nil {
				logger.Error(err)
			}
			return err
		}
		if err != nil && logger != nil {
			logger.Debug(fmt.Sprintf("Base64 Parser :: Successful Deserialized bytes : %v", length))
		}
		return err
	} else {
		err = errors.New("Base64 Parser :: Empty decoded array!!!")
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
	var writer LocalWriter = NewLocalWriter()
	var byteArray []byte
	var err error
	dec := gob.NewEncoder(writer)
	if err = dec.Encode(mask); err == nil {
		byteArray, err = writer.GetBytes()
		if err != nil {
			if logger != nil {
				logger.Error(err)
			}
			return nil, err
		}
	} else {
		if logger != nil {
			logger.Error(err)
		}
		return nil, err
	}
	outStr := base64.StdEncoding.EncodeToString(byteArray)
	return []byte(outStr), nil
}

func (this *base64ParserStruct) GetEncoding() Encoding {
	return BASE64
}
