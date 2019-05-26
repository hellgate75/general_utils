package streams

import (
	"github.com/hellgate75/general_utils/log"
	"io/ioutil"
	"net/http"
)

var logger log.Logger = log.GetLogger()

func DownloadFile(filepath string, url string) (err error) {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		logger.Error(err)
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error(err)
		return err
	}
	// Writer the body to file
	err = ioutil.WriteFile(filepath, body, 0666)
	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}
