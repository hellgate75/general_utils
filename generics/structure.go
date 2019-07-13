package generics

import (
	"errors"
	"fmt"
	errs "github.com/hellgate75/general_utils/errors"
	"reflect"
)

// Generic Serializable Type
type Serializable interface {
	//Retrive attributes map
	GetAsMap() map[string]interface{}
	//Set attributes from attributes map map
	SetFromMap(map[string]interface{}) (interface{}, error)
}

type SerializableStruct struct {
}

func (this *SerializableStruct) GetAsMap() map[string]interface{} {
	return GetAsMapImpl(this)
}

func (this *SerializableStruct) SetFromMap(attributes map[string]interface{}) (interface{}, error) {
	return SetFromMapImpl(this, attributes)
}

// Type containing most relevant reflection information for any of the attributes in your interface
type FieldConfig struct {
	Name           string
	Type           reflect.Type
	Value          reflect.Value
	IsNil          bool
	FieldStructure reflect.StructField
}

// Implement Serializable#GetAsMap standard solution
//
// Parameters:
//   yourInterface (interface{} NOT POINTER) recover information from your interface and return the attributes map
//
// Returns:
//   map[string]interface{} That represent the attributes of the interface as map
func GetAsMapImpl(yourInterface interface{}) map[string]interface{} {
	var attributes map[string]interface{} = make(map[string]interface{})
	objVal := reflect.ValueOf(yourInterface)
	//	fmt.Println(objVal.Type().Kind().String())
	noFlds := objVal.Type().NumField()
	for i := 0; i < noFlds; i++ {
		fld := objVal.Type().Field(i)
		k := fld.Name
		//		fmt.Println(objVal.Field(i).Type().Kind().String())
		v := objVal.Field(i).Interface()
		attributes[fmt.Sprintf("%v", k)] = v
	}
	return attributes
}

// Extract FieldConfig from interace is not Nil, and it's a structure
//
// Parameters:
//    element (interface{}) Element we want inspect into
//
// Returns:
//   []FieldConfig - Field configuration list accordingly to interface structure

func PrepareFileStructList(element interface{}) ([]FieldConfig, error) {
	var err error
	defer func() {
		itf := recover()
		if errs.IsError(itf) {
			err = itf.(error)
			if logger != nil {
				logger.ErrorS(fmt.Sprintf("Error in PrepareFileStructList, message : %s", err.Error()))
			}

		} else {
			err = errors.New(fmt.Sprintf("%v", itf))
			if logger != nil {
				logger.ErrorS(fmt.Sprintf("Error in PrepareFileStructList, message : %v", itf))
			}
		}
	}()
	var configList []FieldConfig = []FieldConfig{}
	if element != nil && reflect.ValueOf(element).Kind() == reflect.Struct {
		elemRef := reflect.Indirect(reflect.ValueOf(element))
		for i := 0; i < elemRef.NumField(); i++ {
			fs := elemRef.Type().Field(i)
			configList = append(configList, FieldConfig{
				Name:           fs.Name,
				Type:           fs.Type,
				Value:          elemRef.Field(i),
				IsNil:          false,
				FieldStructure: fs,
			})
		}
	}
	return configList, err
}

// Extract FieldConfig from interace is not Nil and it's a structure
//
// Parameters:
//    element (interface{}) Element we want inspect into
//
// Returns:
//   map[string]FieldConfig - Field configuration map by field name, accordingly to interface structure

func PrepareFileStructMap(element interface{}) (map[string]FieldConfig, error) {
	var configMap map[string]FieldConfig = make(map[string]FieldConfig)
	var err error
	defer func() {
		itf := recover()
		if logger != nil {
			if errs.IsError(itf) {
				err = itf.(error)
				if logger != nil {
					logger.ErrorS(fmt.Sprintf("Error in PrepareFileStructMap, message : %s", err.Error()))
				}
			} else {
				err = errors.New(fmt.Sprintf("%v", itf))
				if logger != nil {
					logger.ErrorS(fmt.Sprintf("Error in PrepareFileStructMap, message : %v", itf))
				}
			}
		}
	}()
	if element != nil && reflect.ValueOf(element).Kind() == reflect.Struct {
		elemRef := reflect.Indirect(reflect.ValueOf(element))
		for i := 0; i < elemRef.NumField(); i++ {
			fs := elemRef.Type().Field(i)
			configMap[fs.Name] = FieldConfig{
				Name:           fs.Name,
				Type:           fs.Type,
				Value:          elemRef.Field(i),
				IsNil:          false,
				FieldStructure: fs,
			}
		}
	}
	return configMap, err
}

// Implement Serializable#SetFromMap standard solution
//
// Parameters:
//   yourInterface (interface{} pointer) set attributes into your interface
//   attributes (map[string]interface{}) Map of attribute names and values to set up into the interface
//
// Returns:
//   error Any arror that can occur during the computation
func SetFromMapImpl(yourInterface interface{}, attributes map[string]interface{}) (interface{}, error) {
	//	var fieldList []FieldConfig
	//	var wrongFieldList []FieldConfig
	//	var discardedFieldList []FieldConfig
	//	var yourValue reflect.Value = reflect.Indirect(reflect.ValueOf(yourInterface))
	//	var yourElement reflect.Value = reflect.Indirect(reflect.ValueOf(yourInterface)).Elem()
	yourElement := reflect.ValueOf(&yourInterface).Elem()
	//	var yourType reflect.Type = yourValue.Type()
	//	var attributesMap []reflect.StructField
	var err error
	defer func() {
		itf := recover()
		if errs.IsError(itf) {
			err = itf.(error)
			if logger != nil {
				logger.ErrorS(fmt.Sprintf("Error in SetFromMapImpl(), message : %s", err.Error()))
			}
		} else {
			err = errors.New(fmt.Sprintf("%v", itf))
			if logger != nil {
				logger.ErrorS(fmt.Sprintf("Error in SetFromMapImpl(), message : %v", itf))
			}
		}
	}()
	for k, v := range attributes {
		if fd := yourElement.FieldByName(k); fd.IsValid() {
			if !fd.CanSet() {
				return yourInterface, errors.New(fmt.Sprintf("Field Name '%s' cannot be setted up!!!", fd.Type().Name()))
			}
			fd.Set(reflect.ValueOf(v))
		}
	}
	return yourInterface, err
}

// Returns a new Emty serializable object instance
//
// Returns:
//   Serializable - Empty Serializable interface instance
func NewSerializable() Serializable {
	return &SerializableStruct{}
}
