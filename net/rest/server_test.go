package rest

import (
	"fmt"
	"github.com/hellgate75/general_utils/net/common"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
)

func TestNewRestServer(t *testing.T) {
	var logLevel common.LogLevel = common.INFO
	var port common.Port = common.Port(50837)
	var stateHandler common.HttpStateHandler
	var err error
	var errorMap map[int]common.HTTPAction = make(map[int]common.HTTPAction)
	handle404Error := func(w http.ResponseWriter, r *http.Request) error {
		var message404Json string = fmt.Sprintf("{code=\"404\",message=\"Rest Endpoint '%s': Not Found\"}", r.URL.Path)
		w.WriteHeader(404)
		w.Write([]byte(message404Json))
		return nil
	}
	errorMap[404] = common.HTTPAction(handle404Error)
	handle500Error := func(w http.ResponseWriter, r *http.Request) error {
		var message500Json string = fmt.Sprintf("{code=\"500\",message=\"Rest Endpoint '%s': Internal Server Error\"}", r.URL.Path)
		w.WriteHeader(500)
		w.Write([]byte(message500Json))
		return nil
	}
	errorMap[500] = common.HTTPAction(handle500Error)
	stateHandler, err = common.NewHttpStateHandler(errorMap)
	if err != nil {
		t.Fatal(fmt.Sprintf("Error Creating New HttpStateHandler : %s", err.Error()))
	}
	var endPointMap map[string]RestEndpoint = make(map[string]RestEndpoint)
	var expectedEntriesJsonMessage string = "{\"list\" : [{\"name\":\"William\",\"surname\":\"Smith\",\"age\":\"47\"},{\"name\":\"Mark\",\"surname\":\"White\",\"age\":\"27\"}]}"
	entriesEndointFunc := func(handler common.HttpStateHandler, query url.Values, outChan *chan interface{}, w http.ResponseWriter, r *http.Request, ctx common.NetContext) error {
		//		w.WriteHeader(http.StatusOK)
		w.Write([]byte(expectedEntriesJsonMessage))
		return nil
	}
	endPointMap["/entries"] = RestEndpoint(entriesEndointFunc)
	var entriesMap RestEntriesMap = RestEntriesMap(endPointMap)
	var restServer common.Server = NewRestServer(logLevel, "localhost", port, stateHandler, entriesMap)
	go restServer.Open()
	defer func() {
		restServer.Close()
		restServer.Destroy()
	}()
	var urlMap map[string][]string = make(map[string][]string)
	urlMap["Request"] = []string{"entriesList"}
	urlMap["State"] = []string{"open"}
	var values url.Values = url.Values(urlMap)
	response, errR := http.PostForm("http://localhost:50837/entries", values)
	if errR != nil {
		t.Fatal(fmt.Sprintf("Unexpected error requesting rest endpoint '/entries', Error is %s", errR.Error()))
	}
	status := response.Status
	if "200 OK" != status {
		t.Fatal(fmt.Sprintf("Error comparing POST status. Expected <%s> but Given <%s>", "200 OK", status))
	}
	byteArray, errT := ioutil.ReadAll(response.Body)
	if errT != nil {
		t.Fatal(fmt.Sprintf("Unexpected error reading rest endpoint '/entries' response, Error is %s", errT.Error()))
	}
	var text string = string(byteArray)
	if text != expectedEntriesJsonMessage {
		t.Fatal(fmt.Sprintf("Error comparing values. Expected <%s> but Given <%s>", expectedEntriesJsonMessage, text))
	}
}
