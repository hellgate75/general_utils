package utils

import (
	"fmt"
	"github.com/hellgate75/general_utils/common"
)

type NavAttr struct {
	arr_sel []common.Type
	index   int
}

type IntNavAttr struct {
	arr_sel []int
	index   int
}

type FloatNavAttr struct {
	arr_sel []float64
	index   int
}

type BoolNavAttr struct {
	arr_sel []bool
	index   int
}

func (nav *NavAttr) Next() bool {
	flag := nav.index < len(nav.arr_sel)-1
	if flag {
		nav.index++
	}
	return flag
}

func (nav *IntNavAttr) Next() bool {
	flag := nav.index < len(nav.arr_sel)-1
	if flag {
		nav.index++
	}
	return flag
}

func (nav *FloatNavAttr) Next() bool {
	flag := nav.index < len(nav.arr_sel)-1
	if flag {
		nav.index++
	}
	return flag
}

func (nav *BoolNavAttr) Next() bool {
	flag := nav.index < len(nav.arr_sel)-1
	if flag {
		nav.index++
	}
	return flag
}

func (nav *NavAttr) Prev() bool {
	flag := nav.index >= 1
	if flag {
		nav.index--
	}
	return flag
}

func (nav *IntNavAttr) Prev() bool {
	flag := nav.index >= 1
	if flag {
		nav.index--
	}
	return flag
}

func (nav *FloatNavAttr) Prev() bool {
	flag := nav.index >= 1
	if flag {
		nav.index--
	}
	return flag
}

func (nav *BoolNavAttr) Prev() bool {
	flag := nav.index >= 1
	if flag {
		nav.index--
	}
	return flag
}

func (nav *NavAttr) Get() common.Type {
	if nav.index >= 0 {
		return nav.arr_sel[nav.index]
	}
	return nil
}

func (nav *IntNavAttr) Get() int {
	if nav.index >= 0 {
		return nav.arr_sel[nav.index]
	}
	return 0
}

func (nav *FloatNavAttr) Get() float64 {
	if nav.index >= 0 {
		return nav.arr_sel[nav.index]
	}
	return 0.0
}

func (nav *BoolNavAttr) Get() bool {
	if nav.index >= 0 {
		return nav.arr_sel[nav.index]
	}
	return false
}

func (nav *NavAttr) Print() {
	fmt.Println("NavAttr")
}

func (nav *NavAttr) Len() int {
	return len(nav.arr_sel)
}

func (nav *IntNavAttr) Len() int {
	return len(nav.arr_sel)
}

func (nav *FloatNavAttr) Len() int {
	return len(nav.arr_sel)
}

func (nav *BoolNavAttr) Len() int {
	return len(nav.arr_sel)
}

func (nav *NavAttr) Position() int {
	return nav.index
}

func (nav *IntNavAttr) Position() int {
	return nav.index
}

func (nav *FloatNavAttr) Position() int {
	return nav.index
}

func (nav *BoolNavAttr) Position() int {
	return nav.index
}

type baseArrayNav interface {
	Prev() bool
	Next() bool
	Len() int
	Position() int
}
type ArrayNav interface {
	baseArrayNav
	Get() common.Type
}

type IntArrayNav interface {
	baseArrayNav
	Get() int
}

type FloatArrayNav interface {
	baseArrayNav
	Get() float64
}

type BoolArrayNav interface {
	baseArrayNav
	Get() bool
}

func NewArrayNav(arr []common.Type) ArrayNav {
	return &NavAttr{
		arr_sel: arr,
		index:   -1,
	}
}

func NewIntArrayNav(arr []int) IntArrayNav {
	return &IntNavAttr{
		arr_sel: arr,
		index:   -1,
	}
}

func NewFloatArrayNav(arr []float64) FloatArrayNav {
	return &FloatNavAttr{
		arr_sel: arr,
		index:   -1,
	}
}

func NewBoolArrayNav(arr []bool) BoolArrayNav {
	return &BoolNavAttr{
		arr_sel: arr,
		index:   -1,
	}
}

/*
func arr_nav(arr []interface{}) interface{} {
	arr_sel := arr
	index := 0
	return interface{
		next : func() interface{} {
			if index < len(arr_sel) {
				index++
				return arr_sel[index]
			} else {
				return nil;
			}
		},
		prev : func() interface{} {
			if index > 1 {
				index--
				return arr_sel[index]
			} else {
				return nil;
			}
		}
		hasNext : func() bool {
			return index < len(arr_sel) - 1
		},
		hasPrev : func() bool {
			return index < 1
		},
		index : func() int {
			return index;
		},
		len : func() int {
			return len(arr_sel)
		},
	};
}
*/
