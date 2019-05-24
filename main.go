package main

import (
	"fmt"
	ut "hellgate75/general_utils/utils"
)

func main() {
	fmt.Println("Preparing ...")
	var arr []int = make([]int, 5, 10)
	for index := 0; index < 5; index++ {
		arr[index] = index + 1
	}
	fmt.Println("Elements : ", len(arr))
	fmt.Println("Starting ...")
	nav := ut.IntArrayNav(arr)
	for nav.Next() {
		fmt.Println(nav.Get())
	}
}
