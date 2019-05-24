package utils

import (
	"fmt"
	"strconv"
	"testing"
)

func TestArrayNav(t *testing.T) {
	var arr []interface{} = make([]string, 5, 10)
	for index := 0; index < 5; index++ {
		arr[index] = string(index + 1)
	}
	nav := ArrayNav(arr)
	var expectedValue int = 1
	for nav.Next() {
		expectedValue++
		var valueStr string = (nav.Get()).(string)
		if string(expectedValue) != valueStr {
			t.Fatal("Expected <" + string(expectedValue) + "> but given <" + valueStr + ">")
		}
	}

}

func TestIntArrayNav(t *testing.T) {
	var arr []int = make([]int, 5, 10)
	for index := 0; index < 5; index++ {
		arr[index] = index + 1
	}
	nav := IntArrayNav(arr)
	var expectedValue int = 0
	for nav.Next() {
		expectedValue++
		if expectedValue != nav.Get() {
			t.Fatal("Expected <" + string(expectedValue) + "> but given <" + string(nav.Get()) + ">")
		}
	}

}

func TestFloatArrayNav(t *testing.T) {
	var arr []float64 = make([]float64, 5, 10)
	for index := 0; index < 5; index++ {
		arr[index] = float64(index) + 1.0
	}
	nav := FloatArrayNav(arr)
	var expectedValue float64 = 0.0
	for nav.Next() {
		expectedValue++
		if expectedValue != nav.Get() {
			var valueStr string = fmt.Sprintf("%f", nav.Get())
			var expectedStr string = fmt.Sprintf("%f", expectedValue)
			t.Fatal("Expected <" + expectedStr + "> but given <" + valueStr + ">")
		}
	}

}

func TestBoolArrayNav(t *testing.T) {
	var arr []bool = make([]bool, 5, 10)
	for index := 0; index < 5; index++ {
		arr[index] = (index%2 == 0)
	}
	nav := BoolArrayNav(arr)
	var counter int = 0
	var expectedValue bool = true
	for nav.Next() {
		counter++
		expectedValue = (counter%2 == 0)
		if expectedValue != nav.Get() {
			var valueStr string = strconv.FormatBool(nav.Get())
			var expectedStr string = strconv.FormatBool(expectedValue)
			t.Fatal("Expected <" + expectedStr + "> but given <" + valueStr + ">")
		}
	}

}
