package web

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/hellgate75/general_utils/log"
	"html/template"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

var __fileSeparator string = fmt.Sprintf("%c", os.PathSeparator)

var logger log.Logger

func InitLogger() {
	currentLogger, err := log.New("net/web")
	if err != nil {
		panic(err.Error())
	}
	logger = currentLogger
}

type ServerContext struct {
	ServerPath  string
	Environment map[string]interface{}
	Validator   *WebServerValidator
}

//Web Page Element
type WebPage struct {
	//Page title
	Title string
	//Page extension
	Extension string
	//Page content bytes
	Body []byte
	//Server Configuration Context
	Context *ServerContext
}

func __parseWebPageName(p *WebPage) (string, error) {
	if p == nil {
		return "", errors.New("Cannot find file name(s) for nil Web Page!!")
	}
	path := p.Context.ServerPath
	if strings.LastIndex(path, __fileSeparator) == len(path)-1 {
		path = path[0 : len(path)-1]
	}
	filename := path + __fileSeparator + p.Title
	if strings.TrimSpace(p.Extension) != "" {
		filename = fmt.Sprintf("%s.%s", filename, p.Extension)
	}
	return filename, nil
}

// Save web page to server path file
// Returns:
//    error Any error that can occurs during computation
func (p *WebPage) Save() error {
	filename, err := __parseWebPageName(p)
	if err != nil {
		if logger != nil {
			logger.Error(err)
		}
		return err
	}
	err = ioutil.WriteFile(filename, p.Body, 0660)
	if err != nil {
		if logger != nil {
			logger.Error(err)
		}
		return err
	}
	return nil
}

// Save web page from server path file
// Returns:
//    error Any error that can occurs during computation
func (p *WebPage) Load() error {
	filename, err := __parseWebPageName(p)
	if err != nil {
		if logger != nil {
			logger.Error(err)
		}
		return err
	}
	body, err2 := ioutil.ReadFile(filename)
	if err2 != nil {
		if logger != nil {
			logger.Error(err2)
		}
		return err2
	}
	p.Body = body
	return nil
}

// Load web page from file
// Returns:
//    (*WebPage Web Page pointer or nil in case of error,
//    error Any error that can occurs during computation)
func LoadWebPageFromFile(filename string, context *ServerContext) (*WebPage, error) {
	startIdx := strings.LastIndex(filename, __fileSeparator)
	endIdx := strings.LastIndex(filename, ".")
	title := filename
	extension := ""
	if startIdx >= 0 {
		if endIdx > 0 {
			title = filename[startIdx+1 : endIdx-1]

		} else {
			title = filename[startIdx+1:]
		}
	}
	if endIdx >= 0 {
		if startIdx < 0 {
			title = filename[:endIdx-1]
		}
		if endIdx+1 < len(filename) {
			extension = filename[endIdx+1:]
		}
	}
	body, err2 := ioutil.ReadFile(filename)
	if err2 != nil {
		if logger != nil {
			logger.Error(err2)
		}
		return nil, err2
	}
	return &WebPage{Title: title, Extension: extension, Body: body, Context: context}, nil
}

// Render web page from server path file or pre-loaded bytes
// Returns:
//    (string template parsed page text,
//    error Any error that can occurs during computation)
func (p *WebPage) Render() (string, error) {
	var templ *template.Template
	var err error
	if len(p.Body) > 0 {
		templ, err = template.New(p.Title).Parse(string(p.Body))
	} else {
		filename, err0 := __parseWebPageName(p)
		if err0 != nil {
			if logger != nil {
				logger.Error(err0)
			}
			return "", err0
		}
		templ, err = template.ParseFiles(filename)
	}
	if err != nil {
		if logger != nil {
			logger.Error(err)
		}
		return "", err
	}
	wr := bytes.NewBuffer([]byte{})
	err = templ.Execute(wr, p.Context.Environment)
	if err != nil {
		if logger != nil {
			logger.Error(err)
		}
		return "", err
	}
	return wr.String(), nil
}

// Render web page from server path file or pre-loaded bytes to an external writer
// Returns:
//    error Any error that can occurs during computation
func (p *WebPage) RenderOn(wr io.Writer) error {
	text, err := p.Render()
	if err != nil {
		if logger != nil {
			logger.Error(err)
		}
		return err
	}
	_, err = wr.Write([]byte(text))
	if err != nil {
		if logger != nil {
			logger.Error(err)
		}
		return err
	}
	return nil
}
