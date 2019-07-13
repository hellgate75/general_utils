package web

import (
	"errors"
	errs "github.com/hellgate75/general_utils/errors"
	"github.com/hellgate75/general_utils/net/common"
	//	"net"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"
)

// Http Response Handler element
var HandlerFunc func(w http.ResponseWriter, r *http.Request) error

//Define base validators used to redirect contect in the service to specific content
type WebServerValidator struct {
	//Regular Expression pointer that define images, video etc...
	ImageValidator *regexp.Regexp
	//Regular Expression pointer that define scripts, css, etc...
	ScriptsValidator *regexp.Regexp
	//Regular Expression pointer that define pages and types
	PageValidator *regexp.Regexp
	//Regular Expression pointer that define service (eg /{service_name})
	ServiceValidator *regexp.Regexp
}

var imagePath = regexp.MustCompile(".*\\.(gif|ico|png|svg|jpg|jpeg)$")
var scriptingPath = regexp.MustCompile(".*\\.(js|cgi|css)$")
var pagePath = regexp.MustCompile(".*\\.(template|htm|html)$")
var servicePath = regexp.MustCompile("/([a-zA-Z0-9_-]+)$")

type __webServerStruct struct {
	__running        bool
	Config           *common.ServerConfig
	Context          *ServerContext
	__logger         common.ServerLogger
	__notFoundAction common.HTTPAction
}

func writeFileToWriter(path string, file string, w http.ResponseWriter) error {
	filename := fmt.Sprintf("%s%s", path, file)
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	_, err = w.Write(bytes)
	return err
}

func (ws *__webServerStruct) __httpRequestHandler(w http.ResponseWriter, r *http.Request) {
	var path string = r.URL.Path
	var query url.Values = r.URL.Query()
	//	fmt.Println(fmt.Sprintf("Path: %s", path))
	//	fmt.Println(fmt.Sprintf("Query: %v", query))
	p := ws.Context.Validator.PageValidator.FindStringSubmatch(path)
	if p != nil {
		//Found Page call
		//		fmt.Println(fmt.Sprintf("Found PAGE Context for path: %s", path))
		var dataMap map[string]interface{} = ws.Context.Environment
		for k, v := range query {
			dataMap[k] = strings.Join(v, " ")
		}
		var newContext *ServerContext = &ServerContext{
			ServerPath:  ws.Context.ServerPath,
			Environment: dataMap,
		}
		title := path[1:]
		index := strings.LastIndex(title, ".")
		ext := ""
		if index > 0 {
			ext = title[index+1:]
			title = title[:index]
		}
		pg := WebPage{
			Title:     title,
			Extension: ext,
			Context:   newContext,
		}
		err := pg.Load()
		if err != nil {
			var message string = fmt.Sprintf("Error loading PAGE pattern : %s -> err = %s", path, err.Error())
			ws.__logger.Log(common.ERROR, message)
			if logger != nil {
				logger.ErrorS(message)
			} else {
				fmt.Println(message)
			}
			ws.__notFoundAction(w, r)
			return
		}
		err = pg.RenderOn(w)
		if err != nil {
			var message string = fmt.Sprintf("Error rendering PAGE pattern : %s -> err = %s", path, err.Error())
			ws.__logger.Log(common.ERROR, message)
			if logger != nil {
				logger.ErrorS(message)
			} else {
				fmt.Println(message)
			}
			ws.__notFoundAction(w, r)
			return
		}
	} else {
		i := ws.Context.Validator.ImageValidator.FindStringSubmatch(path)
		if i != nil {
			//Found image path
			//			fmt.Println(fmt.Sprintf("Found IMAGE Context for path: %s", path))
			err := writeFileToWriter(ws.Context.ServerPath, path, w)
			if err != nil {
				var message string = fmt.Sprintf("Error sending pattern : %s -> err = %s", path, err.Error())
				ws.__logger.Log(common.ERROR, message)
				if logger != nil {
					logger.ErrorS(message)
				} else {
					fmt.Println(message)
				}
				ws.__notFoundAction(w, r)
				return
			}
		} else {
			s := ws.Context.Validator.ScriptsValidator.FindStringSubmatch(path)
			if s != nil {
				//Found scripting path
				//				fmt.Println(fmt.Sprintf("Found SCRIPT Context for path: %s", path))
				err := writeFileToWriter(ws.Context.ServerPath, path, w)
				if err != nil {
					var message string = fmt.Sprintf("Error sending pattern : %s -> err = %s", path, err.Error())
					ws.__logger.Log(common.ERROR, message)
					if logger != nil {
						logger.ErrorS(message)
					} else {
						fmt.Println(message)
					}
					ws.__notFoundAction(w, r)
					return
				}
			} else {
				se := ws.Context.Validator.ServiceValidator.FindStringSubmatch(path)
				if se != nil {
					//Found service path
					//					fmt.Println(fmt.Sprintf("Found SERVICE Context for path: %s", path))
					var dataMap map[string]interface{} = ws.Context.Environment
					for k, v := range query {
						dataMap[k] = strings.Join(v, " ")
					}
					var newContext *ServerContext = &ServerContext{
						ServerPath:  ws.Context.ServerPath,
						Environment: dataMap,
					}
					title := path[1:]
					ext := "template"
					pg := WebPage{
						Title:     title,
						Extension: ext,
						Context:   newContext,
					}
					err := pg.Load()
					if err != nil {
						var message string = fmt.Sprintf("Error loading SERVICE pattern : %s -> err = %s", path, err.Error())
						ws.__logger.Log(common.ERROR, message)
						if logger != nil {
							logger.ErrorS(message)
						} else {
							fmt.Println(message)
						}
						ws.__notFoundAction(w, r)
						return
					}
					err = pg.RenderOn(w)
					if err != nil {
						var message string = fmt.Sprintf("Error rendering SERVICE pattern : %s -> err = %s", path, err.Error())
						ws.__logger.Log(common.ERROR, message)
						if logger != nil {
							logger.ErrorS(message)
						} else {
							fmt.Println(message)
						}
						ws.__notFoundAction(w, r)
						return
					}
				} else {
					var message string = fmt.Sprintf("Uknown type for pattern : %s", path)
					ws.__logger.Log(common.ERROR, message)
					if logger != nil {
						logger.ErrorS(message)
					} else {
						fmt.Println(message)
					}
					ws.__notFoundAction(w, r)
					return
				}
			}
		}
	}
}
func (ws *__webServerStruct) Open() error {
	var err error
	defer func() {
		itf := recover()
		if errs.IsError(itf) {
			err = itf.(error)
			ws.__logger.Log(common.ERROR, fmt.Sprintf("Error executing web server Open : %s", err.Error()))
		} else {
			err = errors.New(fmt.Sprintf("%v", itf))
			ws.__logger.Log(common.ERROR, fmt.Sprintf("Error executing web server Open : %v", err))
		}
	}()
	ws.__running = true
	ws.__logger.Open()
	http.HandleFunc("/", ws.__httpRequestHandler)
	http.ListenAndServe(fmt.Sprintf(":%v", ws.Config.ListeningPort), nil)
	return err
}
func (ws *__webServerStruct) Close() error {
	var err error
	defer func() {
		itf := recover()
		if errs.IsError(itf) {
			err = itf.(error)
			ws.__logger.Log(common.ERROR, fmt.Sprintf("Error executing web server Close : %s", err.Error()))
		} else {
			err = errors.New(fmt.Sprintf("%v", itf))
			ws.__logger.Log(common.ERROR, fmt.Sprintf("Error executing web server Close : %v", err))
		}
	}()
	ws.__logger.Close()
	ws.__running = false
	return err
}
func (ws *__webServerStruct) IsListening() bool {
	return ws.__running
}
func (ws *__webServerStruct) IsRunning() bool {
	return ws.__running
}
func (ws *__webServerStruct) Stream(interface{}) error {
	return errors.New("Not Implemeted!!")

}
func (ws *__webServerStruct) Receive() (interface{}, error) {
	return nil, errors.New("Not Implemeted!!")
}
func (ws *__webServerStruct) HandleConnOn(hf common.ServerHablerFunc) error {
	return errors.New("Not Implemeted!!")
}
func (ws *__webServerStruct) HandlingFuncs() []common.ServerHablerFunc {
	return []common.ServerHablerFunc{}
}
func (ws *__webServerStruct) RevokeHandler(hf common.ServerHablerFunc) error {
	return errors.New("Not Implemeted!!")
}
func (ws *__webServerStruct) Destroy() error {
	var err error
	defer func() {
		itf := recover()
		if errs.IsError(itf) {
			err = itf.(error)
			if ws.__logger != nil {
				ws.__logger.Log(common.ERROR, fmt.Sprintf("Error executing web server Destroy : %s", err.Error()))
			}
		} else {
			err = errors.New(fmt.Sprintf("%v", itf))
			if ws.__logger != nil {
				ws.__logger.Log(common.ERROR, fmt.Sprintf("Error executing web server Destroy : %v", err))
			}
		}
	}()
	if ws.__logger != nil {
		ws.__logger.Close()
		ws.__logger = nil
	}
	ws.Config = nil
	ws.Context = nil
	return err
}
func (ws *__webServerStruct) WaitFor() error {
	var err error
	defer func() {
		itf := recover()
		if errs.IsError(itf) {
			err = itf.(error)
			ws.__logger.Log(common.ERROR, fmt.Sprintf("Error executing web server WaitFor : %s", err.Error()))
		} else {
			err = errors.New(fmt.Sprintf("%v", itf))
			ws.__logger.Log(common.ERROR, fmt.Sprintf("Error executing web server WaitFor : %v", err))
		}
	}()
	for ws.__running {
		time.Sleep(1 * time.Second)
	}
	return err
}
func (ws *__webServerStruct) Clients() []common.ClientRef {
	return []common.ClientRef{}
}
func (ws *__webServerStruct) LogOn(channel *chan interface{}) error {
	return ws.__logger.AddOutChannel(channel)
}
func (ws *__webServerStruct) Logger() common.ServerLogger {
	return ws.__logger
}

// Define New Web server based on inpuit parameters
// Paarameters:
//    baseFolder (string) Base FS folder that container content sources
//    config (common.ServerConfig) Server initialization parameters group
//    env (map[string]interface{}) Map of environment items
//    validator (*WebServerValidator) Pointer to Web Server Service/Call Validator
//    notFoundAction (*common.HTTPAction) Pointer to 404 (Not found) state http handler
// Returns:
//    common.Server Web Server Instance
func NewWebServer(baseFolder string, config common.ServerConfig, env map[string]interface{}, validator *WebServerValidator, notFoundAction *common.HTTPAction) common.Server {
	if validator == nil {
		validator = &WebServerValidator{
			ImageValidator:   imagePath,
			PageValidator:    pagePath,
			ScriptsValidator: scriptingPath,
			ServiceValidator: servicePath,
		}
	} else {
		if validator.ImageValidator == nil {
			validator.ImageValidator = imagePath
		}
		if validator.PageValidator == nil {
			validator.PageValidator = pagePath
		}
		if validator.ScriptsValidator == nil {
			validator.ScriptsValidator = scriptingPath
		}
		if validator.ServiceValidator == nil {
			validator.ServiceValidator = servicePath
		}
	}
	var action common.HTTPAction
	if notFoundAction == nil {
		action = common.DefaultNotFoundAction
	} else {
		action = *notFoundAction
	}
	context := &ServerContext{
		ServerPath:  baseFolder,
		Environment: env,
		Validator:   validator,
	}
	logger := common.NewServerLogger(config.LogLevel)
	return &__webServerStruct{
		Config:           &config,
		Context:          context,
		__logger:         logger,
		__notFoundAction: action,
	}
}
