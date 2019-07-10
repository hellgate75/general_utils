package web

import (
	"bytes"
	"fmt"
	"github.com/hellgate75/general_utils/streams"
	"io/ioutil"
	"os"
	"testing"
)

var __testEnvironment map[string]interface{} = make(map[string]interface{})
var __testContext *ServerContext
var __testWebPage *WebPage

func __getDefaultPageCode() string {
	return "<h1>{{.Title}}</h1>\n" +
		"<p>[<a href=\"{{.Link}}\">load</a>]</p>\n" +
		"<p>Welcome {{.User}}</p>\n" +
		"<div>Content : {{printf \"%s\" .Body}}</div>"

}

func __getWrongPageCode() string {
	return "<h1>{{.Title2}}</h1>\n" +
		"<p>[<a href=\"{{.Link}\">load</a>]</p>\n" +
		"<p>Welcome {{.User2}}</p>\n" +
		"<div>Content : {{printf \"%s\" .Body2}}</div>"

}

func __initTests() {
	__testEnvironment["Title"] = "Test Page"
	__testEnvironment["User"] = "Fabrizio"
	__testEnvironment["Link"] = "#"
	__testEnvironment["Body"] = []byte("This is page content")
	__testContext = &ServerContext{
		ServerPath:  fmt.Sprintf("%s%s%s%s%s%s%s%s%s%s%s", streams.GetCurrentPath(), __fileSeparator, "..", __fileSeparator, "..", __fileSeparator, "test", __fileSeparator, "servers", __fileSeparator, "web"),
		Environment: __testEnvironment,
	}
	__testWebPage = &WebPage{
		Title:     "test_html",
		Extension: "template",
		Context:   __testContext,
	}
	os.MkdirAll(__testContext.ServerPath, 0777)
}

func __destroyTests() {
	os.RemoveAll(__testContext.ServerPath)
}

func TestParseWebPageName(t *testing.T) {
	_, err := __parseWebPageName(nil)
	if err == nil {
		t.Fatal("Unexpected nil error in page name parse for nil web page")
	}
	__initTests()
	_, err = __parseWebPageName(__testWebPage)
	if err != nil {
		t.Fatal(fmt.Sprintf("Unexpected error parsing page url, error is : %s", err.Error()))
	}
	__destroyTests()
}

func TestWebPageSave(t *testing.T) {
	__initTests()
	__testWebPage.Body = []byte(__getDefaultPageCode())
	__testWebPage.Save()
	path, err := __parseWebPageName(__testWebPage)
	if err != nil {
		t.Fatal(fmt.Sprintf("Unexpected error parsing page name, error is : %s", err.Error()))
	}
	_, err = os.Stat(path)
	if err != nil {
		t.Fatal(fmt.Sprintf("Saved Web Page not found, error is : %s", err.Error()))
	}
	byteArray, errRead := ioutil.ReadFile(path)
	if errRead != nil {
		t.Fatal(fmt.Sprintf("Unexpected error reading page name from file %s, error is : %s", path, errRead.Error()))
	}
	var expectedBytes []byte = []byte{60, 104, 49, 62, 123, 123, 46, 84, 105, 116, 108, 101, 125, 125, 60, 47, 104, 49, 62, 10, 60, 112, 62, 91, 60, 97, 32, 104, 114, 101, 102, 61, 34, 123, 123, 46, 76, 105, 110, 107, 125, 125, 34, 62, 108, 111, 97, 100, 60, 47, 97, 62, 93, 60, 47, 112, 62, 110, 60, 112, 62, 87, 101, 108, 99, 111, 109, 101, 32, 123, 123, 46, 85, 115, 101, 114, 125, 125, 60, 47, 112, 62, 10, 60, 100, 105, 118, 62, 67, 111, 110, 116, 101, 110, 116, 32, 58, 32, 123, 123, 112, 114, 105, 110, 116, 102, 32, 34, 37, 115, 34, 32, 46, 66, 111, 100, 121, 125, 125, 60, 47, 100, 105, 118, 62}
	var expectedText string = __getDefaultPageCode()
	if len(byteArray) != len(expectedBytes) {
		t.Fatal(fmt.Sprintf("Error comparing values. Expected <%v> but Given <%v>", expectedBytes, byteArray))
	}
	if string(byteArray) != expectedText {
		t.Fatal(fmt.Sprintf("Error comparing values. Expected <%s> but Given <%s>", expectedText, byteArray))
	}
	__destroyTests()
}

func TestWebPageLoad(t *testing.T) {
	__initTests()
	filename, err := __parseWebPageName(__testWebPage)
	if err != nil {
		t.Fatal(fmt.Sprintf("Unexpected error parsing page name, error is : %s", err.Error()))
	}
	err = ioutil.WriteFile(filename, []byte(__getDefaultPageCode()), 0777)
	if err != nil {
		t.Fatal(fmt.Sprintf("Unexpected error saving page from file %s, error is : %s", filename, err.Error()))
	}
	err = __testWebPage.Load()
	if err != nil {
		t.Fatal(fmt.Sprintf("Unexpected error loading page , error is : %s", err.Error()))
	}
	var expectedBytes []byte = []byte{60, 104, 49, 62, 123, 123, 46, 84, 105, 116, 108, 101, 125, 125, 60, 47, 104, 49, 62, 10, 60, 112, 62, 91, 60, 97, 32, 104, 114, 101, 102, 61, 34, 123, 123, 46, 76, 105, 110, 107, 125, 125, 34, 62, 108, 111, 97, 100, 60, 47, 97, 62, 93, 60, 47, 112, 62, 110, 60, 112, 62, 87, 101, 108, 99, 111, 109, 101, 32, 123, 123, 46, 85, 115, 101, 114, 125, 125, 60, 47, 112, 62, 10, 60, 100, 105, 118, 62, 67, 111, 110, 116, 101, 110, 116, 32, 58, 32, 123, 123, 112, 114, 105, 110, 116, 102, 32, 34, 37, 115, 34, 32, 46, 66, 111, 100, 121, 125, 125, 60, 47, 100, 105, 118, 62}
	var expectedText string = __getDefaultPageCode()
	if len(__testWebPage.Body) != len(expectedBytes) {
		t.Fatal(fmt.Sprintf("Error comparing values. Expected <%v> but Given <%v>", expectedBytes, __testWebPage.Body))
	}
	if string(__testWebPage.Body) != expectedText {
		t.Fatal(fmt.Sprintf("Error comparing values. Expected <%s> but Given <%s>", expectedText, __testWebPage.Body))
	}
	__destroyTests()
}

func TestWebPageRender(t *testing.T) {
	__initTests()
	filename, err := __parseWebPageName(__testWebPage)
	if err != nil {
		t.Fatal(fmt.Sprintf("Unexpected error parsing page name, error is : %s", err.Error()))
	}
	err = ioutil.WriteFile(filename, []byte(__getDefaultPageCode()), 0777)
	if err != nil {
		t.Fatal(fmt.Sprintf("Unexpected error saving page from file %s, error is : %s", filename, err.Error()))
	}
	err = __testWebPage.Load()
	if err != nil {
		t.Fatal(fmt.Sprintf("Unexpected error loading page , error is : %s", err.Error()))
	}
	var text string
	text, err = __testWebPage.Render()
	if err != nil {
		t.Fatal(fmt.Sprintf("Unexpected error rendering template page error is : %s", err.Error()))
	}
	var expectedText string = "<h1>Test Page</h1>\n<p>[<a href=\"#\">load</a>]</p>\n<p>Welcome Fabrizio</p>\n<div>Content : This is page content</div>"
	if text != expectedText {
		t.Fatal(fmt.Sprintf("Error comparing values. Expected <%s> but Given <%s>", expectedText, text))
	}
	__testWebPage.Body = []byte{}
	text, err = __testWebPage.Render()
	if err != nil {
		t.Fatal(fmt.Sprintf("Unexpected error rendering template page  from file %s,, error is : %s", filename, err.Error()))
	}
	if text != expectedText {
		t.Fatal(fmt.Sprintf("Error comparing values. Expected <%s> but Given <%s>", expectedText, text))
	}
	err = ioutil.WriteFile(filename, []byte(__getWrongPageCode()), 0777)
	if err != nil {
		t.Fatal(fmt.Sprintf("Unexpected error saving page from file %s, error is : %s", filename, err.Error()))
	}
	err = __testWebPage.Load()
	if err != nil {
		t.Fatal(fmt.Sprintf("Unexpected error loading page , error is : %s", err.Error()))
	}
	_, err = __testWebPage.Render()
	if err == nil {
		t.Fatal("Unexpected nil error in rendering template for unparsable page")
	}
	__testWebPage.Body = []byte{}
	_, err = __testWebPage.Render()
	if err == nil {
		t.Fatal("Unexpected nil error in rendering template for unparsable page")
	}
	__destroyTests()
}

func TestWebPageRenderOn(t *testing.T) {
	__initTests()
	filename, err := __parseWebPageName(__testWebPage)
	if err != nil {
		t.Fatal(fmt.Sprintf("Unexpected error parsing page name, error is : %s", err.Error()))
	}
	err = ioutil.WriteFile(filename, []byte(__getDefaultPageCode()), 0777)
	if err != nil {
		t.Fatal(fmt.Sprintf("Unexpected error saving page from file %s, error is : %s", filename, err.Error()))
	}
	err = __testWebPage.Load()
	if err != nil {
		t.Fatal(fmt.Sprintf("Unexpected error loading page , error is : %s", err.Error()))
	}
	var text string
	writer := bytes.NewBuffer([]byte{})
	err = __testWebPage.RenderOn(writer)
	text = writer.String()
	if err != nil {
		t.Fatal(fmt.Sprintf("Unexpected error rendering template page error is : %s", err.Error()))
	}
	var expectedText string = "<h1>Test Page</h1>\n<p>[<a href=\"#\">load</a>]</p>\n<p>Welcome Fabrizio</p>\n<div>Content : This is page content</div>"
	if text != expectedText {
		t.Fatal(fmt.Sprintf("Error comparing values. Expected <%s> but Given <%s>", expectedText, text))
	}
	__testWebPage.Body = []byte{}
	text, err = __testWebPage.Render()
	if err != nil {
		t.Fatal(fmt.Sprintf("Unexpected error rendering template page  from file %s,, error is : %s", filename, err.Error()))
	}
	if text != expectedText {
		t.Fatal(fmt.Sprintf("Error comparing values. Expected <%s> but Given <%s>", expectedText, text))
	}
	err = ioutil.WriteFile(filename, []byte(__getWrongPageCode()), 0777)
	if err != nil {
		t.Fatal(fmt.Sprintf("Unexpected error saving page from file %s, error is : %s", filename, err.Error()))
	}
	err = __testWebPage.Load()
	if err != nil {
		t.Fatal(fmt.Sprintf("Unexpected error loading page , error is : %s", err.Error()))
	}
	writer = bytes.NewBuffer([]byte{})
	err = __testWebPage.RenderOn(writer)
	if err == nil {
		t.Fatal("Unexpected nil error in rendering template for unparsable page")
	}
	__testWebPage.Body = []byte{}
	writer = bytes.NewBuffer([]byte{})
	err = __testWebPage.RenderOn(writer)
	if err == nil {
		t.Fatal("Unexpected nil error in rendering template for unparsable page")
	}
	__destroyTests()
}

func TestLoadWebPageFromFile(t *testing.T) {
	__initTests()
	filename, err := __parseWebPageName(__testWebPage)
	if err != nil {
		t.Fatal(fmt.Sprintf("Unexpected error parsing page name, error is : %s", err.Error()))
	}
	err = ioutil.WriteFile(filename, []byte(__getDefaultPageCode()), 0777)
	if err != nil {
		t.Fatal(fmt.Sprintf("Unexpected error saving page from file %s, error is : %s", filename, err.Error()))
	}
	var webpage *WebPage
	webpage, err = LoadWebPageFromFile(filename, __testContext)
	if err != nil {
		t.Fatal(fmt.Sprintf("Unexpected error loading page , error is : %s", err.Error()))
	}
	var expectedBytes []byte = []byte{60, 104, 49, 62, 123, 123, 46, 84, 105, 116, 108, 101, 125, 125, 60, 47, 104, 49, 62, 10, 60, 112, 62, 91, 60, 97, 32, 104, 114, 101, 102, 61, 34, 123, 123, 46, 76, 105, 110, 107, 125, 125, 34, 62, 108, 111, 97, 100, 60, 47, 97, 62, 93, 60, 47, 112, 62, 110, 60, 112, 62, 87, 101, 108, 99, 111, 109, 101, 32, 123, 123, 46, 85, 115, 101, 114, 125, 125, 60, 47, 112, 62, 10, 60, 100, 105, 118, 62, 67, 111, 110, 116, 101, 110, 116, 32, 58, 32, 123, 123, 112, 114, 105, 110, 116, 102, 32, 34, 37, 115, 34, 32, 46, 66, 111, 100, 121, 125, 125, 60, 47, 100, 105, 118, 62}
	var expectedText string = __getDefaultPageCode()
	if len(webpage.Body) != len(expectedBytes) {
		t.Fatal(fmt.Sprintf("Error comparing values. Expected <%v> but Given <%v>", expectedBytes, webpage.Body))
	}
	if string(webpage.Body) != expectedText {
		t.Fatal(fmt.Sprintf("Error comparing values. Expected <%s> but Given <%s>", expectedText, webpage.Body))
	}
	_, err = LoadWebPageFromFile(fmt.Sprintf("%s23", filename), __testContext)
	if err == nil {
		t.Fatal("Unexpected nil error in loading template for not existing page file")
	}
	__destroyTests()
}
