package tcp

import (
	"fmt"
	"github.com/hellgate75/general_utils/net/common"
	"io/ioutil"
	"net"
	"testing"
	"time"
)

func TestNewTcpServer(t *testing.T) {
	var logLevel common.LogLevel = common.INFO
	var port common.Port = common.Port(50837)
	var tcpActionsList []TcpEndpoint

	var expectedEntriesJsonMessage string = "{\"list\" : [{\"name\":\"William\",\"surname\":\"Smith\",\"age\":\"47\"},{\"name\":\"Mark\",\"surname\":\"White\",\"age\":\"27\"}]}"
	entriesEndpointFunc := func(conn net.Conn, outChan *chan interface{}, ctx common.NetContext) error {
		conn.Write([]byte(expectedEntriesJsonMessage))
		return nil
	}
	tcpActionsList = append(tcpActionsList, entriesEndpointFunc)
	var entriesList TcpEntriesList = TcpEntriesList(tcpActionsList)
	var restServer common.Server = NewTcpServer(logLevel, "", port, entriesList, nil)
	go restServer.Open()
	defer func() {
		restServer.Close()
		restServer.Destroy()
	}()
	time.Sleep(1 * time.Second)
	conn, errR := net.Dial("tcp", "localhost:50837")
	if errR != nil {
		t.Fatal(fmt.Sprintf("Unexpected error requesting tcp endpoint, Error is %s", errR.Error()))
	}
	byteArray, errT := ioutil.ReadAll(conn)
	if errT != nil {
		t.Fatal(fmt.Sprintf("Unexpected error reading tcp endpoint response, Error is %s", errT.Error()))
	}
	var text string = string(byteArray)
	if text != expectedEntriesJsonMessage {
		t.Fatal(fmt.Sprintf("Error comparing values. Expected <%s> but Given <%s>", expectedEntriesJsonMessage, text))
	}
}
