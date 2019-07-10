package streams

import (
	"io/ioutil"
	"net/http"
)

// Dowload file from url and save locally.
//
// Parameters:
//   filepath (string) Absolute destination file path
//   url (string) Source file url
//
// Returns:
//   error Any suitable error risen during code execution
func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		if logger != nil {
			logger.Error(err)
		}
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		if logger != nil {
			logger.Error(err)
		}
		return err
	}
	// Writer the body to file
	err = ioutil.WriteFile(filepath, body, 0777)
	if err != nil {
		if logger != nil {
			logger.Error(err)
		}
		return err
	}

	return nil
}

// Dowload file from url and return content byte array.
//
// Parameters:
//   url (string) Source file url
//
// Returns:
//   []byte Bytes contained in the remote support
//   error Any suitable error risen during code execution
func DownloadFileAsByteArray(url string) ([]byte, error) {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		if logger != nil {
			logger.Error(err)
		}
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		if logger != nil {
			logger.Error(err)
		}
		return nil, err
	}
	return body, nil
}
