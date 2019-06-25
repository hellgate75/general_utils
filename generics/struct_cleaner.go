package generics

import (
	"fmt"
)

type StructureCleaner interface {
	CleanupInterfaceArray(in []interface{}) []interface{}
	CleanupInterfaceMap(in map[interface{}]interface{}) map[string]interface{}
	CleanupMapValue(v interface{}) interface{}
	MapToInterface(v map[interface{}]interface{}) interface{}
	CleanupInterfaceObject(in interface{}) interface{}
}

type _cleanerStruct struct{}

func (c *_cleanerStruct) CleanupInterfaceArray(in []interface{}) []interface{} {
	res := make([]interface{}, len(in))
	for i, v := range in {
		res[i] = c.CleanupMapValue(v)
	}
	return res
}

func (c *_cleanerStruct) CleanupInterfaceMap(in map[interface{}]interface{}) map[string]interface{} {
	res := make(map[string]interface{})
	for k, v := range in {
		res[fmt.Sprintf("%v", k)] = c.CleanupMapValue(v)
	}
	return res
}

func (c *_cleanerStruct) CleanupInterfaceObject(in interface{}) interface{} {
	//	res := in
	//	for k, v := range in {
	//		res[fmt.Sprintf("%v", k)] = c.CleanupMapValue(v)
	//	}
	//	return res
	return in
}

func (c *_cleanerStruct) CleanupMapValue(v interface{}) interface{} {
	switch v := v.(type) {
	case []interface{}:
		return c.CleanupInterfaceArray(v)
	case map[interface{}]interface{}:
		return c.CleanupInterfaceMap(v)
		//	case string, bool, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr, byte, rune, float32, float64, complex64, complex128:
	case string, bool, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr, float32, float64, complex64, complex128:
		return v
	default:
		return c.CleanupInterfaceObject(v)
	}
}

func (c *_cleanerStruct) MapToInterface(in map[interface{}]interface{}) interface{} {
	var out interface{}
	//	for k, v := range in {
	//		out[fmt.Sprintf("%v", k)] = c.CleanupMapValue(v)
	//	}
	return out
}

func NewStructureCleaner() StructureCleaner {
	return &_cleanerStruct{}
}
