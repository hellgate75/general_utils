package common

import (
	"fmt"
	"strings"
	"testing"
)

func TestStringToStorageUnit(t *testing.T) {
	var val StorageUnit
	var err error
	val, err = StringToStorageUnit("B")
	if err != nil {
		t.Fatal("Error parsing Byte: 'B' : ", err.Error())
	} else if val != Bt {
		t.Fatal(fmt.Sprintf("Wrong parsing of Byte: 'B' : expected <%v> insted found <%v> !!", Bt, val))
	}
	val, err = StringToStorageUnit("KB")
	if err != nil {
		t.Fatal("Error parsing KiloByte: 'KB' : ", err.Error())
	} else if val != Kbt {
		t.Fatal(fmt.Sprintf("Wrong parsing of KiloByte: 'KB' : expected <%v> insted found <%v> !!", Kbt, val))
	}
	val, err = StringToStorageUnit("MB")
	if err != nil {
		t.Fatal("Error parsing MegaByte: 'MB' : ", err.Error())
	} else if val != Mbt {
		t.Fatal(fmt.Sprintf("Wrong parsing of MegaByte: 'MB' : expected <%v> insted found <%v> !!", Mbt, val))
	}
	val, err = StringToStorageUnit("GB")
	if err != nil {
		t.Fatal("Error parsing GigaByte: 'GB' : ", err.Error())
	} else if val != Gbt {
		t.Fatal(fmt.Sprintf("Wrong parsing of GigaByte: 'GB' : expected <%v> insted found <%v> !!", Gbt, val))
	}
	val, err = StringToStorageUnit("TB")
	if err != nil {
		t.Fatal("Error parsing TeraByte: 'TB' : ", err.Error())
	} else if val != Tbt {
		t.Fatal(fmt.Sprintf("Wrong parsing of TeraByte: 'TB' : expected <%v> insted found <%v> !!", Tbt, val))
	}
	val, err = StringToStorageUnit("UB")
	if err == nil {
		t.Fatal("Error due to parsing of UltraByte: 'UB' because ultra bytes unit doesn't exist ")
	}
	val, err = StringToStorageUnit("")
	if err == nil {
		t.Fatal("Error due to parsing of EmptyByte: '' because empty bytes unit doesn't exist ")
	}
}

func TestStorageUnitToString(t *testing.T) {
	var val string
	var err error
	val, err = StorageUnitToString(Bt)
	if err != nil {
		t.Fatal("Error parsing Byte: 'B' : ", err.Error())
	} else if strings.ToUpper(val) != "B" {
		t.Fatal(fmt.Sprintf("Wrong parsing of Byte: expected <%s> insted found <%s> !!", "B", val))
	}
	val, err = StorageUnitToString(Kbt)
	if err != nil {
		t.Fatal("Error parsing KiloByte: 'KB' : ", err.Error())
	} else if strings.ToUpper(val) != "KB" {
		t.Fatal(fmt.Sprintf("Wrong parsing of KiloByte: 'KB' : expected <%s> insted found <%s> !!", "KB", val))
	}
	val, err = StorageUnitToString(Mbt)
	if err != nil {
		t.Fatal("Error parsing MegaByte: 'MB' : ", err.Error())
	} else if strings.ToUpper(val) != "MB" {
		t.Fatal(fmt.Sprintf("Wrong parsing of MegaByte: 'MB' : expected <%s> insted found <%s> !!", "MB", val))
	}
	val, err = StorageUnitToString(Gbt)
	if err != nil {
		t.Fatal("Error parsing GigaByte: 'GB' : ", err.Error())
	} else if strings.ToUpper(val) != "GB" {
		t.Fatal(fmt.Sprintf("Wrong parsing of GigaByte: 'GB' : expected <%s> insted found <%s> !!", "GB", val))
	}
	val, err = StorageUnitToString(Tbt)
	if err != nil {
		t.Fatal("Error parsing TeraByte: 'TB' : ", err.Error())
	} else if strings.ToUpper(val) != "TB" {
		t.Fatal(fmt.Sprintf("Wrong parsing of TeraByte: 'TB' : expected <%s> insted found <%s> !!", "TB", val))
	}
	var wrongSU StorageUnit = 1000
	val, err = StorageUnitToString(wrongSU)
	if err == nil {
		t.Fatal(fmt.Sprintf("Error due to parsing of Wrong Storage Unit: '%v' because ultra bytes unit doesn't exist ", wrongSU))
	}
}

func TestStorageUnitBytesFactor(t *testing.T) {
	var val int64
	var err error
	val, err = StorageUnitBytesFactor(Bt)
	if err != nil {
		t.Fatal("Error calling function for Byte: 'B' : ", err.Error())
	} else if val != 1 {
		t.Fatal(fmt.Sprintf("Wrong call of function for Byte: expected <%d> insted found <%d> !!", 1, val))
	}
	val, err = StorageUnitBytesFactor(Kbt)
	if err != nil {
		t.Fatal("Error calling function for KiloByte: 'KB' : ", err.Error())
	} else if val != 1024 {
		t.Fatal(fmt.Sprintf("Wrong call of function for KiloByte: 'KB' : expected <%d> insted found <%d> !!", 1024, val))
	}
	val, err = StorageUnitBytesFactor(Mbt)
	if err != nil {
		t.Fatal("Error calling function for MegaByte: 'MB' : ", err.Error())
	} else if val != 1048576 {
		t.Fatal(fmt.Sprintf("Wrong call of function for MegaByte: 'MB' : expected <%d> insted found <%d> !!", 1048576, val))
	}
	val, err = StorageUnitBytesFactor(Gbt)
	if err != nil {
		t.Fatal("Error calling function for GigaByte: 'GB' : ", err.Error())
	} else if val != 1073741824 {
		t.Fatal(fmt.Sprintf("Wrong call of function for GigaByte: 'GB' : expected <%d> insted found <%d> !!", 1073741824, val))
	}
	val, err = StorageUnitBytesFactor(Tbt)
	if err != nil {
		t.Fatal("Error calling function for TeraByte: 'TB' : ", err.Error())
	} else if val != 1099511627776 {
		t.Fatal(fmt.Sprintf("Wrong call of function for TeraByte: 'TB' : expected <%d> insted found <%d> !!", 1099511627776, val))
	}
	var wrongSU StorageUnit = 1000
	val, err = StorageUnitBytesFactor(wrongSU)
	if err == nil {
		t.Fatal(fmt.Sprintf("Error due to calling function for Wrong Storage Unit: '%v' because ultra bytes unit doesn't exist ", wrongSU))
	}
}

func TestStringToStreamInOutFormat(t *testing.T) {
	var val StreamInOutFormat
	var err error
	val, err = StringToStreamInOutFormat("text/plain")
	if err != nil {
		t.Fatal(fmt.Sprintf("Error calling for : '%s' : ", "text/plain"), err.Error())
	} else if val != PlainTextFormat {
		t.Fatal(fmt.Sprintf("Wrong call of function for: '%s' : expected <%v> insted found <%v> !!", "text/plain", PlainTextFormat, val))
	}
	val, err = StringToStreamInOutFormat("application/json")
	if err != nil {
		t.Fatal(fmt.Sprintf("Error calling for : '%s' : ", "application/json"), err.Error())
	} else if val != JsonFormat {
		t.Fatal(fmt.Sprintf("Wrong call of function for: '%s' : expected <%v> insted found <%v> !!", "application/json", JsonFormat, val))
	}
	val, err = StringToStreamInOutFormat("application/yaml")
	if err != nil {
		t.Fatal(fmt.Sprintf("Error calling for : '%s' : ", "application/yaml"), err.Error())
	} else if val != YamlFormat {
		t.Fatal(fmt.Sprintf("Wrong call of function for: '%s' : expected <%v> insted found <%v> !!", "application/yaml", YamlFormat, val))
	}
	val, err = StringToStreamInOutFormat("application/xml")
	if err != nil {
		t.Fatal(fmt.Sprintf("Error calling for : '%s' : ", "application/xml"), err.Error())
	} else if val != XmlFormat {
		t.Fatal(fmt.Sprintf("Wrong call of function for: '%s' : expected <%v> insted found <%v> !!", "application/xml", XmlFormat, val))
	}
	val, err = StringToStreamInOutFormat("application/base64")
	if err != nil {
		t.Fatal(fmt.Sprintf("Error calling for : '%s' : ", "application/base64"), err.Error())
	} else if val != EncryptedFormat {
		t.Fatal(fmt.Sprintf("Wrong call of function for: '%s' : expected <%v> insted found <%v> !!", "application/base64", EncryptedFormat, val))
	}
	val, err = StringToStreamInOutFormat("application/binary")
	if err != nil {
		t.Fatal(fmt.Sprintf("Error calling for : '%s' : ", "application/binary"), err.Error())
	} else if val != BinaryFormat {
		t.Fatal(fmt.Sprintf("Wrong call of function for: '%s' : expected <%v> insted found <%v> !!", "application/binary", BinaryFormat, val))
	}
	val, err = StringToStreamInOutFormat("")
	if err == nil {
		t.Fatal(fmt.Sprintf("Error not occurred calling for : '%s' : ", "<empty>"))
	} else if val != StreamInOutFormat(0) {
		t.Fatal(fmt.Sprintf("Wrong call of function for: '%s' : expected <%v> insted found <%v> !!", "<empty>", StreamInOutFormat(0), val))
	}
	val, err = StringToStreamInOutFormat("application/notexistingformat")
	if err == nil {
		t.Fatal(fmt.Sprintf("Error not occurred calling for : '%s' : ", "application/notexistingformat"))
	} else if val != StreamInOutFormat(0) {
		t.Fatal(fmt.Sprintf("Wrong call of function for: '%s' : expected <%v> insted found <%v> !!", "application/notexistingformat", StreamInOutFormat(0), val))
	}
}

func TestStreamInOutFormatToString(t *testing.T) {
	var val string
	var err error
	val, err = StreamInOutFormatToString(PlainTextFormat)
	if err != nil {
		t.Fatal(fmt.Sprintf("Error calling for : '%s' : ", "text/plain"), err.Error())
	} else if val != "text/plain" {
		t.Fatal(fmt.Sprintf("Wrong call of function for: '%s' : expected <%v> insted found <%v> !!", "PlainTextFormat", "text/plain", val))
	}
	val, err = StreamInOutFormatToString(JsonFormat)
	if err != nil {
		t.Fatal(fmt.Sprintf("Error calling for : '%s' : ", "application/json"), err.Error())
	} else if val != "application/json" {
		t.Fatal(fmt.Sprintf("Wrong call of function for: '%s' : expected <%v> insted found <%v> !!", "JsonFormat", "application/json", val))
	}
	val, err = StreamInOutFormatToString(YamlFormat)
	if err != nil {
		t.Fatal(fmt.Sprintf("Error calling for : '%s' : ", "application/yaml"), err.Error())
	} else if val != "application/yaml" {
		t.Fatal(fmt.Sprintf("Wrong call of function for: '%s' : expected <%v> insted found <%v> !!", "YamlFormat", "application/yaml", val))
	}
	val, err = StreamInOutFormatToString(XmlFormat)
	if err != nil {
		t.Fatal(fmt.Sprintf("Error calling for : '%s' : ", "application/xml"), err.Error())
	} else if val != "application/xml" {
		t.Fatal(fmt.Sprintf("Wrong call of function for: '%s' : expected <%v> insted found <%v> !!", "XmlFormat", "application/xml", val))
	}
	val, err = StreamInOutFormatToString(EncryptedFormat)
	if err != nil {
		t.Fatal(fmt.Sprintf("Error calling for : '%s' : ", "application/base64"), err.Error())
	} else if val != "application/base64" {
		t.Fatal(fmt.Sprintf("Wrong call of function for: '%s' : expected <%v> insted found <%v> !!", "EncryptedFormat", "application/base64", val))
	}
	val, err = StreamInOutFormatToString(BinaryFormat)
	if err != nil {
		t.Fatal(fmt.Sprintf("Error calling for : '%s' : ", "application/binary"), err.Error())
	} else if val != "application/binary" {
		t.Fatal(fmt.Sprintf("Wrong call of function for: '%s' : expected <%v> insted found <%v> !!", "BinaryFormat", "application/binary", val))
	}
	val, err = StreamInOutFormatToString(-1)
	if err == nil {
		t.Fatal(fmt.Sprintf("Error not occurred calling for : '%s' : ", "<negative value>"))
	} else if val != "" {
		t.Fatal(fmt.Sprintf("Wrong call of function for: '%s' : expected <%v> insted found <%v> !!", "<negative value>", StreamInOutFormat(0), val))
	}
	val, err = StreamInOutFormatToString(StreamInOutFormat(1000))
	if err == nil {
		t.Fatal(fmt.Sprintf("Error not occurred calling for : '%s' ", "<worng value>"))
	} else if val != "" {
		t.Fatal(fmt.Sprintf("Wrong call of function for: '%s' : expected <%v> insted found <%v> !!", "<worng value", StreamInOutFormat(0), val))
	}
}

func TestStringToWriterType(t *testing.T) {
	var val WriterType
	var err error
	val, err = StringToWriterType("stdout")
	if err != nil {
		t.Fatal(fmt.Sprintf("Error calling for : '%s' : ", "StdOutWriter"), err.Error())
	} else if val != StdOutWriter {
		t.Fatal(fmt.Sprintf("Wrong call of function for: '%s' : expected <%s> insted found <%s> !!", "StdOutWriter", StdOutWriter, val))
	}
	val, err = StringToWriterType("file")
	if err != nil {
		t.Fatal(fmt.Sprintf("Error calling for : '%s' : ", "FileWriter"), err.Error())
	} else if val != FileWriter {
		t.Fatal(fmt.Sprintf("Wrong call of function for: '%s' : expected <%s> insted found <%s> !!", "FileWriter", FileWriter, val))
	}
	val, err = StringToWriterType("url")
	if err != nil {
		t.Fatal(fmt.Sprintf("Error calling for : '%s' : ", "UrlWriter"), err.Error())
	} else if val != UrlWriter {
		t.Fatal(fmt.Sprintf("Wrong call of function for: '%s' : expected <%s> insted found <%s> !!", "UrlWriter", UrlWriter, val))
	}
	val, err = StringToWriterType("")
	if err == nil {
		t.Fatal(fmt.Sprintf("Error not occurred calling for : '%s' : ", "<empty>"))
	} else if val != WriterType("") {
		t.Fatal(fmt.Sprintf("Wrong call of function for: '%s' : expected <%s> insted found <%s> !!", "<empty>", WriterType(""), val))
	}
	val, err = StringToWriterType("notexistingtype")
	if err == nil {
		t.Fatal(fmt.Sprintf("Error not occurred calling for : '%s' : ", "notexistingtype"))
	} else if val != WriterType("") {
		t.Fatal(fmt.Sprintf("Wrong call of function for: '%s' : expected <%s> insted found <%s> !!", "notexistingtype", WriterType(""), val))
	}
}

func TestWriterTypeToString(t *testing.T) {
	var val string
	var err error
	val, err = WriterTypeToString(StdOutWriter)
	if err != nil {
		t.Fatal(fmt.Sprintf("Error calling for : '%s' : ", "StdOutWriter"), err.Error())
	} else if val != "StdOut" {
		t.Fatal(fmt.Sprintf("Wrong call of function for: '%s' : expected <%s> insted found <%s> !!", "StdOutWriter", "StdOut", val))
	}
	val, err = WriterTypeToString(FileWriter)
	if err != nil {
		t.Fatal(fmt.Sprintf("Error calling for : '%s' : ", "FileWriter"), err.Error())
	} else if val != "File" {
		t.Fatal(fmt.Sprintf("Wrong call of function for: '%s' : expected <%s> insted found <%s> !!", "FileWriter", "File", val))
	}
	val, err = WriterTypeToString(UrlWriter)
	if err != nil {
		t.Fatal(fmt.Sprintf("Error calling for : '%s' : ", "UrlWriter"), err.Error())
	} else if val != "Url" {
		t.Fatal(fmt.Sprintf("Wrong call of function for: '%s' : expected <%s> insted found <%s> !!", "UrlWriter", "Url", val))
	}
	val, err = WriterTypeToString(WriterType(""))
	if err == nil {
		t.Fatal(fmt.Sprintf("Error not occurred calling for : '%s' : ", "<empty>"))
	} else if val != "" {
		t.Fatal(fmt.Sprintf("Wrong call of function for: '%s' : expected <%s> insted found <%s> !!", "<empty>", "", val))
	}
	val, err = WriterTypeToString(WriterType("notexistingtype"))
	if err == nil {
		t.Fatal(fmt.Sprintf("Error not occurred calling for : '%s' : ", "notexistingtype"))
	} else if val != "" {
		t.Fatal(fmt.Sprintf("Wrong call of function for: '%s' : expected <%s> insted found <%s> !!", "notexistingtype", "", val))
	}
}
