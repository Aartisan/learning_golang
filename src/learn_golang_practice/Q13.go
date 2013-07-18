package main

import (
	"fmt"
)

func max(arr []int) int {
	max := 0
	for _, i := range arr {
		if i > max {
			max = i
		}
	}
	return max
}

func min(arr []int) int {
	var min int
	min = arr[0]
	for _, i := range arr {
		if i < min {
			min = i
		}
	}
	return min
}

func mix(arr []int, mark string) int {
	var return_value int
	return_value = arr[0]
	if mark == "max" {
		for _, i := range arr {
			if i > return_value {
				return_value = i
			}
		}
	} else {
		for _, i := range arr {
			if i < return_value {
				return_value = i
			}
		}
	}
	return return_value
}

func main() {
	arr := []int{5, 6, 71, 23, 46, 9, 234, 64}
	fmt.Printf("In arr max is %d\n", max(arr))
	fmt.Printf("In arr min is %d\n", min(arr))
	fmt.Printf("In arr in mix func min is %d\n", mix(arr, "min"))
	fmt.Printf("In arr in mix func max is %d\n", mix(arr, "max"))
}
