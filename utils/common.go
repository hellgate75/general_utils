package utils

import (
	"errors"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func ExistsFile(file string) bool {
	_, err := os.Stat(file)
	return err == nil
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
