package logger

import (
	"github.com/hellgate75/general_utils/log"
	"github.com/hellgate75/general_utils/parsers"
)

// Provides read access to log configuration and parse from a File.
//
// Parameters:
//   encoding (parsers.Encoding) File Encoding type (JSON, XML, YAML)
//   filePath (string) File full path
//
// Returns:
// error Any suitable error risen during code execution
func LoadLoggerConfigFromFile(encoding parsers.Encoding, filePath string) (*log.LogConfig, error) {
	var parser parsers.Parser
	var err error
	parser, err = parsers.New(encoding)

	if err != nil {
		return nil, err
	}
	var conf log.LogConfig = log.LogConfig{}
	//TODO: Complete Remapping structure
	err = parser.DeserializeFromFile(filePath, conf)
	return &conf, err
}

var conf log.LogConfig = log.LogConfig{}

// Set loaded log configuration.
//
// Paraeters:
// config (log.LogConfig) Loaded Log Configuration
func SetLogConfig(config log.LogConfig) {
	conf = config
}

// Retrieve loaded log configuration.
//
// Returns:
// log.LogConfig Loaded Log Configuration
func GetLogConfig() log.LogConfig {
	return conf
}
