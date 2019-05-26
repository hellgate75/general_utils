package config

import (
	"bytes"
	"fmt"
	"github.com/blakesmith/ar"
	"github.com/hellgate75/general_utils/log"
	"github.com/hellgate75/general_utils/parsers"
	"github.com/hellgate75/general_utils/streams"
	"io"
	"os"
	"strings"
)

var logger log.Logger = log.GetLogger()

func InitDatabaseConfig(arFilePath string, mask *parsers.Type, deserializer func([]byte, *parsers.Type) error) error {
	if _, err := os.Stat(arFilePath); err == nil {
		err, xMap := ReadDatabaseConfig(arFilePath)
		if err == nil {
			if val, ok := xMap[".settings"]; ok {
				err := deserializer(val, mask)
				if err != nil {
					logger.Error(err)
					return err
				}
				logger.Info(fmt.Sprintf("Configuration reloaded from file : %s", arFilePath))
			}
		}
	} else {
		var elems []string = strings.Split(arFilePath, "/")
		var fileName string = elems[len(elems)-1]
		LoadDatabaseFromURL(fileName, arFilePath, mask, deserializer)
	}
	return nil
}

func ReadDatabaseConfig(arFileName string) (error, map[string][]byte) {
	var files map[string][]byte = make(map[string][]byte, 0)
	file, err := os.Open(arFileName)
	if err != nil {
		return err, files
	}
	reader := ar.NewReader(file)

	var header *ar.Header
	var errH error

	if header, errH = reader.Next(); errH == nil {
		var buf bytes.Buffer
		io.Copy(&buf, reader)
		files[header.Name] = buf.Bytes()
	} else {
		logger.ErrorS(fmt.Sprintf("Error in Configuration reloaded from file : %s", arFileName))
		logger.Error(errH)
		return errH, nil
	}
	return nil, files
}

func LoadDatabaseFromURL(arFileName string, url string, mask *parsers.Type, deserializer func([]byte, *parsers.Type) error) error {
	if _, err := os.Stat(arFileName); err != nil {
		streams.DownloadFile(arFileName, url)
	}
	err, xMap := ReadDatabaseConfig(arFileName)
	if err == nil {
		logger.Debug(fmt.Sprintf("Database loaded from url: %s", url))
		if val, ok := xMap[".settings"]; ok {
			err := deserializer(val, mask)
			if err != nil {
				logger.Error(err)
				return err
			}
			logger.Warning(fmt.Sprintf("Configuration reloaded from file : %s", arFileName))
		}
	} else {
		logger.ErrorS(fmt.Sprintf("Error in Configuration reloaded from file : %s", arFileName))
		logger.Error(err)
		return err
	}
	return nil
}
