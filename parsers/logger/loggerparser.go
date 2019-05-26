package parsers

import (
	"github.com/hellgate75/general_utils/parsers"
)

var conf LogConfig = LogConfig{}

func InitStaticFromFile(encoding parsers.Encoding, filePath string) error {
	var parser parsers.Parser
	var err error
	parser, err = parsers.New(encoding)

	if err != nil {
		return err
	}
	//TODO Complete and fix issues
	//err = parser.DeserializeFromFile(filePath, &conf)
	return err
}

func GetLogConfig() LogConfig {
	return conf
}
