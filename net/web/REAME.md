## package web // import "github.com\\hellgate75\\general_utils\\net\\web"


### VARIABLES

##### var HandlerFunc func(w http.ResponseWriter, r *http.Request) error
    Http Response Handler element


### FUNCTIONS

#### func InitLogger()
     Initialize package logger if not started

#### func NewWebServer(baseFolder string, config common.ServerConfig, env map[string]interface{}, validator *WebServerValidator, notFoundAction *common.HTTPAction) common.Server
    Define New Web server based on inpuit parameters 
#####     Parameters:
       baseFolder (string) Base FS folder that container content sources
       config (common.ServerConfig) Server initialization parameters group
       env (map[string]interface{}) Map of environment items
       validator (*WebServerValidator) Pointer to Web Server Service/Call Validator
       notFoundAction (*common.HTTPAction) Pointer to 404 (Not found) state http handler
#####    Returns:

       common.Server Web Server Instance


### TYPES

##### type ServerContext struct {
##### 	ServerPath  string
##### 	Environment map[string]interface{}
##### 	Validator   *WebServerValidator
##### }
	Server Context element

##### type WebPage struct {
##### 	//Page title
##### 	Title string
##### 	//Page extension
##### 	Extension string
##### 	//Page content bytes
##### 	Body []byte
##### 	//Server Configuration Context
##### 	Context *ServerContext
##### }
    Web Page Element

#### func LoadWebPageFromFile(filename string, context *ServerContext) (*WebPage, error)
    Load web page from file 
#####     Returns:
       (*WebPage Web Page pointer or nil in case of error,
       error Any error that can occurs during computation)

#### func (p *WebPage) Load() error
    Save web page from server path file
##### 	Returns:
       error Any error that can occurs during computation

#### func (p *WebPage) Render() (string, error)
    Render web page from server path file or pre-loaded bytes 
#####     Returns:
       (string template parsed page text,
       error Any error that can occurs during computation)

#### func (p *WebPage) RenderOn(wr io.Writer) error
    Render web page from server path file or pre-loaded bytes to an external
    writer
##### 	Returns:
       error Any error that can occurs during computation

#### func (p *WebPage) Save() error
    Save web page to server path file
##### 	Returns:
       error Any error that can occurs during computation

##### type WebServerValidator struct {
##### 	//Regular Expression pointer that define images, video etc...
##### 	ImageValidator *regexp.Regexp
##### 	//Regular Expression pointer that define scripts, css, etc...
##### 	ScriptsValidator *regexp.Regexp
##### 	//Regular Expression pointer that define pages and types
##### 	PageValidator *regexp.Regexp
##### 	//Regular Expression pointer that define service (eg /{service_name})
##### 	ServiceValidator *regexp.Regexp
##### }
    Define base validators used to redirect contect in the service to specific
    content

