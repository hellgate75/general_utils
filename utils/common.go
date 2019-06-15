package utils

import (
	"github.com/hellgate75/general_utils/log"
	"os"
	"strconv"
	"strings"
)

var logger log.Logger

func InitLogger() {
	currentLogger, err := log.New("utils")
	if err != nil {
		panic(err.Error())
	}
	logger = currentLogger
}

func FileExists(file string) bool {
	_, err := os.Stat(file)
	return err == nil
}

func CreateFileIfNotExists(file string) error {
	if FileExists(file) {
		return nil
	}
	emptyFile, err := os.Create(file)
	if err != nil {
		return err
	}
	defer func(f *os.File) {
		f.Close()
	}(emptyFile)
	return nil
}

func CreateFile(file string) (*os.File, error) {
	if FileExists(file) {
		return os.Open(file)
	}
	emptyFile, err := os.Create(file)
	if err != nil {
		return nil, err
	}
	return emptyFile, nil
}

func CreateFileAndUse(file string, consumer func(*os.File) (interface{}, error)) (interface{}, error) {
	var emptyFile *os.File
	var err error
	if !FileExists(file) {
		emptyFile, err = os.Create(file)
	} else {
		emptyFile, err = os.Open(file)
	}
	if err != nil {
		return nil, err
	}
	defer func(f *os.File) {
		f.Close()
	}(emptyFile)
	return consumer(emptyFile)
}

func DeleteIfExists(file string) error {
	_, err := os.Stat(file)
	if err == nil {
		return os.Remove(file)
	}
	return err
}

func MakeFolderIfNotExists(folder string) error {
	if _, err := os.Stat(folder); err != nil {
		err := os.MkdirAll(folder, 0777)
		return err
	}
	return nil
}

func EncodeBytes(decodedByteArray []byte) []byte {
	newBytes := make([]byte, 0)
	for _, byteElem := range decodedByteArray {
		newBytes = append(newBytes, byteElem-20)
	}
	return newBytes
	//return decodedByteArray
}

func DecodeBytes(encodedByteArray []byte) []byte {
	newBytes := make([]byte, 0)
	for _, byteElem := range encodedByteArray {
		newBytes = append(newBytes, byteElem+20)
	}
	return newBytes
	//return encodedByteArray
}

func CorrectInput(input string) string {
	return strings.TrimSpace(strings.ToLower(input))
}

func StringToInt(s string) (int, error) {
	return strconv.Atoi(s)
}

func IntToString(n int) string {
	return strconv.Itoa(n)
}
