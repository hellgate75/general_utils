package web

import (
	"fmt"
	"github.com/hellgate75/general_utils/net/common"
	"github.com/hellgate75/general_utils/streams"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
	"time"
)

func __getEnrichedPageTemplateCode() string {
	return "<h1>{{.Title}} (host: {{.Host}} - Port {{.Port}})</h1>\n" +
		"<p><image href=\"/golang.png\"></image></p>\n" +
		"<p>[<a href=\"{{.Link}}\">load</a>]</p>\n" +
		"<p>Welcome {{.User}}</p>\n" +
		"<div>Content : {{printf \"%s\" .Body}}</div>"

}

func TestNewWebServer(t *testing.T) {
	__initTests()
	baseFolder := fmt.Sprintf("%s%s%s%s%s%s%s%s%s%s%s", streams.GetCurrentPath(), __fileSeparator, "..", __fileSeparator, "..", __fileSeparator, "test", __fileSeparator, "servers", __fileSeparator, "web")
	__testWebPage.Body = []byte(__getEnrichedPageTemplateCode())
	os.MkdirAll(baseFolder, 0666)

	byteArr, errImg := ioutil.ReadFile(fmt.Sprintf("%s/testdata/golang.png", streams.GetCurrentPath()))
	if errImg != nil {
		t.Fatal(fmt.Sprintf("Error reading sample image : %s", errImg.Error()))
	} else {
		ioutil.WriteFile(fmt.Sprintf("%s%s%s", baseFolder, __fileSeparator, "golang.png"), byteArr, 0666)
	}
	byteArr, errImg = ioutil.ReadFile(fmt.Sprintf("%s/testdata/sample.js", streams.GetCurrentPath()))
	if errImg != nil {
		t.Fatal(fmt.Sprintf("Error reading sample javascript : %s", errImg.Error()))
	} else {
		ioutil.WriteFile(fmt.Sprintf("%s%s%s", baseFolder, __fileSeparator, "sample.js"), byteArr, 0666)
	}
	config := common.ServerConfig{
		ListeningPort: 8080,
		LogLevel:      common.INFO,
	}
	__testWebPage.Context.Environment["Port"] = 8080
	__testWebPage.Context.Environment["Host"] = "0.0.0.0"
	__testWebPage.Context.ServerPath = baseFolder
	__testWebPage.Save()
	ws := NewWebServer(baseFolder, config, __testWebPage.Context.Environment, nil, nil)
	ws = NewWebServer(baseFolder, config, __testWebPage.Context.Environment, &WebServerValidator{}, nil)
	defer func() {
		os.RemoveAll(baseFolder)
	}()
	go ws.Open()
	time.Sleep(500 * time.Millisecond)
	response, errX := http.Get("http://localhost:8080/test_html")
	if errX != nil {
		t.Fatal(fmt.Sprintf("Unexpected error reading page url, error is : %s", errX.Error()))
	}
	var bytes []byte = make([]byte, response.ContentLength)
	response.Body.Read(bytes)
	var output string = fmt.Sprintf("%s", bytes)
	var expected string = "<h1>Test Page (host: 0.0.0.0 - Port 8080)</h1>\n" +
		"<p><image href=\"/golang.png\"></image></p>\n" +
		"<p>[<a href=\"#\">load</a>]</p>\n" +
		"<p>Welcome Fabrizio</p>\n" +
		"<div>Content : This is page content</div>"
	if output != expected {
		t.Fatal(fmt.Sprintf("Error comparing values. Expected <%s> but Given <%s>", expected, output))
	}
	response, errX = http.Get("http://localhost:8080/test_html.template")
	if errX != nil {
		t.Fatal(fmt.Sprintf("Unexpected error reading page url, error is : %s", errX.Error()))
	}
	bytes = make([]byte, response.ContentLength)
	response.Body.Read(bytes)
	output = fmt.Sprintf("%s", bytes)
	if output != expected {
		t.Fatal(fmt.Sprintf("Error comparing values. Expected <%s> but Given <%s>", expected, output))
	}
	response, errX = http.Get("http://localhost:8080/golang.png")
	if errX != nil {
		t.Fatal(fmt.Sprintf("Unexpected error reading page image, error is : %s", errX.Error()))
	}
	response, errX = http.Get("http://localhost:8080/sample.js")
	if errX != nil {
		t.Fatal(fmt.Sprintf("Unexpected error reading page script, error is : %s", errX.Error()))
	}
	ws.Close()
	ws.Destroy()
	__destroyTests()
}
