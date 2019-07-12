package generics

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

type __serializableTestStruct struct {
	Name      string
	Age       int
	Birthdate time.Time
}

func (this *__serializableTestStruct) GetAsMap() map[string]interface{} {
	return GetAsMapImpl(*this)
}

func (this *__serializableTestStruct) SetFromMap(attributes map[string]interface{}) (interface{}, error) {
	return SetFromMapImpl(*this, attributes)
}

func TestNewSerializable(t *testing.T) {
	serializable := NewSerializable()
	m1, ok1 := reflect.ValueOf(serializable).Type().MethodByName("GetAsMap")
	if !ok1 {
		t.Fatal("Error checking GetAsMap method into the default serializable interface")
	}
	if m1.Type.NumOut() != 1 {
		t.Fatal(fmt.Sprintf("Error in number of out args in GetAsMap : Expected <%d> but Given <%d>", 1, m1.Type.NumOut()))
	}
	if m1.Type.Out(0).Kind().String() != "map" {
		t.Fatal(fmt.Sprintf("Error in type of out args in GetAsMap : Expected <%s> but Given <%s>", "map", m1.Type.Out(0).Kind().String()))
	}
	m2, ok2 := reflect.ValueOf(serializable).Type().MethodByName("SetFromMap")
	if !ok2 {
		t.Fatal("Error checking SetFromMap method into the default serializable interface")
	}
	if m2.Type.NumOut() != 2 {
		t.Fatal(fmt.Sprintf("Error in number of in args in SetFromMap : Expected <%d> but Given <%d>", 2, m2.Type.NumIn()))
	}
	if m2.Type.NumIn() != 2 {
		t.Fatal(fmt.Sprintf("Error in number of in args in SetFromMap : Expected <%d> but Given <%d>", 2, m2.Type.NumIn()-m2.Type.NumOut()))
	}
	if m2.Type.In(1).Kind().String() != "map" {
		t.Fatal(fmt.Sprintf("Error in type of in args in SetFromMap : Expected <%s> but Given <%s>", "map", m2.Type.In(1).Kind().String()))
	}
	if m2.Type.Out(0).Kind().String() != "interface" {
		t.Fatal(fmt.Sprintf("Error in type of out args in SetFromMap : Expected <%s> but Given <%s>", "interface", m2.Type.Out(0).Kind().String()))
	}
}

func TestExtendedSerializableStruct(t *testing.T) {
	var now time.Time = time.Now()
	var name string = "Fabrizio"
	var age int = 44
	mp := &__serializableTestStruct{
		Name:      name,
		Age:       age,
		Birthdate: now,
	}
	attributesMap := mp.GetAsMap()
	nameMap, ok1 := attributesMap["Name"]
	ageMap, ok2 := attributesMap["Age"]
	bd, ok3 := attributesMap["Birthdate"]
	if !ok1 || !ok2 || !ok3 || nameMap != name || ageMap != age || bd != now {
		t.Fatal(fmt.Sprintf("Error Reading object : Expected <%v> but Given <%v>", mp, attributesMap))
	}
	fmt.Println(fmt.Sprintf("%v", mp))
	var newAge int = 45
	attributesMap["Age"] = newAge
	mp.SetFromMap(attributesMap)
	ageMap, ok2 = attributesMap["Age"]
	if !ok1 || !ok2 || !ok3 || nameMap != name || ageMap != newAge || bd != now {
		t.Fatal(fmt.Sprintf("Error Reading object : Expected <%v> but Given <%v>", mp, attributesMap))
	}

}
