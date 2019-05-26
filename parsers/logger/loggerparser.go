package parsers

import (
	"github.com/hellgate75/general_utils/log"
	"github.com/hellgate75/general_utils/parsers"
)

var conf log.LogConfig = log.LogConfig{}

// Provides read access to log configuration and parse from a File.
//
// Parameters:
//   encoding (parsers.Encoding) File Encoding type (JSON, XML, YAML)
//   filePath (string) File full path
//
// Returns:
// error Any suitable error during code execution
func InitStaticFromFile(encoding parsers.Encoding, filePath string) error {
	var parser parsers.Parser
	var err error
	parser, err = parsers.New(encoding)

	if err != nil {
		return err
	}
	err = parser.DeserializeFromFile(filePath, conf)
	return err
}

// Retrieve loaded log configuration.
//
// Returns:
// LogConfig Loaded Log Configuration
func GetLogConfig() log.LogConfig {
	return conf
}
