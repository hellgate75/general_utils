package common

import (
	"errors"
	"strings"
)

// Generic Type
type Type interface{}

//Storage Unit representing type
type StorageUnit int

const (
	//bytes = 1b
	Bt StorageUnit = 1
	//KiloBytes = 1024b
	Kbt StorageUnit = 2
	//MegaBytes = 1048576b
	Mbt StorageUnit = 3
	//GigaBytes = 1073741824b
	Gbt StorageUnit = 4
	//TeraBytes = 1099511627776b
	Tbt StorageUnit = 5
)

// Transform text in Storage Unit.
//
// Parameters:
//   text (string) Text to convert
//
// Returns:
// common.StorageUnit Storage Unit representing the text data or 0 anyway
// error Any suitable error risen during code execution
func StringToStorageUnit(text string) (StorageUnit, error) {
	var testText = strings.ToUpper(strings.TrimSpace(text))
	if testText == "" {
		return 0, errors.New("common :: StringToStorageUnit : Empty input text")
	}
	switch testText {
	case "B":
		return Bt, nil
	case "KB":
		return Kbt, nil
	case "MB":
		return Mbt, nil
	case "GB":
		return Gbt, nil
	case "TB":
		return Tbt, nil
	}

	return 0, errors.New("common :: StringToStorageUnit : Storage Unit is undefined")
}

// Transform Storage Unit in representing text.
//
// Parameters:
//   unit (common.StorageUnit) Storage Unit
//
// Returns:
// string Text representing the Storage Unit or 0 anyway
// error Any suitable error risen during code execution
func StorageUnitToString(unit StorageUnit) (string, error) {
	if unit < Bt || unit > Tbt {
		return "", errors.New("common :: StorageUnitToString : Storage Unit out of range")
	}
	switch unit {
	case Bt:
		return "b", nil
	case Kbt:
		return "Kb", nil
	case Mbt:
		return "Mb", nil
	case Gbt:
		return "Gb", nil
	case Tbt:
		return "Tb", nil
	}

	return "", errors.New("common :: StorageUnitToString : Storage Unit is undefined")
}

// Returns Storage Unit representin factor in number of bytes.
//
// Parameters:
//   unit (common.StorageUnit) Reference Storage Unit
//
// Returns:
// int64 Factor of conversion to bytes or 0 anyway
// error Any suitable error risen during code execution
func StorageUnitBytesFactor(unit StorageUnit) (int64, error) {
	if unit < Bt || unit > Tbt {
		return 0, errors.New("common :: StorageUnitMultiplier : Storage Unit out of range")
	}
	switch unit {
	case Bt:
		return 1, nil
	case Kbt:
		return 1024, nil
	case Mbt:
		return 1048576, nil
	case Gbt:
		return 1073741824, nil
	case Tbt:
		return 1099511627776, nil
	}

	return 0, errors.New("common :: StorageUnitMultiplier : Storage Unit is undefined")
}

//Storage Unit representing type
type StreamInOutFormat int

const (
	//Plain text input/output format
	PlainTextFormat StreamInOutFormat = 11
	//Json input/output format
	JsonFormat StreamInOutFormat = 12
	//Yaml input/output format
	YamlFormat StreamInOutFormat = 13
	//Xml input/output format
	XmlFormat StreamInOutFormat = 14
	//Base64 encrypted text input/output format
	EncryptedFormat StreamInOutFormat = 15
	//Text encrypted input/output format
	BinaryFormat StreamInOutFormat = 16
)

// Transform text in Stream Input/Output format.
//
// Parameters:
//   text (string) Text to convert
//
// Returns:
// common.StreamInOutFormat Stream Input/Output format representing the text data or 0 anyway
// error Any suitable error risen during code execution
func StringToStreamInOutFormat(text string) (StreamInOutFormat, error) {
	var testText = strings.ToLower(strings.TrimSpace(text))
	if testText == "" {
		return 0, errors.New("common :: StringToStreamInOutFormat : Empty input text")
	}
	switch testText {
	case "text/plain":
		return PlainTextFormat, nil
	case "application/json":
		return JsonFormat, nil
	case "application/yaml":
		return YamlFormat, nil
	case "application/xml":
		return XmlFormat, nil
	case "application/base64":
		return EncryptedFormat, nil
	case "application/binary":
		return BinaryFormat, nil
	}

	return 0, errors.New("common :: StringToStreamInOutFormat : Stream Input/Output format is undefined")
}

// Transform Stream Input/Output format in representing text.
//
// Parameters:
//   format (common.StreamInOutFormat) Stream Input/Output format
//
// Returns:
// string Text representing the Stream Input/Output format or 0 anyway
// error Any suitable error risen during code execution
func StreamInOutFormatToString(format StreamInOutFormat) (string, error) {
	if format < PlainTextFormat || format > BinaryFormat {
		return "", errors.New("common :: StorageUnitToString : Stream Input/Output format out of range")
	}
	switch format {
	case PlainTextFormat:
		return "text/plain", nil
	case JsonFormat:
		return "application/json", nil
	case YamlFormat:
		return "application/yaml", nil
	case XmlFormat:
		return "application/xml", nil
	case EncryptedFormat:
		return "application/base64", nil
	case BinaryFormat:
		return "application/binary", nil
	}

	return "", errors.New("common :: StringToStreamInOutFormat : Stream Input/Output format is undefined")
}
