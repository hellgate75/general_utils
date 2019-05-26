package streams

import (
	"io/ioutil"
	"os"
)

// Get current Go execution path.
//
// Returns:
//   string Current absolute path
func GetCurrentPath() string {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return pwd
}

// Load all content in a file as string.
//
// Parameters:
//   path (string) Absolute file path
//
// Returns:
//   string File content
//   error Any suitable error risen during code execution
func LoadFileContent(path string) (string, error) {
	var err error
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}
	return string(bytes), err
}

// Load all content in a file as byte array.
//
// Parameters:
//   path (string) Absolute file path
//
// Returns:
//   []byte File content
//   error Any suitable error risen during code execution
func LoadFileBytes(path string) ([]byte, error) {
	var err error
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return bytes, err
}
